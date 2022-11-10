import unittest
from is_leap_year import *

class TestIsLeapYear(unittest.TestCase):
    def testLeapYear(self):
        self.assertEqual(is_leap_year(2000), True)
        self.assertEqual(is_leap_year(2001), False)
        self.assertEqual(is_leap_year(2002), False)
        self.assertEqual(is_leap_year(2003), False)
        self.assertEqual(is_leap_year(2004), True)
        self.assertEqual(is_leap_year(2100), False)