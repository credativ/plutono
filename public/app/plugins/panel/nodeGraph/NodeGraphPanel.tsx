import React from 'react';
import { PanelProps } from '@credativ/plutono-data';
import { Options } from './types';
import { NodeGraph } from '@credativ/plutono-ui';
import { useLinks } from '../../../features/explore/utils/links';

export const NodeGraphPanel: React.FunctionComponent<PanelProps<Options>> = ({ width, height, data }) => {
  if (!data || !data.series.length) {
    return (
      <div className="panel-empty">
        <p>No data found in response</p>
      </div>
    );
  }

  const getLinks = useLinks(data.timeRange);

  return (
    <div style={{ width, height }}>
      <NodeGraph dataFrames={data.series} getLinks={getLinks} />
    </div>
  );
};
