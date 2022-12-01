import React, { FunctionComponent } from 'react';
import { ValiQueryFieldForm, ValiQueryFieldFormProps } from './ValiQueryFieldForm';

type ValiQueryFieldProps = Omit<
  ValiQueryFieldFormProps,
  'labelsLoaded' | 'onLoadOptions' | 'onLabelsRefresh' | 'absoluteRange'
>;

export const ValiQueryField: FunctionComponent<ValiQueryFieldProps> = (props) => {
  const { datasource, range, ...otherProps } = props;
  const absoluteTimeRange = { from: range!.from!.valueOf(), to: range!.to!.valueOf() }; // Range here is never optional

  return <ValiQueryFieldForm datasource={datasource} absoluteRange={absoluteTimeRange} {...otherProps} />;
};

export default ValiQueryField;
