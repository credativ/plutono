import { ActiveDimensions, TooltipMode } from '../../Chart/Tooltip';
import { Dimension, Dimensions, TimeZone } from '@credativ/plutono-data';

export interface GraphTooltipOptions {
  mode: TooltipMode;
}

export interface GraphDimensions extends Dimensions {
  xAxis: Dimension<number>;
  yAxis: Dimension<number>;
}

export interface GraphTooltipContentProps {
  dimensions: GraphDimensions; // Dimension[]
  activeDimensions: ActiveDimensions<GraphDimensions>;
  timeZone?: TimeZone;
}
