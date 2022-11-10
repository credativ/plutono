import { PlutonoTheme } from '@credativ/plutono-data';
import { stylesFactory } from '@credativ/plutono-ui';
import { css } from 'emotion';

export const getStyles = stylesFactory((theme: PlutonoTheme, hidden: boolean) => ({
  color:
    hidden &&
    css`
      &,
      &:hover,
      label,
      a {
        color: ${hidden ? theme.colors.textFaint : theme.colors.text};
      }
    `,
}));
