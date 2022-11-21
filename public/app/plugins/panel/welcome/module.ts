import { PanelPlugin } from '@credativ/plutono-data';
import { WelcomeBanner } from './Welcome';

export const plugin = new PanelPlugin(WelcomeBanner).setNoPadding();
