const assert = require('assert');
const add_numbers = require('./sum.js');

describe('sum', () =>{
    it('should add two numbers', () => {
        assert.equal(add_numbers(1, 2), 3);
        assert.equal(add_numbers(1, -2), -1);
        assert.equal(add_numbers(0, 0), 0);
        assert.equal(add_numbers(-1, 2), 1);
        assert.equal(add_numbers(-1, -2), -3);
    });
});