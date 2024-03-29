// O(n) time | O(n) space - where n is the length of the input array
function sortedSquaredArray(array) {
	const sortedSquares = new Array(array.length).fill(0);
  
	let smallerValueIdx = 0;
	let largerValueIdx = array.length - 1;
  
	for (let idx = array.length - 1; idx >= 0; idx--) {
	  const smallerValue = array[smallerValueIdx];
	  const largerValue = array[largerValueIdx];
  
	  if (Math.abs(smallerValue) > Math.abs(largerValue)) {
		sortedSquares[idx] = smallerValue * smallerValue;
		smallerValueIdx += 1;
	  } else {
		sortedSquares[idx] = largerValue * largerValue;
		largerValueIdx -= 1;
	  }
	}
  
	return sortedSquares;
  }
  
  const array = [-5, -4, 1, 2, 3, 5];
  console.log(sortedSquaredArray(array));

  module.exports = sortedSquaredArray;