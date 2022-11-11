const threeNumberSum = (array, target) => {
	array.sort((a, b) => a - b);
	const triplets = [];
	for (let i = 0; i < array.length - 2; i++) {
		let left = i + 1;
		let right = array.length - 1;
		while (left < right) {
			const currentSum = array[i] + array[left] + array[right];
			if (currentSum === target) {
				triplets.push([array[i], array[left], array[right]]);
				left += 1;
				right -= 1;
			} else if (currentSum < target) {
				left += 1;
			} else if (currentSum > target) {
				right -= 1;
			}
		}
	}
	return triplets;
};

const array = [12, 3, 1, 2, -6, 5, -8, 6];
const target = 0;
console.log(threeNumberSum(array, target));

module.exports = threeNumberSum;
