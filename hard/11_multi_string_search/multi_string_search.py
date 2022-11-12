def multiStringSearch(bigString, smallStrings):
    trie = Trie()
    for subString in smallStrings:
        trie.insert(subString)
    containedStrings = {}
    for i in range(0, len(bigString)):
        findSmallStringsIn(bigString, i, trie, containedStrings)
    return [subString in containedStrings for subString in smallStrings]

def findSmallStringsIn(subString, startIdx, trie, containedStrings):
    currentNode = trie.root
    for i in range(startIdx, len(subString)):
        currentChar = subString[i]
        if currentChar not in currentNode:
            break
        currentNode = currentNode[currentChar]
        if trie.endSymbol in currentNode:
            containedStrings[currentNode[trie.endSymbol]] = True

class Trie:
    def __init__(self):
        self.root = {}
        self.endSymbol = "*"

    def insert(self, subString):
        current = self.root
        for i in range(0, len(subString)):
            if subString[i] not in current:
                current[subString[i]] = {}
            current = current[subString[i]]
        current[self.endSymbol] = subString

print(multiStringSearch("this is a big subString", ["this", "yo", "is", "a", "bigger", "subString", "kappa"]))