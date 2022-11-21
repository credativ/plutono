import { UrlQueryMap } from '@credativ/plutono-data';

export interface LocationState {
  url: string;
  path: string;
  query: UrlQueryMap;
  routeParams: UrlQueryMap;
  replace: boolean;
  lastUpdated: number;
}
