function multiStringSearch(bigString, smallStrings) {
  const trie = new Trie();
  for (const subString of smallStrings) {
    trie.insert(subString);
  }
  const containedStrings = {};
  for (let i = 0; i < bigString.length; i++) {
    findSmallStringsIn(bigString, i, trie, containedStrings);
  }
  return smallStrings.map((subString) => subString in containedStrings);
}

function findSmallStringsIn(subString, startIdx, trie, containedStrings) {
  let currentNode = trie.root;
  for (let i = startIdx; i < subString.length; i++) {
    const currentChar = subString[i];
    if (!(currentChar in currentNode)) break;
    currentNode = currentNode[currentChar];
    if (trie.endSymbol in currentNode)
      containedStrings[currentNode[trie.endSymbol]] = true;
  }
}

class Trie {
  constructor() {
    this.root = {};
    this.endSymbol = "*";
  }

  insert(subString) {
    let current = this.root;
    for (let i = 0; i < subString.length; i++) {
      if (!(subString[i] in current)) {
        current[subString[i]] = {};
      }
      current = current[subString[i]];
    }
    current[this.endSymbol] = subString;
  }
}

console.log(
  multiStringSearch("this is a big subString", [
    "this",
    "yo",
    "is",
    "a",
    "bigger",
    "subString",
    "kappa",
  ])
);

module.exports = multiStringSearch;
