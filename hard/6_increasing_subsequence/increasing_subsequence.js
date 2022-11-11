const MaxSumIncreasingSubsequence = (array) => {
	const sequences = new Array(array.length);
	const sums = array.slice();
	let maxSumIndex = 0;
	for (let i = 0; i < array.length; i++) {
		const currentNum = array[i];
		for (let j = 0; j < i; j++) {
			const otherNum = array[j];
			if (otherNum < currentNum && sums[j] + currentNum >= sums[i]) {
				sums[i] = sums[j] + currentNum;
				sequences[i] = j;
			}
		}
		if (sums[i] > sums[maxSumIndex]) {
			maxSumIndex = i;
		}
	}
	const maxSum = sums[maxSumIndex];
	const sequence = buildSequence(array, sequences, maxSumIndex);
	return [maxSum, sequence];
}

function buildSequence(array, sequences, index) {
	const sequence = [];
	while (index !== undefined) {
		sequence.push(array[index]);
		index = sequences[index];
	}
	return sequence.reverse();
}

const array = [10, 70, 20, 30, 50, 11, 30];
const [maxSum, sequence] = MaxSumIncreasingSubsequence(array);
console.log(maxSum);
console.log(sequence);

module.exports = MaxSumIncreasingSubsequence;