import { PieChartType, SingleStatBaseOptions, PieChartLabels, PieChartLegendOptions } from '@credativ/plutono-ui';

export interface PieChartOptions extends SingleStatBaseOptions {
  pieType: PieChartType;
  displayLabels: PieChartLabels[];
  legend: PieChartLegendOptions;
}
