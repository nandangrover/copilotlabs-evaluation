function smallestDifference(array1, array2) {
	array1.sort((a, b) => a - b);
	array2.sort((a, b) => a - b);
	let idxOne = 0;
	let idxTwo = 0;
	let smallest = Number.MAX_SAFE_INTEGER;
	let current = Number.MAX_SAFE_INTEGER;
	let smallestPair = [];
	while (idxOne < array1.length && idxTwo < array2.length) {
		const first = array1[idxOne];
		const second = array2[idxTwo];
		if (first < second) {
			current = second - first;
			idxOne += 1;
		} else if (second < first) {
			current = first - second;
			idxTwo += 1;
		} else {
			return [first, second];
		}
		if (smallest > current) {
			smallest = current;
			smallestPair = [first, second];
		}
	}
	return smallestPair;
}

const array1 = [-1, 5, 10, 20, 28, 3];
const array2 = [26, 134, 135, 15, 17];
console.log(smallestDifference(array1, array2));

module.exports = smallestDifference;
