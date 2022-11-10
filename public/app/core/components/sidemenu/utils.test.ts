import { updateConfig } from '../../config';
import { getForcedLoginUrl } from './utils';

describe('getForcedLoginUrl', () => {
  it.each`
    appSubUrl          | url                    | expected
    ${''}              | ${'/whatever?a=1&b=2'} | ${'/whatever?a=1&b=2&forceLogin=true'}
    ${'/plutono'}      | ${'/whatever?a=1&b=2'} | ${'/plutono/whatever?a=1&b=2&forceLogin=true'}
    ${'/plutono/test'} | ${'/whatever?a=1&b=2'} | ${'/plutono/test/whatever?a=1&b=2&forceLogin=true'}
    ${'/plutono'}      | ${''}                  | ${'/plutono?forceLogin=true'}
    ${'/plutono'}      | ${'/whatever'}         | ${'/plutono/whatever?forceLogin=true'}
    ${'/plutono'}      | ${'/whatever/'}        | ${'/plutono/whatever/?forceLogin=true'}
  `(
    "when appUrl set to '$appUrl' and appSubUrl set to '$appSubUrl' then result should be '$expected'",
    ({ appSubUrl, url, expected }) => {
      updateConfig({
        appSubUrl,
      });

      const result = getForcedLoginUrl(url);

      expect(result).toBe(expected);
    }
  );
});
