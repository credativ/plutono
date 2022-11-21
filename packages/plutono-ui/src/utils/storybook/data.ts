import { applyFieldOverrides, DataFrame, PlutonoTheme } from '@credativ/plutono-data';

export function prepDataForStorybook(data: DataFrame[], theme: PlutonoTheme) {
  return applyFieldOverrides({
    data: data,
    fieldConfig: {
      overrides: [],
      defaults: {},
    },
    theme,
    replaceVariables: (value: string) => value,
  });
}
