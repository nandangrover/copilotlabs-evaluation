def longestStringChain(strings):
    # For every string, imagine the longest string chain that starts with it.
    # Set up every string to point to the next string in its respective longest
    # string chain. Also keep track of the lengths of these longest string chains.
    stringChains = {string: {"nextString": "", "maxChainLength": 1} for string in strings}

    # Sort the strings based on their length so that whenever we visit a
    # string (as we iterate through them from left to right), we can
    # already have computed the longest string chains of any smaller strings.
    sortedStrings = sorted(strings, key=len)

    for string in sortedStrings:
        findLongestStringChain(string, stringChains)
    return buildLongestStringChain(strings, stringChains)

def findLongestStringChain(string, stringChains):
    # Try removing every letter of the current string to see if the
    # remaining strings form a string chain.
    for i in range(len(string)):
        smallerString = getSmallerString(string, i)
        if smallerString not in stringChains:
            continue
        tryUpdateLongestStringChain(string, smallerString, stringChains)

def getSmallerString(string, index):
    return string[:index] + string[index+1:]

def tryUpdateLongestStringChain(currentString, smallerString, stringChains):
    smallerStringChainLength = stringChains[smallerString]["maxChainLength"]
    currentStringChainLength = stringChains[currentString]["maxChainLength"]
    # Update the string chain of the current string only if the smaller string leads
    # to a longer string chain.
    if smallerStringChainLength+1 > currentStringChainLength:
        stringChains[currentString]["maxChainLength"] = smallerStringChainLength + 1
        stringChains[currentString]["nextString"] = smallerString

def buildLongestStringChain(strings, stringChains):
    # Find the string that starts the longest string chain.
    maxChainLength = 0
    chainStartingString = ""
    for string in strings:
        if stringChains[string]["maxChainLength"] > maxChainLength:
            maxChainLength = stringChains[string]["maxChainLength"]
            chainStartingString = string

    # Starting at the string found above, build the longest string chain.
    ourLongestStringChain = []
    currentString = chainStartingString
    while currentString != "":
        ourLongestStringChain.append(currentString)
        currentString = stringChains[currentString]["nextString"]
    if len(ourLongestStringChain) == 1:
        return []
    return ourLongestStringChain


strings = ["abde", "abc", "abd", "abcde", "ade", "ae", "1abde", "abcdef"]
print(longestStringChain(strings))