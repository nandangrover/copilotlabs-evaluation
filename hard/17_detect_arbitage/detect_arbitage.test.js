// import { underscorifySubstring } from './underscorify_substring.js';
const detectArbitrage = require('./detect_arbitage.js');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    const input = [[1.0, 0.8631, 0.5903], [1.1586, 1.0, 0.6849], [1.6939, 1.46, 1.0]];
    const expected = true;
    const actual = detectArbitrage(input);
    expect(actual).toEqual(expected);
  });
});