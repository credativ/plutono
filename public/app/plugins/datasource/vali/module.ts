import { DataSourcePlugin } from '@credativ/plutono-data';
import Datasource from './datasource';

import ValiCheatSheet from './components/ValiCheatSheet';
import ValiExploreQueryEditor from './components/ValiExploreQueryEditor';
import ValiQueryEditor from './components/ValiQueryEditor';
import { ValiAnnotationsQueryCtrl } from './ValiAnnotationsQueryCtrl';
import { ConfigEditor } from './configuration/ConfigEditor';

export const plugin = new DataSourcePlugin(Datasource)
  .setQueryEditor(ValiQueryEditor)
  .setConfigEditor(ConfigEditor)
  .setExploreQueryField(ValiExploreQueryEditor)
  .setQueryEditorHelp(ValiCheatSheet)
  .setAnnotationQueryCtrl(ValiAnnotationsQueryCtrl);
