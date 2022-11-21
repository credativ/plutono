import { DataSourceSettings } from '@credativ/plutono-data';
import { PromOptions } from '../types';
import { createDatasourceSettings } from '../../../../features/datasources/mocks';

export function createDefaultConfigOptions(): DataSourceSettings<PromOptions> {
  return createDatasourceSettings<PromOptions>({
    timeInterval: '1m',
    queryTimeout: '1m',
    httpMethod: 'GET',
    directUrl: 'url',
  });
}
