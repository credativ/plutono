// Libraries
import React, { memo } from 'react';

// Types
import { ValiQuery } from '../types';
import { ValiQueryFieldForm } from './ValiQueryFieldForm';
import ValiDatasource from '../datasource';

interface Props {
  expr: string;
  maxLines?: number;
  instant?: boolean;
  datasource: ValiDatasource;
  onChange: (query: ValiQuery) => void;
}

export const ValiAnnotationsQueryEditor = memo(function ValiAnnotationQueryEditor(props: Props) {
  const { expr, maxLines, instant, datasource, onChange } = props;

  // Timerange to get existing labels from. Hard coding like this seems to be good enough right now.
  const absolute = {
    from: Date.now() - 10000,
    to: Date.now(),
  };

  const queryWithRefId: ValiQuery = {
    refId: '',
    expr,
    maxLines,
    instant,
  };
  return (
    <div className="gf-form-group">
      <ValiQueryFieldForm
        datasource={datasource}
        query={queryWithRefId}
        onChange={onChange}
        onRunQuery={() => {}}
        history={[]}
        absoluteRange={absolute}
      />
    </div>
  );
});
