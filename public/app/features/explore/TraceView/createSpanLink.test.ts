import { DataSourceInstanceSettings, ScopedVars } from '@credativ/plutono-data';
import { setDataSourceSrv, setTemplateSrv } from '@credativ/plutono-runtime';
import { createSpanLinkFactory } from './createSpanLink';

describe('createSpanLinkFactory', () => {
  it('returns undefined if there is no data source uid', () => {
    const splitOpenFn = jest.fn();
    const createLink = createSpanLinkFactory(splitOpenFn);
    expect(createLink).not.toBeDefined();
  });

  describe('should return link', () => {
    beforeAll(() => {
      setDataSourceSrv({
        getInstanceSettings(uid: string): DataSourceInstanceSettings | undefined {
          return {
            uid: 'vali1',
            name: 'vali1',
          } as any;
        },
      } as any);

      setTemplateSrv({
        replace(target?: string, scopedVars?: ScopedVars, format?: string | Function): string {
          return target!;
        },
      } as any);
    });

    it('with default keys when tags not configured', () => {
      const splitOpenFn = jest.fn();
      const createLink = createSpanLinkFactory(splitOpenFn, { datasourceUid: 'valiUid' });
      expect(createLink).toBeDefined();
      const linkDef = createLink!({
        startTime: new Date('2020-10-14T01:00:00Z').valueOf() * 1000,
        duration: 1000 * 1000,
        tags: [
          {
            key: 'host',
            value: 'host',
          },
        ],
        process: {
          tags: [
            {
              key: 'cluster',
              value: 'cluster1',
            },
            {
              key: 'hostname',
              value: 'hostname1',
            },
            {
              key: 'label2',
              value: 'val2',
            },
          ],
        } as any,
      } as any);

      expect(linkDef.href).toBe(
        `/explore?left={"range":{"from":"20201014T000000","to":"20201014T010006"},"datasource":"vali1","queries":[{"expr":"{cluster=\\"cluster1\\", hostname=\\"hostname1\\"}","refId":""}]}`
      );
    });

    it('with tags that passed in and without tags that are not in the span', () => {
      const splitOpenFn = jest.fn();
      const createLink = createSpanLinkFactory(splitOpenFn, { datasourceUid: 'valiUid', tags: ['ip', 'newTag'] });
      expect(createLink).toBeDefined();
      const linkDef = createLink!({
        startTime: new Date('2020-10-14T01:00:00Z').valueOf() * 1000,
        duration: 1000 * 1000,
        tags: [
          {
            key: 'host',
            value: 'host',
          },
        ],
        process: {
          tags: [
            {
              key: 'hostname',
              value: 'hostname1',
            },
            {
              key: 'ip',
              value: '192.168.0.1',
            },
          ],
        } as any,
      } as any);

      expect(linkDef.href).toBe(
        `/explore?left={"range":{"from":"20201014T000000","to":"20201014T010006"},"datasource":"vali1","queries":[{"expr":"{ip=\\"192.168.0.1\\"}","refId":""}]}`
      );
    });

    it('from tags and process tags as well', () => {
      const splitOpenFn = jest.fn();
      const createLink = createSpanLinkFactory(splitOpenFn, {
        datasourceUid: 'valiUid',
        tags: ['ip', 'host'],
      });
      expect(createLink).toBeDefined();
      const linkDef = createLink!({
        startTime: new Date('2020-10-14T01:00:00Z').valueOf() * 1000,
        duration: 1000 * 1000,
        tags: [
          {
            key: 'host',
            value: 'host',
          },
        ],
        process: {
          tags: [
            {
              key: 'hostname',
              value: 'hostname1',
            },
            {
              key: 'ip',
              value: '192.168.0.1',
            },
          ],
        } as any,
      } as any);

      expect(linkDef.href).toBe(
        `/explore?left={"range":{"from":"20201014T000000","to":"20201014T010006"},"datasource":"vali1","queries":[{"expr":"{ip=\\"192.168.0.1\\", host=\\"host\\"}","refId":""}]}`
      );
    });
  });
});
