import { config, PlutonoBootConfig } from '@credativ/plutono-runtime';
// Legacy binding paths
export { config, PlutonoBootConfig as Settings };

let plutonoConfig: PlutonoBootConfig = config;

export default plutonoConfig;

export const getConfig = () => {
  return plutonoConfig;
};

export const updateConfig = (update: Partial<PlutonoBootConfig>) => {
  plutonoConfig = {
    ...plutonoConfig,
    ...update,
  };
};
