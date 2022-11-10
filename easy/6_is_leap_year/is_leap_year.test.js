const assert = require('assert');
const is_leap_year = require('./is_leap_year.js');

describe('is_leap_year', () => {
    it('should return True for 2000', () => {
        assert.equal(is_leap_year(2000), true);
    });
    it('should return False for 2001', () => {
        assert.equal(is_leap_year(2001), false);
    });
    it('should return False for 2002', () => {
        assert.equal(is_leap_year(2002), false);
    });
    it('should return False for 2003', () => {
        assert.equal(is_leap_year(2003), false);
    });
    it('should return True for 2004', () => {
        assert.equal(is_leap_year(2004), true);
    });
    it('should return False for 2100', () => {
        assert.equal(is_leap_year(2100), false);
    });
});