import unittest
from knuth_morris_pratt_algorithm import knuthMorrisPrattAlgorithm

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(knuthMorrisPrattAlgorithm("aefoaefcdaefcdaed", "aefcdaed"), True)
