function MinimumWaitingTime(queries) {
	queries.sort((a, b) => a - b);

	let totalWaitingTime = 0;
	for (let idx = 0; idx < queries.length; idx++) {
		const duration = queries[idx];
		const queriesLeft = queries.length - (idx + 1);
		totalWaitingTime += duration * queriesLeft;
	}

	return totalWaitingTime;
}

const queries = [3, 2, 1, 2, 6];
console.log(MinimumWaitingTime(queries));

module.exports = MinimumWaitingTime;
