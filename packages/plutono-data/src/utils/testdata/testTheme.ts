import { PlutonoTheme, PlutonoThemeType } from '../../types/theme';

export function getTestTheme(type: PlutonoThemeType = PlutonoThemeType.Dark): PlutonoTheme {
  return ({
    type,
    isDark: type === PlutonoThemeType.Dark,
    isLight: type === PlutonoThemeType.Light,
    colors: {
      panelBg: 'white',
    },
  } as unknown) as PlutonoTheme;
}
