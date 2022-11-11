import { PlutonoThemeType } from '@credativ/plutono-data';

type VariantDescriptor = { [key in PlutonoThemeType]: string | number };

/**
 * @deprecated use theme.isLight ? or theme.isDark instead
 */
export const selectThemeVariant = (variants: VariantDescriptor, currentTheme?: PlutonoThemeType) => {
  return variants[currentTheme || PlutonoThemeType.Dark];
};
