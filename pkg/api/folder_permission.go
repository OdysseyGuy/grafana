package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/dashboards"
	"github.com/grafana/grafana/pkg/services/guardian"
	"github.com/grafana/grafana/pkg/util"
)

func (hs *HTTPServer) GetFolderPermissionList(c *models.ReqContext) response.Response {
	s := dashboards.NewFolderService(c.OrgId, c.SignedInUser, hs.SQLStore)
	folder, err := s.GetFolderByUID(c.Params(":uid"))

	if err != nil {
		return ToFolderErrorResponse(err)
	}

	g := guardian.New(folder.Id, c.OrgId, c.SignedInUser)

	if canAdmin, err := g.CanAdmin(); err != nil || !canAdmin {
		return ToFolderErrorResponse(models.ErrFolderAccessDenied)
	}

	acl, err := g.GetAcl()
	if err != nil {
		return response.Error(http.StatusInternalServerError, "Failed to get folder permissions", err)
	}

	filteredAcls := make([]*models.DashboardAclInfoDTO, 0, len(acl))
	for _, perm := range acl {
		if perm.UserId > 0 && dtos.IsHiddenUser(perm.UserLogin, c.SignedInUser, hs.Cfg) {
			continue
		}

		perm.FolderId = folder.Id
		perm.DashboardId = 0

		perm.UserAvatarUrl = dtos.GetGravatarUrl(perm.UserEmail)

		if perm.TeamId > 0 {
			perm.TeamAvatarUrl = dtos.GetGravatarUrlWithDefault(perm.TeamEmail, perm.Team)
		}

		if perm.Slug != "" {
			perm.Url = models.GetDashboardFolderUrl(perm.IsFolder, perm.Uid, perm.Slug)
		}

		filteredAcls = append(filteredAcls, perm)
	}

	return response.JSON(http.StatusOK, filteredAcls)
}

func (hs *HTTPServer) UpdateFolderPermissions(c *models.ReqContext, apiCmd dtos.UpdateDashboardAclCommand) response.Response {
	if err := validatePermissionsUpdate(apiCmd); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), err)
	}

	s := dashboards.NewFolderService(c.OrgId, c.SignedInUser, hs.SQLStore)
	folder, err := s.GetFolderByUID(c.Params(":uid"))
	if err != nil {
		return ToFolderErrorResponse(err)
	}

	g := guardian.New(folder.Id, c.OrgId, c.SignedInUser)
	canAdmin, err := g.CanAdmin()
	if err != nil {
		return ToFolderErrorResponse(err)
	}

	if !canAdmin {
		return ToFolderErrorResponse(models.ErrFolderAccessDenied)
	}

	var items []*models.DashboardAcl
	for _, item := range apiCmd.Items {
		items = append(items, &models.DashboardAcl{
			OrgID:       c.OrgId,
			DashboardID: folder.Id,
			UserID:      item.UserID,
			TeamID:      item.TeamID,
			Role:        item.Role,
			Permission:  item.Permission,
			Created:     time.Now(),
			Updated:     time.Now(),
		})
	}

	hiddenACL, err := g.GetHiddenACL(hs.Cfg)
	if err != nil {
		return response.Error(http.StatusInternalServerError, "Error while retrieving hidden permissions", err)
	}
	items = append(items, hiddenACL...)

	if okToUpdate, err := g.CheckPermissionBeforeUpdate(models.PERMISSION_ADMIN, items); err != nil || !okToUpdate {
		if err != nil {
			if errors.Is(err, guardian.ErrGuardianPermissionExists) ||
				errors.Is(err, guardian.ErrGuardianOverride) {
				return response.Error(http.StatusBadRequest, err.Error(), err)
			}

			return response.Error(http.StatusInternalServerError, "Error while checking folder permissions", err)
		}

		return response.Error(http.StatusForbidden, "Cannot remove own admin permission for a folder", nil)
	}

	if err := updateDashboardACL(hs, folder.Id, items); err != nil {
		if errors.Is(err, models.ErrDashboardAclInfoMissing) {
			err = models.ErrFolderAclInfoMissing
		}
		if errors.Is(err, models.ErrDashboardPermissionDashboardEmpty) {
			err = models.ErrFolderPermissionFolderEmpty
		}

		if errors.Is(err, models.ErrFolderAclInfoMissing) || errors.Is(err, models.ErrFolderPermissionFolderEmpty) {
			return response.Error(http.StatusConflict, err.Error(), err)
		}

		return response.Error(http.StatusInternalServerError, "Failed to create permission", err)
	}

	return response.JSON(http.StatusOK, util.DynMap{
		"message": "Folder permissions updated",
		"id":      folder.Id,
		"title":   folder.Title,
	})

}
