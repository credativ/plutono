import React from 'react';
import { DataQueryError } from '@credativ/plutono-data';
import { JSONFormatter } from '@credativ/plutono-ui';

interface InspectErrorTabProps {
  error?: DataQueryError;
}

export const InspectErrorTab: React.FC<InspectErrorTabProps> = ({ error }) => {
  if (!error) {
    return null;
  }
  if (error.data) {
    return (
      <>
        <h3>{error.data.message}</h3>
        <JSONFormatter json={error} open={2} />
      </>
    );
  }
  return <div>{error.message}</div>;
};
