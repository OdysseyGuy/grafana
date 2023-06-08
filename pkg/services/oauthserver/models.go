package oauthserver

import (
	"context"
	"net/http"

	"github.com/grafana/grafana/pkg/plugins/oauth"
	"gopkg.in/square/go-jose.v2"
)

const (
	// TmpOrgID is the orgID we use while global service accounts are not supported.
	TmpOrgID int64 = 1
	// NoServiceAccountID is the ID we use for client that have no service account associated.
	NoServiceAccountID int64 = 0

	// List of scopes used to identify the impersonated user.
	ScopeUsersSelf       = "users:self"
	ScopeGlobalUsersSelf = "global.users:self"
	ScopeTeamsSelf       = "teams:self"

	// Supported encryptions
	RS256 = "RS256"
	ES256 = "ES256"
)

// OAuth2Server represents a service in charge of managing OAuth2 clients
// and handling OAuth2 requests (token, introspection).
type OAuth2Server interface {
	oauth.ExternalServiceRegister
	// GetExternalService retrieves an external service from store by client_id. It populates the SelfPermissions and
	// SignedInUser from the associated service account.
	GetExternalService(ctx context.Context, id string) (*ExternalService, error)

	// HandleTokenRequest handles the client's OAuth2 query to obtain an access_token by presenting its authorization
	// grant (ex: client_credentials, jwtbearer).
	HandleTokenRequest(rw http.ResponseWriter, req *http.Request)
	// HandleIntrospectionRequest handles the OAuth2 query to determine the active state of an OAuth 2.0 token and
	// to determine meta-information about this token.
	HandleIntrospectionRequest(rw http.ResponseWriter, req *http.Request)
}

//go:generate mockery --name Store --structname MockStore --outpkg oauthtest --filename store_mock.go --output ./oauthtest/

type Store interface {
	RegisterExternalService(ctx context.Context, client *ExternalService) error
	SaveExternalService(ctx context.Context, client *ExternalService) error
	GetExternalService(ctx context.Context, id string) (*ExternalService, error)
	GetExternalServiceByName(ctx context.Context, name string) (*ExternalService, error)
	GetExternalServicePublicKey(ctx context.Context, clientID string) (*jose.JSONWebKey, error)
}
