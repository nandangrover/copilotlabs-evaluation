import unittest
from non_attacking_queens import nonAttackingQueens

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = 4
        expected = 2
        actual = nonAttackingQueens(input)
        self.assertEqual(actual, expected)
