import unittest
from traverse_graph import numberOfWaysToTraverseGraph

class TestTraverseGraph(unittest.TestCase):
    def test_case_1(self):
        width = 4
        height = 3
        expected = 10
        actual = numberOfWaysToTraverseGraph(width, height)
        self.assertEqual(actual, expected)
