function detectArbitrage(exchangeRates) {
  // To use exchange rates as edge weights, we must be able to add them.
  // Since log(a*b) = log(a) + log(b), we can convert all rates to
  // -log10(rate) to use them as edge weights.
  const logExchangeRates = convertToLogMatrix(exchangeRates);

  // A negative weight cycle indicates an arbitrage.
  return foundNegativeWeightCycle(logExchangeRates, 0);
}


// Runs the Bellman–Ford Algorithm to detect any negative-weight cycles.
function foundNegativeWeightCycle(graph, start) {
  const distancesFromStart = Array(graph.length).fill(Infinity);
  distancesFromStart[start] = 0;

  for (let i = 0; i < graph.length - 1; i++) {
    // If no update occurs, that means there's no negative cycle.
    if (!relaxEdgesAndUpdateDistances(graph, distancesFromStart)) {
      return false;
    }
  }

  return relaxEdgesAndUpdateDistances(graph, distancesFromStart);
}


// Returns `True` if any distance was updated
function relaxEdgesAndUpdateDistances(graph, distances) {
  let updated = false;
  for (let sourceIdx = 0; sourceIdx < graph.length; sourceIdx++) {
    for (let destinationIdx = 0; destinationIdx < graph[sourceIdx].length; destinationIdx++) {
      const edgeWeight = graph[sourceIdx][destinationIdx];
      const newDistanceToDestination = distances[sourceIdx] + edgeWeight;
      if (newDistanceToDestination < distances[destinationIdx]) {
        updated = true;
        distances[destinationIdx] = newDistanceToDestination;
      }
    }
  }
  
  return updated;
}


function convertToLogMatrix(matrix) {
  const newMatrix = [];
  for (let row = 0; row < matrix.length; row++) {
    newMatrix.push([]);
    for (let rate = 0; rate < matrix[row].length; rate++) {
      newMatrix[row].push(-Math.log10(matrix[row][rate]));
    }
  }

  return newMatrix;
}

console.log(detectArbitrage([[1.0, 0.8631, 0.5903], [1.1586, 1.0, 0.6849], [1.6939, 1.46, 1.0]])); // True

module.exports = detectArbitrage;