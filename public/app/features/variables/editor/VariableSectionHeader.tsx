import React, { PropsWithChildren, ReactElement } from 'react';
import { useStyles } from '@credativ/plutono-ui';
import { PlutonoTheme } from '@credativ/plutono-data';
import { css } from 'emotion';

interface VariableSectionHeaderProps {
  name: string;
}

export function VariableSectionHeader({ name }: PropsWithChildren<VariableSectionHeaderProps>): ReactElement {
  const styles = useStyles(getStyles);

  return <h5 className={styles.sectionHeading}>{name}</h5>;
}

function getStyles(theme: PlutonoTheme) {
  return {
    sectionHeading: css`
      label: sectionHeading;
      font-size: ${theme.typography.size.md};
      margin-bottom: ${theme.spacing.sm};
    `,
  };
}
