const nonAttackingQueens =  require('./non_attacking_queens.js');

describe('nonAttackingQueens', () => {
  it('should return the number of non attacking queens', () => {
    const input = 4;
    const expected = 2;
    const actual = nonAttackingQueens(input);
    expect(actual).toEqual(expected);
  });
});