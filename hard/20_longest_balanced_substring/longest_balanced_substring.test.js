const longestBalancedSubstring = require('./longest_balanced_substring.js');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    const string = '(()))(';
    const expected = 4;
    const actual = longestBalancedSubstring(string);
    expect(actual).toEqual(expected);
  });
});
