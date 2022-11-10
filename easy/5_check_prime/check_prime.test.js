const assert = require('assert');
const is_prime = require('./check_prime.js');

describe('check prime', () =>{
    it('should check if number is prime', () => {
        assert.equal(is_prime(1), false);
        assert.equal(is_prime(17), true);
        assert.equal(is_prime(2), true);
        assert.equal(is_prime(3), true);
    });
});