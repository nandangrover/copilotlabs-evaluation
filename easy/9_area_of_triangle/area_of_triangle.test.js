const assert = require('assert');
const areaOfTriangle = require('./area_of_triangle.js');

describe('Area of triangle', () => {
    it('should return area of triangle', () => {
        assert.equal(areaOfTriangle(2, 3), 3);
        assert.equal(areaOfTriangle(3, 4), 6);
        assert.equal(areaOfTriangle(4, 5), 10);
        assert.equal(areaOfTriangle(5, 6), 15);
    });
});