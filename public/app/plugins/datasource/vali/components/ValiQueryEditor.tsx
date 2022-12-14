// Libraries
import React, { memo } from 'react';

// Types
import { QueryEditorProps } from '@credativ/plutono-data';
import { InlineFormLabel } from '@credativ/plutono-ui';
import { ValiDatasource } from '../datasource';
import { ValiQuery, ValiOptions } from '../types';
import { ValiQueryField } from './ValiQueryField';

type Props = QueryEditorProps<ValiDatasource, ValiQuery, ValiOptions>;

export function ValiQueryEditor(props: Props) {
  const { range, query, data, datasource, onChange, onRunQuery } = props;

  const onLegendChange = (e: React.SyntheticEvent<HTMLInputElement>) => {
    const nextQuery = { ...query, legendFormat: e.currentTarget.value };
    onChange(nextQuery);
  };

  const legendField = (
    <div className="gf-form-inline">
      <div className="gf-form">
        <InlineFormLabel
          width={6}
          tooltip="Controls the name of the time series, using name or pattern. For example
        {{hostname}} will be replaced with label value for the label hostname. The legend only applies to metric queries."
        >
          Legend
        </InlineFormLabel>
        <input
          type="text"
          className="gf-form-input"
          placeholder="legend format"
          value={query.legendFormat || ''}
          onChange={onLegendChange}
          onBlur={onRunQuery}
        />
      </div>
    </div>
  );

  return (
    <ValiQueryField
      datasource={datasource}
      query={query}
      onChange={onChange}
      onRunQuery={onRunQuery}
      onBlur={onRunQuery}
      history={[]}
      data={data}
      range={range}
      runOnBlur={true}
      ExtraFieldElement={legendField}
    />
  );
}

export default memo(ValiQueryEditor);
