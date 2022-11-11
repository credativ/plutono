import { PlutonoThemeType } from '@credativ/plutono-data';
import { selectThemeVariant } from './selectThemeVariant';
import { mockTheme } from './index';

const lightThemeMock = {
  color: {
    red: '#ff0000',
    green: '#00ff00',
  },
};

const darkThemeMock = {
  color: {
    red: '#ff0000',
    green: '#00ff00',
  },
};

describe('Theme variable variant selector', () => {
  // @ts-ignore
  const restoreTheme = mockTheme((name) => (name === PlutonoThemeType.Light ? lightThemeMock : darkThemeMock));

  afterAll(() => {
    restoreTheme();
  });
  it('return correct variable value for given theme', () => {
    const theme = lightThemeMock;

    const selectedValue = selectThemeVariant(
      {
        dark: theme.color.red,
        light: theme.color.green,
      },
      PlutonoThemeType.Light
    );

    expect(selectedValue).toBe(lightThemeMock.color.green);
  });

  it('return dark theme variant if no theme given', () => {
    const theme = lightThemeMock;

    const selectedValue = selectThemeVariant({
      dark: theme.color.red,
      light: theme.color.green,
    });

    expect(selectedValue).toBe(lightThemeMock.color.red);
  });
});
