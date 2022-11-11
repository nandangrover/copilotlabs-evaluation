const  sortedSquares = require('./sorted_squares');
const assert = require('assert')

describe('TestSortedSquares', () => {
	it('TestSortedSquares', () => {
	  const input = [1, 2, 3, 5, 6, 8, 9];
	  const expected = [1, 4, 9, 25, 36, 64, 81];
	  const actual = sortedSquares(input);
	  expect(actual).toEqual(expected);
	});
  });
