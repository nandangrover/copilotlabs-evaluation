import unittest
from knapsack_problem import knapsackProblem

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(knapsackProblem([[1, 2], [4, 3], [5, 6], [6, 7]], 10), [10, [1, 3]])
