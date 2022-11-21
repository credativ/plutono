import { PanelPlugin } from '@credativ/plutono-data';
import { NodeGraphPanel } from './NodeGraphPanel';
import { Options } from './types';

export const plugin = new PanelPlugin<Options>(NodeGraphPanel);
