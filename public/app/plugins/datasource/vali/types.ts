import { DataQuery, DataSourceJsonData, QueryResultMeta, ScopedVars } from '@credativ/plutono-data';

export interface ValiInstantQueryRequest {
  query: string;
  limit?: number;
  time?: string;
  direction?: 'BACKWARD' | 'FORWARD';
}

export interface ValiRangeQueryRequest {
  query: string;
  limit?: number;
  start?: number;
  end?: number;
  step?: number;
  direction?: 'BACKWARD' | 'FORWARD';
}

export enum ValiResultType {
  Stream = 'streams',
  Vector = 'vector',
  Matrix = 'matrix',
}

export interface ValiQuery extends DataQuery {
  expr: string;
  query?: string;
  format?: string;
  reverse?: boolean;
  legendFormat?: string;
  valueWithRefId?: boolean;
  maxLines?: number;
  range?: boolean;
  instant?: boolean;
}

export interface ValiOptions extends DataSourceJsonData {
  maxLines?: string;
  derivedFields?: DerivedFieldConfig[];
}

export interface ValiStats {
  [component: string]: {
    [label: string]: number;
  };
}

export interface ValiVectorResult {
  metric: { [label: string]: string };
  value: [number, string];
}

export interface ValiVectorResponse {
  status: string;
  data: {
    resultType: ValiResultType.Vector;
    result: ValiVectorResult[];
    stats?: ValiStats;
  };
}

export interface ValiMatrixResult {
  metric: Record<string, string>;
  values: Array<[number, string]>;
}

export interface ValiMatrixResponse {
  status: string;
  data: {
    resultType: ValiResultType.Matrix;
    result: ValiMatrixResult[];
    stats?: ValiStats;
  };
}

export interface ValiStreamResult {
  stream: Record<string, string>;
  values: Array<[string, string]>;
}

export interface ValiStreamResponse {
  status: string;
  data: {
    resultType: ValiResultType.Stream;
    result: ValiStreamResult[];
    stats?: ValiStats;
  };
}

export interface ValiTailResponse {
  streams: ValiStreamResult[];
  dropped_entries?: Array<{
    labels: Record<string, string>;
    timestamp: string;
  }> | null;
}

export type ValiResult = ValiVectorResult | ValiMatrixResult | ValiStreamResult;
export type ValiResponse = ValiVectorResponse | ValiMatrixResponse | ValiStreamResponse;

export interface ValiLogsStreamEntry {
  line: string;
  ts: string;
}

export interface ValiExpression {
  regexp: string;
  query: string;
}

export type DerivedFieldConfig = {
  matcherRegex: string;
  name: string;
  url?: string;
  datasourceUid?: string;
};

export interface TransformerOptions {
  format?: string;
  legendFormat?: string;
  step: number;
  start: number;
  end: number;
  query: string;
  responseListLength: number;
  refId: string;
  scopedVars: ScopedVars;
  meta?: QueryResultMeta;
  valueWithRefId?: boolean;
}
