const knapsackProblem = require('./knapsack_problem');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    expect(knapsackProblem([[1, 2], [4, 3], [5, 6], [6, 7]], 10)).toEqual([10, [1, 3]]);
  });
});
