def laptopRentals(times):
    if len(times) == 0:
        return 0

    usedLaptops = 0
    startTimes = sorted([interval[0] for interval in times])
    endTimes = sorted([interval[1] for interval in times])

    startIterator = 0
    endIterator = 0

    while startIterator < len(times):
        if startTimes[startIterator] >= endTimes[endIterator]:
            usedLaptops -= 1
            endIterator += 1

        usedLaptops += 1
        startIterator += 1

    return usedLaptops
print(laptopRentals([[0, 2], [1, 4], [4, 6], [0, 4], [7, 8], [9, 11], [3, 10]]))