const  threeNumberSum = require('./three_number_sum');

describe("Three Number Sum", () => {
	it("Test Case #1", () => {
		const queries = [12, 3, 1, 2, -6, 5, -8, 6];
		const target = 0;
		const expected = [[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]];

		const actual = threeNumberSum(queries, target);
		expect(actual).toEqual(expected);
	});

	it("Test Case #2", () => {
		const queries = [1, 2, 3];
		const target = 6;
		const expected = [[1, 2, 3]];

		const actual = threeNumberSum(queries, target);
		expect(actual).toEqual(expected);
	});
});