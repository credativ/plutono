import { VizLegendOptions, GraphTooltipOptions } from '@credativ/plutono-ui';

export interface GraphOptions {
  // nothing for now
}

export interface OptionsWithLegend {
  legend: VizLegendOptions;
}

export interface Options extends OptionsWithLegend {
  graph: GraphOptions;
  tooltipOptions: GraphTooltipOptions;
}
