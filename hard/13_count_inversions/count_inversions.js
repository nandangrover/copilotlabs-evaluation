const countInversions = (array) => countSubArrayInversions(array, 0, array.length);

const countSubArrayInversions = (array, start, end) => {
    if (end - start <= 1) {
        return 0;
    }

    const middle = start + Math.floor((end - start) / 2);
    const leftInversions = countSubArrayInversions(array, start, middle);
    const rightInversions = countSubArrayInversions(array, middle, end);
    const mergedArrayInversions = mergeSortAndCountInversions(array, start, middle, end);
    return leftInversions + rightInversions + mergedArrayInversions;
};

const mergeSortAndCountInversions = (array, start, middle, end) => {
    const sortedArray = [];
    let left = start;
    let right = middle;
    let inversions = 0;

    while (left < middle && right < end) {
        if (array[left] <= array[right]) {
            sortedArray.push(array[left]);
            left += 1;
        } else {
            inversions += middle - left;
            sortedArray.push(array[right]);
            right += 1;
        }
    }

    sortedArray.push(...array.slice(left, middle), ...array.slice(right, end));
    for (let i = 0; i < sortedArray.length; i += 1) {
        array[start + i] = sortedArray[i];
    }

    return inversions;
};

console.log(countInversions([8, 5, 2, 9, 5, 6, 3]));
console.log(countInversions([1, 2, 3, 4, 5, 6, 7, 8]));

module.exports = countInversions;