import { DataSourcePluginOptionsEditorProps } from '@credativ/plutono-data';
import { DataSourceHttpSettings } from '@credativ/plutono-ui';
import { TraceToLogsSettings } from 'app/core/components/TraceToLogsSettings';
import React from 'react';

export type Props = DataSourcePluginOptionsEditorProps;

export const ConfigEditor: React.FC<Props> = ({ options, onOptionsChange }) => {
  return (
    <>
      <DataSourceHttpSettings
        defaultUrl="http://localhost:9411"
        dataSourceConfig={options}
        showAccessOptions={false}
        onChange={onOptionsChange}
      />

      <TraceToLogsSettings options={options} onOptionsChange={onOptionsChange} />
    </>
  );
};
