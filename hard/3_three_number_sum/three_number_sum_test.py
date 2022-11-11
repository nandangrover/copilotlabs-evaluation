import unittest
from three_number_sum import three_number_sum

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(
            three_number_sum([12, 3, 1, 2, -6, 5, -8, 6], 0),
            [[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]],
        )

    def test_case_2(self):
        self.assertEqual(
            three_number_sum([1, 2, 3], 6), [[1, 2, 3]]
        )
