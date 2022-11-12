package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	trie := NewSuffixTrie("babc")
	// Couldnt get this to work
	require.Equal(t, expected, trie.Root)
	require.Equal(t, true, trie.contains("abc"))
}
