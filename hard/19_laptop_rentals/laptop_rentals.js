function laptopRentals(times) {
  if (times.length === 0) {
    return 0;
  }

  let usedLaptops = 0;
  const startTimes = times.map((interval) => interval[0]).sort((a, b) => a - b);
  const endTimes = times.map((interval) => interval[1]).sort((a, b) => a - b);

  let startIterator = 0;
  let endIterator = 0;

  while (startIterator < times.length) {
    if (startTimes[startIterator] >= endTimes[endIterator]) {
      usedLaptops -= 1;
      endIterator += 1;
    }

    usedLaptops += 1;
    startIterator += 1;
  }

  return usedLaptops;
}
console.log(laptopRentals([[0, 2], [1, 4], [4, 6], [0, 4], [7, 8], [9, 11], [3, 10]]));

module.exports = laptopRentals;