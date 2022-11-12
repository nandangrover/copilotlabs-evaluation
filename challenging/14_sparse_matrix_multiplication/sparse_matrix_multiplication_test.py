from sparse_matrix_multiplication import sparse_matrix_multiplication
import unittest

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        matrix_a = [
            [0, 2, 0],
            [0, -3, 5],
        ]
        matrix_b = [
            [0, 10, 0],
            [0, 0, 0],
            [0, 0, 4],
        ]
        expected = [
            [0, 0, 0],
            [0, 0, 20],
        ]
        actual = sparse_matrix_multiplication(matrix_a, matrix_b)
        self.assertEqual(actual, expected)
