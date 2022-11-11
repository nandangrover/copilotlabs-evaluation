const  MinimumWaitingTime = require('./total_waiting_time');

describe('Test Total Waiting Time', () => {
	it('Test Case 1', () => {
		const queries = [3, 2, 1, 2, 6];
		const expected = 17;
		const actual = MinimumWaitingTime(queries);
		expect(actual).toEqual(expected);
	});

	it('Test Case 2', () => {
		const queries = [1, 4, 5];
		const expected = 6;
		const actual = MinimumWaitingTime(queries);
		expect(actual).toEqual(expected);
	});
});