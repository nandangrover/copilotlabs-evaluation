import unittest
import io
import sys
from hello_world import *

class TestHellowWorld(unittest.TestCase):
    def testHelloWorld(self):
        capturedOutput = io.StringIO()
        sys.stdout = capturedOutput
        helloworld()
        sys.stdout = sys.__stdout__
        self.assertEqual(capturedOutput.getvalue().strip(), "Hello World")