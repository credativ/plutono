import { removeEmpty } from './utils';

describe('removeEmpty', () => {
  it('Should remove all empty', () => {
    const original = {
      stringsShouldBeKept: 'Something',
      unlessTheyAreEmpty: '',
      nullToBeRemoved: null,
      undefinedToBeRemoved: null,
      zeroShouldBeKept: 0,
      booleansShouldBeKept: false,
      emptyObjectsShouldBeRemoved: {},
      emptyArrayShouldBeRemoved: [],
      nonEmptyArraysShouldBeKept: [1, 2, 3],
      nestedObjToBeRemoved: {
        toBeRemoved: undefined,
      },
      nestedObjectToKeep: {
        thisShouldBeRemoved: null,
        thisShouldBeKept: 'Hello, Plutono',
      },
    };

    const expectedResult = {
      stringsShouldBeKept: 'Something',
      zeroShouldBeKept: 0,
      booleansShouldBeKept: false,
      nonEmptyArraysShouldBeKept: [1, 2, 3],
      nestedObjectToKeep: {
        thisShouldBeKept: 'Hello, Plutono',
      },
    };

    expect(removeEmpty(original)).toStrictEqual(expectedResult);
  });
});
