import unittest
from smallest_difference import smallest_difference

class TestSmallestDifference(unittest.TestCase):

    def test_case_1(self):
        expected = [28, 26]
        output = smallest_difference([-1, 5, 10, 20, 28, 3], [26, 134, 135, 15, 17])
        self.assertEqual(output, expected)

    def test_case_2(self):
        expected = [25, 1005]
        output = smallest_difference([10, 0, 20, 25], [1005, 1006, 1014, 1032, 1031])
        self.assertEqual(output, expected)
