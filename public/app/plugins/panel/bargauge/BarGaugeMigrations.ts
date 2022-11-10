import { PanelModel } from '@credativ/plutono-data';
import { sharedSingleStatMigrationHandler } from '@credativ/plutono-ui';
import { BarGaugeOptions } from './types';

export const barGaugePanelMigrationHandler = (panel: PanelModel<BarGaugeOptions>): Partial<BarGaugeOptions> => {
  return sharedSingleStatMigrationHandler(panel);
};
