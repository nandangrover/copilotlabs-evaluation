import unittest
from area_of_triangle import areaOfTriangle

class TestAreaOfTriangle(unittest.TestCase):
    def test_area_of_triangle(self):
        self.assertEqual(areaOfTriangle(2, 3), 3)
        self.assertEqual(areaOfTriangle(3, 4), 6)
        self.assertEqual(areaOfTriangle(4, 5), 10)
        self.assertEqual(areaOfTriangle(5, 6), 15)
