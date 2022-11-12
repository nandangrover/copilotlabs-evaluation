import unittest
from detect_arbitage import detectArbitrage

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = [[1.0, 0.8631, 0.5903], [1.1586, 1.0, 0.6849], [1.6939, 1.46, 1.0]]
        expected = True
        actual = detectArbitrage(input)
        self.assertEqual(actual, expected)
