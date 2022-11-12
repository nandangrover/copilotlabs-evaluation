import math

# O(n^3) time | O(n^2) space - where n is the number of currencies
def detectArbitrage(exchangeRates):
    # To use exchange rates as edge weights, we must be able to add them.
    # Since log(a*b) = log(a) + log(b), we can convert all rates to
    # -log10(rate) to use them as edge weights.
    logExchangeRates = convertToLogMatrix(exchangeRates)

    # A negative weight cycle indicates an arbitrage.
    return foundNegativeWeightCycle(logExchangeRates, 0)


# Runs the Bellman–Ford Algorithm to detect any negative-weight cycles.
def foundNegativeWeightCycle(graph, start):
    distancesFromStart = [float("inf") for _ in range(len(graph))]
    distancesFromStart[start] = 0

    for _ in range(len(graph) - 1):
        # If no update occurs, that means there's no negative cycle.
        if not relaxEdgesAndUpdateDistances(graph, distancesFromStart):
            return False

    return relaxEdgesAndUpdateDistances(graph, distancesFromStart)


# Returns `True` if any distance was updated
def relaxEdgesAndUpdateDistances(graph, distances):
    updated = False
    for sourceIdx, edges in enumerate(graph):
        for destinationIdx, edgeWeight in enumerate(edges):
            newDistanceToDestination = distances[sourceIdx] + edgeWeight
            if newDistanceToDestination < distances[destinationIdx]:
                updated = True
                distances[destinationIdx] = newDistanceToDestination

    return updated


def convertToLogMatrix(matrix):
    newMatrix = []
    for row, rates in enumerate(matrix):
        newMatrix.append([])
        for rate in rates:
            newMatrix[row].append(-math.log10(rate))

    return newMatrix

print(detectArbitrage([[1.0, 0.8631, 0.5903], [1.1586, 1.0, 0.6849], [1.6939, 1.46, 1.0]])) # True
