import unittest
from palindrome_partitioning import palindromePartitioningMinCuts

class TestTraverseGraph(unittest.TestCase):
    def test_case_1(self):
        input = "noonabbad"
        expected = 2
        actual = palindromePartitioningMinCuts(input)
        self.assertEqual(actual, expected)
