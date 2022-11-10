import { DataSourcePlugin } from '@credativ/plutono-data';
import { TempoDatasource } from './datasource';
import { TempoQueryField } from './QueryField';
import { ConfigEditor } from './ConfigEditor';

export const plugin = new DataSourcePlugin(TempoDatasource)
  .setConfigEditor(ConfigEditor)
  .setExploreQueryField(TempoQueryField);
