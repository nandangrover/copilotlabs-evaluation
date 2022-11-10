import unittest
import io
import sys
from sum import *

class TestSum(unittest.TestCase):
    def testSum(self):
        self.assertEqual(add_numbers(1, 2), 3)
        self.assertEqual(add_numbers(1, -2), -1)
        self.assertEqual(add_numbers(0, 0), 0)
        self.assertEqual(add_numbers(-1, 2), 1)
        self.assertEqual(add_numbers(-1, -2), -3)