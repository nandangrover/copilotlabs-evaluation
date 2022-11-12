def numberOfWaysToTraverseGraph(width, height):
    numberOfWays = []
    for i in range(height + 1):
        numberOfWays.append([])
        for j in range(width + 1):
            numberOfWays[i].append(0)

    for widthIdx in range(1, width + 1):
        for heightIdx in range(1, height + 1):
            if widthIdx == 1 or heightIdx == 1:
                numberOfWays[heightIdx][widthIdx] = 1
            else:
                waysLeft = numberOfWays[heightIdx][widthIdx - 1]
                waysUp = numberOfWays[heightIdx - 1][widthIdx]
                numberOfWays[heightIdx][widthIdx] = waysLeft + waysUp

    return numberOfWays[height][width]


width = 4
height = 3
print(numberOfWaysToTraverseGraph(width, height))
