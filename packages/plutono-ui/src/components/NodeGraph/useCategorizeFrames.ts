import { useMemo } from 'react';
import { DataFrame } from '@credativ/plutono-data';

/**
 * As we need 2 dataframes for the service map, one with nodes and one with edges we have to figure out which is which.
 * Right now we do not have any metadata for it so we just check preferredVisualisationType and then column names.
 * TODO: maybe we could use column labels to have a better way to do this
 */
export function useCategorizeFrames(series: DataFrame[]) {
  return useMemo(() => {
    const serviceMapFrames = series.filter((frame) => {
      // Use the data frame if the preferred visualization type was set to nodeGraph.
      if (frame.meta?.preferredVisualisationType === 'nodeGraph') {
        return true;
      }
      // Use the data frame if its name is 'nodes' or 'edges'.
      if (frame.name === 'nodes' || frame.name === 'edges') {
        return true;
      }
      // Use the data frame if the query was named 'nodes' or 'edges'.
      if (frame.refId === 'nodes' || frame.refId === 'edges') {
        return true;
      }
      return false;
    });
    return serviceMapFrames.reduce(
      (acc, frame) => {
        const sourceField = frame.fields.filter((f) => f.name === 'source');
        if (sourceField.length) {
          acc.edges.push(frame);
        } else {
          acc.nodes.push(frame);
        }
        return acc;
      },
      { edges: [], nodes: [] } as { nodes: DataFrame[]; edges: DataFrame[] }
    );
  }, [series]);
}
