import { ValiDatasource, VALI_ENDPOINT } from './datasource';
import { AbsoluteTimeRange, DataSourceSettings } from '@credativ/plutono-data';
import { ValiOptions } from './types';
import { createDatasourceSettings } from '../../../features/datasources/mocks';

interface Labels {
  [label: string]: string[];
}

interface Series {
  [label: string]: string;
}

interface SeriesForSelector {
  [selector: string]: Series[];
}

export function makeMockValiDatasource(labelsAndValues: Labels, series?: SeriesForSelector): ValiDatasource {
  const valiLabelsAndValuesEndpointRegex = /^\/vali\/api\/v1\/label\/(\w*)\/values/;
  const valiSeriesEndpointRegex = /^\/vali\/api\/v1\/series/;

  const valiLabelsEndpoint = `${VALI_ENDPOINT}/label`;
  const rangeMock: AbsoluteTimeRange = {
    from: 1560153109000,
    to: 1560163909000,
  };

  const labels = Object.keys(labelsAndValues);
  return {
    getTimeRangeParams: () => rangeMock,
    metadataRequest: (url: string, params?: { [key: string]: string }) => {
      if (url === valiLabelsEndpoint) {
        return labels;
      } else {
        const labelsMatch = url.match(valiLabelsAndValuesEndpointRegex);
        const seriesMatch = url.match(valiSeriesEndpointRegex);
        if (labelsMatch) {
          return labelsAndValues[labelsMatch[1]] || [];
        } else if (seriesMatch && series && params) {
          return series[params.match] || [];
        } else {
          throw new Error(`Unexpected url error, ${url}`);
        }
      }
    },
  } as any;
}

export function createDefaultConfigOptions(): DataSourceSettings<ValiOptions> {
  return createDatasourceSettings<ValiOptions>({
    maxLines: '531',
  });
}
