import unittest
from underscorify_substring import underscorifySubstring

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(
                    underscorifySubstring("testthis is a testtest to see if testestest it works", "test"),
                    "_test_this is a _testtest_ to see if _testestest_ it works",
                )
