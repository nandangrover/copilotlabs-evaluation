def maxSumIncreasingSubsequence(array):
    sequences = [None for x in array]
    sums = array[:]
    maxSumIndex = 0
    for i in range(len(array)):
        currentNum = array[i]
        for j in range(0, i):
            otherNum = array[j]
            if otherNum < currentNum and sums[j] + currentNum >= sums[i]:
                sums[i] = sums[j] + currentNum
                sequences[i] = j
        if sums[i] >= sums[maxSumIndex]:
            maxSumIndex = i
    return [sums[maxSumIndex], buildSequence(array, sequences, maxSumIndex)]

def buildSequence(array, sequences, index):
    sequence = []
    while index is not None:
        sequence.append(array[index])
        index = sequences[index]
    return list(reversed(sequence))

print(maxSumIncreasingSubsequence([10, 70, 20, 30, 50, 11, 30]))