function twoEdgeConnectedGraph(edges) {
  if (edges.length === 0) return true;

  const arrivalTimes = new Array(edges.length).fill(-1);
  const startVertex = 0;

  if (getMinimumArrivalTimeOfAncestors(startVertex, -1, 0, arrivalTimes, edges) === -1) {
    return false;
  }

  return areAllVerticesVisited(arrivalTimes);
}

function areAllVerticesVisited(arrivalTimes) {
  for (const time of arrivalTimes) {
    if (time === -1) return false;
  }

  return true;
}

function getMinimumArrivalTimeOfAncestors(currentVertex, parent, currentTime, arrivalTimes, edges) {
  arrivalTimes[currentVertex] = currentTime;

  let minimumArrivalTime = currentTime;

  for (const destination of edges[currentVertex]) {
    if (arrivalTimes[destination] === -1) {
      minimumArrivalTime = Math.min(
        minimumArrivalTime,
        getMinimumArrivalTimeOfAncestors(destination, currentVertex, currentTime + 1, arrivalTimes, edges),
      );
    } else if (destination !== parent) {
      minimumArrivalTime = Math.min(minimumArrivalTime, arrivalTimes[destination]);
    }
  }

  // A bridge was detected, which means the graph isn't two-edge-connected.
  if (minimumArrivalTime === currentTime && parent !== -1) return -1;

  return minimumArrivalTime;
}

const input1 = [
  [1, 2, 5],
  [0, 2],
  [0, 1, 3],
  [2, 4, 5],
  [3, 5],
  [0, 3, 4],
];
const input2 = [
  [1, 2, 5],
  [0, 2],
  [0, 1, 3],
  [2, 4, 5],
  [3, 5],
  [0, 3, 4, 6],
  [5],
];
console.log(twoEdgeConnectedGraph(input1));
console.log(twoEdgeConnectedGraph(input2));

module.exports = twoEdgeConnectedGraph;
