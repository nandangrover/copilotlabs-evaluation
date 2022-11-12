import unittest
from laptop_rentals import laptopRentals

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = [[0, 2], [1, 4], [4, 6], [0, 4], [7, 8], [9, 11], [3, 10]]
        expected = 3
        actual = laptopRentals(input)
        self.assertEqual(actual, expected)

