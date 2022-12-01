// Libraries
import React, { memo } from 'react';
import _ from 'lodash';

// Types
import { ExploreQueryFieldProps } from '@credativ/plutono-data';
import { ValiDatasource } from '../datasource';
import { ValiQuery, ValiOptions } from '../types';
import { ValiQueryField } from './ValiQueryField';

type Props = ExploreQueryFieldProps<ValiDatasource, ValiQuery, ValiOptions>;

export function ValiExploreQueryEditor(props: Props) {
  const { range, query, data, datasource, history, onChange, onRunQuery } = props;

  return (
    <ValiQueryField
      datasource={datasource}
      query={query}
      onChange={onChange}
      onBlur={() => {}}
      onRunQuery={onRunQuery}
      history={history}
      data={data}
      range={range}
    />
  );
}

export default memo(ValiExploreQueryEditor);
