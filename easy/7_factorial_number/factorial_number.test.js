const assert = require('assert');
const factorial = require('./factorial_number.js');

describe('factorial number', () => {
    it('should return factorial number', () => {
        assert.equal(factorial(1), 1);
        assert.equal(factorial(2), 2);
        assert.equal(factorial(3), 6);
        assert.equal(factorial(4), 24);
        assert.equal(factorial(5), 120);
        assert.equal(factorial(6), 720);
    });
});