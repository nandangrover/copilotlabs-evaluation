import unittest
import io
import sys
from inverted_star import *

class TestInvertedStar(unittest.TestCase):
    def testAlphabetTriangle(self):
        capturedOutput = io.StringIO()
        sys.stdout = capturedOutput
        invertedStar(5)
        sys.stdout = sys.__stdout__
        self.assertEqual(capturedOutput.getvalue().strip(), "*\n   **\n  ***\n ****\n*****")