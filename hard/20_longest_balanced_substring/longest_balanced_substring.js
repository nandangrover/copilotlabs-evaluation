const LongestBalancedSubstring = (str) => {
	return Math.max(
		getLongestBalancedInDirection(str, true),
		getLongestBalancedInDirection(str, false),
	);
};

const getLongestBalancedInDirection = (str, leftToRight) => {
	let openingParens = "(";
	let startIdx = 0;
	let step = 1;
	if (!leftToRight) {
		openingParens = ")";
		startIdx = str.length - 1;
		step = -1;
	}

	let maxLength = 0;

	let openingCount = 0;
	let closingCount = 0;

	let idx = startIdx;
	while (idx >= 0 && idx < str.length) {
		let char = str[idx];

		if (char === openingParens) {
			openingCount++;
		} else {
			closingCount++;
		}
		if (openingCount === closingCount) {
			maxLength = Math.max(maxLength, closingCount * 2);
		} else if (closingCount > openingCount) {
			openingCount = 0;
			closingCount = 0;
		}

		idx += step;
	}

	return maxLength;
};

console.log(LongestBalancedSubstring("(()))(")); // 4

module.exports = LongestBalancedSubstring;
