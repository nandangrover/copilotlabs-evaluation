const  SmallestDifference = require('./smallest_difference');
const assert = require('assert')

describe('TestSmallestDifference', function () {
	it('TestSmallestDifference1', function () {
		const expected = [28, 26]
		const output = SmallestDifference([-1, 5, 10, 20, 28, 3], [26, 134, 135, 15, 17])
		assert.deepEqual(output, expected)
	})
	it('TestSmallestDifference2', function () {
		const expected = [25, 1005]
		const output = SmallestDifference([10, 0, 20, 25], [1005, 1006, 1014, 1032, 1031])
		assert.deepEqual(output, expected)
	})
})
