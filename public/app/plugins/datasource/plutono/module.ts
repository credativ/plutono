import { DataSourcePlugin } from '@credativ/plutono-data';
import { PlutonoDatasource } from './datasource';
import { QueryEditor } from './components/QueryEditor';
import { PlutonoQuery } from './types';
import { PlutonoAnnotationsQueryCtrl } from './annotation_ctrl';

export const plugin = new DataSourcePlugin<PlutonoDatasource, PlutonoQuery>(PlutonoDatasource)
  .setQueryEditor(QueryEditor)
  .setAnnotationQueryCtrl(PlutonoAnnotationsQueryCtrl);
