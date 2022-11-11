import unittest
from sorted_squares import sortedSquaredArray

class TestSortedSquares(unittest.TestCase):
	def test_sorted_squares(self):
		input = [1, 2, 3, 5, 6, 8, 9]
		expected = [1, 4, 9, 25, 36, 64, 81]
		actual = sortedSquaredArray(input)
		self.assertEqual(expected, actual)
