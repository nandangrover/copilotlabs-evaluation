const  palindromePartitioning = require('./palindrome_partitioning');

test("Test Case #1", () => {
  const input = 'noonabbad';
  const expected = 2;
  const actual = palindromePartitioning(input);
  expect(actual).toEqual(expected);
});
