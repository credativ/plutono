import darkTheme from './dark';
import lightTheme from './light';
import { PlutonoTheme } from '@credativ/plutono-data';

let themeMock: ((name?: string) => PlutonoTheme) | null;

export const getTheme = (name?: string) =>
  (themeMock && themeMock(name)) || (name === 'light' ? lightTheme : darkTheme);

export const mockTheme = (mock: (name?: string) => PlutonoTheme) => {
  themeMock = mock;
  return () => {
    themeMock = null;
  };
};
