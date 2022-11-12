const laptopRentals = require('./laptop_rentals');

describe('TestProgram', () => {
  it('Test Case 1', () => {
    const input = [
      [0, 2],
      [1, 4],
      [4, 6],
      [0, 4],
      [7, 8],
      [9, 11],
      [3, 10],
    ];
    const expected = 3;
    const actual = laptopRentals(input);
    expect(actual).toEqual(expected);
  });
});
