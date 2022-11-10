import unittest
from swap_numbers import swap_numbers

class TestSwapNumbers(unittest.TestCase):
    def test_swap_numbers(self):
        self.assertEqual(swap_numbers(1, 2), [2, 1])
        self.assertEqual(swap_numbers(1, -2), [-2, 1])
        self.assertEqual(swap_numbers(0, 0), [0, 0])
        self.assertEqual(swap_numbers(-1, 2), [2, -1])
