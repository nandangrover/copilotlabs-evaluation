const  MaxSumIncreasingSubsequence = require('./increasing_subsequence');
const assert = require('assert')

test("Test Case #1", () => {
  const [outputSum, outputSequence] = MaxSumIncreasingSubsequence([10, 70, 20, 30, 50, 11, 30]);
  expect(outputSum).toEqual(110);
  expect(outputSequence).toEqual([10, 20, 30, 50]);
});
