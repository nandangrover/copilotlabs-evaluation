import unittest
import io
import sys
from alphabet_triangle import *

class TestHellowWorld(unittest.TestCase):
    def testAlphabetTriangle(self):
        capturedOutput = io.StringIO()
        sys.stdout = capturedOutput
        alphabetTriangle()
        sys.stdout = sys.__stdout__
        self.assertEqual(capturedOutput.getvalue().strip(), "A \nA B \nA B C \nA B C D \nA B C D E \nA B C D E F \nA B C D E F G")