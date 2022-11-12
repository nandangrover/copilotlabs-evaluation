const countInversions = require('./count_inversions');

test('Test Case #1', () => {
  const input = [2, 3, 3, 1, 9, 5, 6];
  const expected = 5;
  const actual = countInversions(input);
  expect(actual).toEqual(expected);
});