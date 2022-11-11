import unittest
from total_waiting_time import minimum_waiting_time

class Test(unittest.TestCase):
	def test_total_waiting_time(self):
		queries = [3, 2, 1, 2, 6]
		expected = 17
		actual = minimum_waiting_time(queries)
		self.assertEqual(expected, actual)

		queries = [1, 4, 5]
		expected = 6
		actual = minimum_waiting_time(queries)
		self.assertEqual(expected, actual)

