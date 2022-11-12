import unittest
from longest_balanced_substring import LongestBalancedSubstring

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        string = "(()))("
        expected = 4
        actual = LongestBalancedSubstring(string)
        self.assertEqual(actual, expected)
