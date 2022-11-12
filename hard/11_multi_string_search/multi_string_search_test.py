import unittest
from multi_string_search import multiStringSearch

class TestTraverseGraph(unittest.TestCase):
    def test_case_1(self):
        expected = [True, False, True, True, False, True, False]
        actual = multiStringSearch('this is a big string', ['this', 'yo', 'is', 'a', 'bigger', 'string', 'kappa'])
        self.assertEqual(actual, expected)
