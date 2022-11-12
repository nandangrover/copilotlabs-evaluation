const getStatistics = require('./get_statistics');

describe("getStatistics", () => {
  const testCases = [
    {
      input: [2, 1, 3, 4, 4, 5, 6, 7],
      expected: {
        mean: 4,
        median: 4,
        mode: 4,
        sampleVariance: 4,
        sampleStandardDeviation: 2,
        meanConfidenceInterval: [2.6141, 5.3859],
      },
    },
  ];

  testCases.forEach(({ input, expected }) => {
    it(`should return ${JSON.stringify(
      expected
    )} for input ${JSON.stringify(input)}`, () => {
      expect(getStatistics(input)).toEqual(expected);
    });
  });
});