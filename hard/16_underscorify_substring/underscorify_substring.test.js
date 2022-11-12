// import { underscorifySubstring } from './underscorify_substring.js';
const underscorifySubstring = require('./underscorify_substring.js');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    expect(underscorifySubstring("testthis is a testtest to see if testestest it works", "test")).toEqual("_test_this is a _testtest_ to see if _testestest_ it works");
  });
});