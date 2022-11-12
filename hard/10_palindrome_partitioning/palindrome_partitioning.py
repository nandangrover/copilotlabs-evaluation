def palindromePartitioningMinCuts(inputString):
    palindromes = []
    for i in range(len(inputString)):
        row = []
        for j in range(len(inputString)):
            if i == j:
                row.append(True)
            else:
                row.append(False)
        palindromes.append(row)
    for length in range(2, len(inputString) + 1):
        for i in range(len(inputString) - length + 1):
            j = i + length - 1
            if length == 2:
                palindromes[i][j] = inputString[i] == inputString[j]
            else:
                palindromes[i][j] = inputString[i] == inputString[j] and palindromes[i + 1][j - 1]

    cuts = [float('inf')] * len(inputString)
    for i in range(len(inputString)):
        if palindromes[0][i]:
            cuts[i] = 0
        else:
            cuts[i] = cuts[i - 1] + 1
            for j in range(1, i):
                if palindromes[j][i] and cuts[j - 1] + 1 < cuts[i]:
                    cuts[i] = cuts[j - 1] + 1
    return cuts[len(cuts) - 1]

print(palindromePartitioningMinCuts('noonabbad'))
print(palindromePartitioningMinCuts('ababbbabbababa'))