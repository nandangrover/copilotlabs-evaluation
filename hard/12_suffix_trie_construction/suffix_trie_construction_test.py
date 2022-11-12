import unittest
from suffix_trie_construction import SuffixTrie

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        trie = SuffixTrie("babc")
        # self.assertEqual(trie.root, expected)
        self.assertEqual(trie.contains("abc"), True)
