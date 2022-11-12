package main

import (
	"fmt"
)

type SuffixTrie struct {
	root      map[string]*SuffixTrie
	endSymbol string
}

func NewSuffixTrie(str string) *SuffixTrie {
	root := make(map[string]*SuffixTrie)
	endSymbol := "*"
	suffixTrie := &SuffixTrie{root, endSymbol}
	suffixTrie.populateSuffixTrieFrom(str)
	return suffixTrie
}

// O(n^2) time | O(n^2) space
func (s *SuffixTrie) populateSuffixTrieFrom(str string) {
	for i := 0; i < len(str); i++ {
		s.insertSubstringStartingAt(i, str)
	}
}

func (s *SuffixTrie) insertSubstringStartingAt(i int, str string) {
	node := s.root
	for j := i; j < len(str); j++ {
		letter := string(str[j])
		if _, ok := node[letter]; !ok {
			node[letter] = &SuffixTrie{make(map[string]*SuffixTrie), s.endSymbol}
		}
		node = node[letter].root
	}
	node[s.endSymbol] = &SuffixTrie{make(map[string]*SuffixTrie), s.endSymbol}
}

// O(m) time | O(1) space
func (s *SuffixTrie) contains(str string) bool {
	node := s.root
	for _, letter := range str {
		if _, ok := node[string(letter)]; !ok {
			return false
		}
		node = node[string(letter)].root
	}
	if _, ok := node[s.endSymbol]; ok {
		return true
	}
	return false
}

func main() {
	trie := NewSuffixTrie("babc")
	fmt.Println(trie.contains("abc")) // true

	trie2 := NewSuffixTrie("babc")
	fmt.Println(trie2.contains("ab")) // false
}
