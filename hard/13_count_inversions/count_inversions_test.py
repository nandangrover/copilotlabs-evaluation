import unittest
from count_inversions import countInversions

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = [2, 3, 3, 1, 9, 5, 6]
        expected = 5
        actual = countInversions(input)
        self.assertEqual(actual, expected)
