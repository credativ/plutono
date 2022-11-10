import { DashboardDatasource } from './datasource';
import { DataSourcePlugin } from '@credativ/plutono-data';

export const plugin = new DataSourcePlugin(DashboardDatasource);
