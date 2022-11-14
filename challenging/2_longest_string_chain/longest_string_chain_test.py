import unittest
from longest_string_chain import longestStringChain

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        strings = ["abde", "abc", "abd", "abcde", "ade", "ae", "1abde", "abcdef"]
        expected = ["abcdef", "abcde", "abde", "ade", "ae"]
        self.assertEqual(longestStringChain(strings), expected)
