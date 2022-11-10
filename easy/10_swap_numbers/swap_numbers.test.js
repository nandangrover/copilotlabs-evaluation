const assert = require('assert');
const swap_numbers = require('./swap_numbers.js');

describe('Swap Numbers', () => {
    it('should return swapped numbers', () => {
        assert.deepEqual(swap_numbers(1, 2), [2, 1]);
        assert.deepEqual(swap_numbers(1, -2), [-2, 1]);
        assert.deepEqual(swap_numbers(0, 0), [0, 0]);
        assert.deepEqual(swap_numbers(-1, 2), [2, -1]);
    });
});
