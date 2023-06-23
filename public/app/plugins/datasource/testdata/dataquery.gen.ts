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

export const DataQueryModelVersion = Object.freeze([0, 0]);

export enum TestDataQueryType {
  Annotations = 'annotations',
  Arrow = 'arrow',
  CSVContent = 'csv_content',
  CSVFile = 'csv_file',
  CSVMetricValues = 'csv_metric_values',
  DataPointsOutsideRange = 'datapoints_outside_range',
  ExponentialHeatmapBucketData = 'exponential_heatmap_bucket_data',
  FlameGraph = 'flame_graph',
  GrafanaAPI = 'grafana_api',
  LinearHeatmapBucketData = 'linear_heatmap_bucket_data',
  Live = 'live',
  Logs = 'logs',
  ManualEntry = 'manual_entry',
  NoDataPoints = 'no_data_points',
  NodeGraph = 'node_graph',
  PredictableCSVWave = 'predictable_csv_wave',
  PredictablePulse = 'predictable_pulse',
  RandomWalk = 'random_walk',
  RandomWalkTable = 'random_walk_table',
  RandomWalkWithError = 'random_walk_with_error',
  RawFrame = 'raw_frame',
  ServerError500 = 'server_error_500',
  Simulation = 'simulation',
  SlowQuery = 'slow_query',
  StreamingClient = 'streaming_client',
  TableStatic = 'table_static',
  Trace = 'trace',
  USA = 'usa',
  VariablesQuery = 'variables-query',
}

export interface StreamingQuery {
  bands?: number;
  noise: number;
  speed: number;
  spread: number;
  type: ('signal' | 'logs' | 'fetch');
  url?: string;
}

export interface PulseWaveQuery {
  offCount?: number;
  offValue?: number;
  onCount?: number;
  onValue?: number;
  timeStep?: number;
}

export interface SimulationQuery {
  config?: Record<string, unknown>;
  key: {
    type: string;
    tick: number;
    uid?: string;
  };
  last?: boolean;
  stream?: boolean;
}

export interface NodesQuery {
  count?: number;
  type?: ('random' | 'response' | 'random edges');
}

export interface USAQuery {
  fields?: Array<string>;
  mode?: string;
  period?: string;
  states?: Array<string>;
}

export const defaultUSAQuery: Partial<USAQuery> = {
  fields: [],
  states: [],
};

export interface CSVWave {
  labels?: string;
  name?: string;
  timeStep?: number;
  valuesCSV?: string;
}

/**
 * TODO: Should this live here given it's not used in the dataquery?
 */
export interface Scenario {
  description?: string;
  hideAliasField?: boolean;
  id: string;
  name: string;
  stringInput: string;
}

export interface TestData extends common.DataQuery {
  alias?: string;
  channel?: string;
  csvContent?: string;
  csvFileName?: string;
  csvWave?: Array<CSVWave>; // TODO can we prevent partial from being generated
  /**
   * Drop percentage (the chance we will lose a point 0-100)
   */
  dropPercent?: number;
  errorType?: ('server_panic' | 'frontend_exception' | 'frontend_observable');
  labels?: string;
  levelColumn?: boolean;
  lines?: number;
  nodes?: NodesQuery;
  points?: Array<Array<(string | number)>>;
  pulseWave?: PulseWaveQuery;
  rawFrameContent?: string;
  scenarioId?: TestDataQueryType;
  seriesCount?: number;
  sim?: SimulationQuery;
  spanCount?: number;
  stream?: StreamingQuery;
  stringInput?: string;
  usa?: USAQuery;
}

export const defaultTestData: Partial<TestData> = {
  csvWave: [],
  points: [],
  scenarioId: TestDataQueryType.RandomWalk,
};
