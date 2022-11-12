package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiStringSearch("this is a big subString", []string{"this", "yo", "is", "a", "bigger", "subString", "kappa"}))
}

func multiStringSearch(bigString string, smallStrings []string) []bool {
	trie := NewTrie()
	for _, subString := range smallStrings {
		trie.Insert(subString)
	}
	containedStrings := make(map[string]bool)
	for i := 0; i < len(bigString); i++ {
		findSmallStringsIn(bigString, i, trie, containedStrings)
	}
	results := make([]bool, len(smallStrings))
	for i, subString := range smallStrings {
		_, exists := containedStrings[subString]
		results[i] = exists
	}
	return results
}

func findSmallStringsIn(subString string, startIdx int, trie *Trie, containedStrings map[string]bool) {
	currentNode := trie.Root
	for i := startIdx; i < len(subString); i++ {
		currentChar := subString[i]
		if _, exists := currentNode[currentChar]; !exists {
			break
		}
		currentNode = currentNode[currentChar]
		if _, exists := currentNode[trie.EndSymbol]; exists {
			containedStrings[currentNode[trie.EndSymbol]] = true
		}
	}
}

type Trie struct {
	Root      map[rune]map[rune]string
	EndSymbol rune
}

func NewTrie() *Trie {
	trie := &Trie{
		Root:      make(map[rune]map[rune]string),
		EndSymbol: '*',
	}
	return trie
}

func (t *Trie) Insert(subString string) {
	current := t.Root
	for _, char := range subString {
		if _, exists := current[char]; !exists {
			current[char] = make(map[rune]string)
		}
		current = current[char]
	}
	current[t.EndSymbol] = subString
}
