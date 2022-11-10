import unittest
import io
import sys
from number_pyramid_pattern import pyramidPattern

class TestFactorial(unittest.TestCase):
    def test_pyramid_pattern(self):
        capturedOutput = io.StringIO()                  # Create StringIO object
        sys.stdout = capturedOutput                     #  and redirect stdout.
        pyramidPattern(5)                              # Call function.
        sys.stdout = sys.__stdout__                     # Reset redirect.
        self.assertEqual(capturedOutput.getvalue().strip(), "1\n   22\n  333\n 4444\n55555")
