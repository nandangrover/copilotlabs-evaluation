import unittest
from check_prime import *

class TestPrime(unittest.TestCase):
    def testPrime(self):
        self.assertEqual(is_prime(1), False)
        self.assertEqual(is_prime(17), True)
        self.assertEqual(is_prime(2), True)
        self.assertEqual(is_prime(3), True)