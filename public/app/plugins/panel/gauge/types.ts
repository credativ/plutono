import { VizOrientation, SelectableValue } from '@credativ/plutono-data';
import { SingleStatBaseOptions } from '@credativ/plutono-ui/src/components/SingleStatShared/SingleStatBaseOptions';

export interface GaugeOptions extends SingleStatBaseOptions {
  showThresholdLabels: boolean;
  showThresholdMarkers: boolean;
}

export const orientationOptions: Array<SelectableValue<VizOrientation>> = [
  { value: VizOrientation.Auto, label: 'Auto' },
  { value: VizOrientation.Horizontal, label: 'Horizontal' },
  { value: VizOrientation.Vertical, label: 'Vertical' },
];
