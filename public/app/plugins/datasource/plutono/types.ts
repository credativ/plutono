import { AnnotationQuery, DataQuery } from '@credativ/plutono-data';
import { MeasurementsQuery } from '@credativ/plutono-runtime';

//----------------------------------------------
// Query
//----------------------------------------------

export enum PlutonoQueryType {
  RandomWalk = 'randomWalk',
  LiveMeasurements = 'measurements',
}

export interface PlutonoQuery extends DataQuery {
  queryType: PlutonoQueryType; // RandomWalk by default
  channel?: string;
  measurements?: MeasurementsQuery;
}

export const defaultQuery: PlutonoQuery = {
  refId: 'A',
  queryType: PlutonoQueryType.RandomWalk,
};

//----------------------------------------------
// Annotations
//----------------------------------------------

export enum PlutonoAnnotationType {
  Dashboard = 'dashboard',
  Tags = 'tags',
}

export interface PlutonoAnnotationQuery extends AnnotationQuery<PlutonoQuery> {
  type: PlutonoAnnotationType; // tags
  limit: number; // 100
  tags?: string[];
  matchAny?: boolean; // By default Plutono only shows annotations that match all tags in the query. Enabling this returns annotations that match any of the tags in the query.
}
