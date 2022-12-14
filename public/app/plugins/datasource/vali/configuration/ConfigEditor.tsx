import React from 'react';
import { DataSourcePluginOptionsEditorProps, DataSourceSettings } from '@credativ/plutono-data';
import { DataSourceHttpSettings } from '@credativ/plutono-ui';
import { ValiOptions } from '../types';
import { MaxLinesField } from './MaxLinesField';
import { DerivedFields } from './DerivedFields';

export type Props = DataSourcePluginOptionsEditorProps<ValiOptions>;

const makeJsonUpdater = <T extends any>(field: keyof ValiOptions) => (
  options: DataSourceSettings<ValiOptions>,
  value: T
): DataSourceSettings<ValiOptions> => {
  return {
    ...options,
    jsonData: {
      ...options.jsonData,
      [field]: value,
    },
  };
};

const setMaxLines = makeJsonUpdater('maxLines');
const setDerivedFields = makeJsonUpdater('derivedFields');

export const ConfigEditor = (props: Props) => {
  const { options, onOptionsChange } = props;

  return (
    <>
      <DataSourceHttpSettings
        defaultUrl={'http://localhost:3100'}
        dataSourceConfig={options}
        showAccessOptions={false}
        onChange={onOptionsChange}
      />

      <div className="gf-form-group">
        <div className="gf-form-inline">
          <div className="gf-form">
            <MaxLinesField
              value={options.jsonData.maxLines || ''}
              onChange={(value) => onOptionsChange(setMaxLines(options, value))}
            />
          </div>
        </div>
      </div>

      <DerivedFields
        value={options.jsonData.derivedFields}
        onChange={(value) => onOptionsChange(setDerivedFields(options, value))}
      />
    </>
  );
};
