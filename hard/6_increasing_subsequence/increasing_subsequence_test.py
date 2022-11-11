import unittest
from increasing_subsequence import maxSumIncreasingSubsequence

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        outputSum, outputSequence = maxSumIncreasingSubsequence([10, 70, 20, 30, 50, 11, 30])
        self.assertEqual(outputSum, 110)
        self.assertEqual(outputSequence, [10, 20, 30, 50])
