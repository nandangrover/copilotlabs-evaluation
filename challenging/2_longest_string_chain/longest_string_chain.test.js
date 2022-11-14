const longestStringChain = require('./longest_string_chain');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    const strings = ['abde', 'abc', 'abd', 'abcde', 'ade', 'ae', '1abde', 'abcdef'];
    const expected = ['abcdef', 'abcde', 'abde', 'ade', 'ae'];
    expect(longestStringChain(strings)).toEqual(expected);
  });
});
