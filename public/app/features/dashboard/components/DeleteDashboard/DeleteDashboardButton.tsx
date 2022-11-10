import React from 'react';
import { DeleteDashboardModal } from './DeleteDashboardModal';
import { Button, ModalsController } from '@credativ/plutono-ui';
import { DashboardModel } from '../../state';

type Props = {
  dashboard: DashboardModel;
};

export const DeleteDashboardButton = ({ dashboard }: Props) => (
  <ModalsController>
    {({ showModal, hideModal }) => (
      <Button
        variant="destructive"
        onClick={() => {
          showModal(DeleteDashboardModal, {
            dashboard,
            hideModal,
          });
        }}
        aria-label="Dashboard settings page delete dashboard button"
      >
        Delete Dashboard
      </Button>
    )}
  </ModalsController>
);
