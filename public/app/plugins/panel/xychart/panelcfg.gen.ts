// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     TSTypesJenny
//     PluginTSTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

import * as common from '@grafana/schema';

export const PanelCfgModelVersion = Object.freeze([0, 0]);

export enum SeriesMapping {
  Auto = 'auto',
  Manual = 'manual',
}

export enum ScatterShow {
  Lines = 'lines',
  Points = 'points',
  PointsAndLines = 'points+lines',
}

export interface XYDimensionConfig {
  exclude?: Array<string>;
  frame: number;
  x?: string;
}

export const defaultXYDimensionConfig: Partial<XYDimensionConfig> = {
  exclude: [],
};

export interface PanelFieldConfig extends common.HideableFieldConfig, common.AxisConfig {
  fillOpacity?: number;
  label?: common.VisibilityMode;
  labelValue?: common.TextDimensionConfig;
  lineColor?: common.ColorDimensionConfig;
  lineStyle?: common.LineStyle;
  lineWidth?: number;
  pointColor?: common.ColorDimensionConfig;
  pointSize?: common.ScaleDimensionConfig;
  pointSymbol?: common.ResourceDimensionConfig;
  show?: ScatterShow;
}

export const defaultPanelFieldConfig: Partial<PanelFieldConfig> = {
  fillOpacity: 0.5,
  label: common.VisibilityMode.Auto,
  show: ScatterShow.Points,
};

export interface ScatterSeriesConfig extends PanelFieldConfig {
  name?: string;
  x?: string;
  y?: string;
}

export interface PanelOptions extends common.OptionsWithLegend, common.OptionsWithTooltip {
  dims: XYDimensionConfig;
  series: Array<ScatterSeriesConfig>;
  seriesMapping?: SeriesMapping;
}

export const defaultPanelOptions: Partial<PanelOptions> = {
  series: [],
};
