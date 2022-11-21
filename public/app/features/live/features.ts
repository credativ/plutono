import { LiveChannelConfig } from '@credativ/plutono-data';
import { MeasurementCollector } from '@credativ/plutono-runtime';
import { getDashboardChannelsFeature } from './dashboard/dashboardWatcher';
import { LiveMeasurementsSupport } from './measurements/measurementsSupport';
import { plutonoLiveCoreFeatures } from './scopes';

export function registerLiveFeatures() {
  const random2s = new MeasurementCollector();
  const randomFlakey = new MeasurementCollector();
  const channels: LiveChannelConfig[] = [
    {
      path: 'random-2s-stream',
      description: 'Random stream with points every 2s',
      getController: () => random2s,
      processMessage: random2s.addBatch,
    },
    {
      path: 'random-flakey-stream',
      description: 'Random stream with flakey data points',
      getController: () => randomFlakey,
      processMessage: randomFlakey.addBatch,
    },
  ];

  plutonoLiveCoreFeatures.register({
    name: 'testdata',
    support: {
      getChannelConfig: (path: string) => {
        return channels.find((c) => c.path === path);
      },
      getSupportedPaths: () => channels,
    },
    description: 'Test data generations',
  });

  const broadcastConfig: LiveChannelConfig = {
    path: '${path}',
    description: 'Broadcast any messages to a channel',
    canPublish: () => true,
  };

  plutonoLiveCoreFeatures.register({
    name: 'broadcast',
    support: {
      getChannelConfig: (path: string) => {
        return broadcastConfig;
      },
      getSupportedPaths: () => [broadcastConfig],
    },
    description: 'Broadcast will send/receive any JSON object in a channel',
  });

  plutonoLiveCoreFeatures.register({
    name: 'measurements',
    support: new LiveMeasurementsSupport(),
    description: 'These channels listen for measurements and produce DataFrames',
  });

  // dashboard/*
  plutonoLiveCoreFeatures.register(getDashboardChannelsFeature());
}
