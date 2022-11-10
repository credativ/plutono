import { DataQuery, DataSourceJsonData } from '@credativ/plutono-data';

export interface OpenTsdbQuery extends DataQuery {
  metric?: any;
}

export interface OpenTsdbOptions extends DataSourceJsonData {
  tsdbVersion: number;
  tsdbResolution: number;
  lookupLimit: number;
}
