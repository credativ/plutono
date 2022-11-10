import React, { FC, MouseEvent, useCallback } from 'react';

import { PlutonoTheme } from '@credativ/plutono-data';
import { css } from 'emotion';
import { useStyles } from '@credativ/plutono-ui';

import { TagBadge } from '../../../core/components/TagFilter/TagBadge';

interface Props {
  tags?: string[];
  remove?: boolean;
  onClick: (tag: string, remove?: boolean) => void;
}

export const AnnotationListItemTags: FC<Props> = ({ tags, remove, onClick }) => {
  const styles = useStyles(getStyles);
  const onTagClicked = useCallback(
    (e: MouseEvent, tag: string) => {
      e.stopPropagation();
      onClick(tag, remove);
    },
    [remove]
  );

  if (!tags || !tags.length) {
    return null;
  }

  return (
    <div>
      {tags.map((tag) => {
        return (
          <span key={tag} onClick={(e) => onTagClicked(e, tag)} className={styles.pointer}>
            <TagBadge label={tag} removeIcon={Boolean(remove)} count={0} />
          </span>
        );
      })}
    </div>
  );
};

function getStyles(theme: PlutonoTheme) {
  return {
    pointer: css`
      label: pointer;
      cursor: pointer;
      padding: ${theme.spacing.xxs};
    `,
  };
}
