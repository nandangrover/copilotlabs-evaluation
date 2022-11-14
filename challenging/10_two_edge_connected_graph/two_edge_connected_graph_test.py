from two_edge_connected_graph import two_edge_connected_graph
import unittest


class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = [[1, 2, 5], [0, 2], [0, 1, 3], [2, 4, 5], [3, 5], [0, 3, 4]]
        expected = True
        actual = two_edge_connected_graph(input)
        self.assertEqual(actual, expected)
