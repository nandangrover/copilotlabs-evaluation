from pandas_math import pandas_operation
import unittest
import pandas as pd

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        ds1 = pd.Series([2, 4, 6, 8, 10])
        ds2 = pd.Series([1, 3, 5, 7, 9])
        actual = pandas_operation(ds1, ds2)
        self.assertEqual([[3, 7, 11, 15, 19], [1, 1, 1, 1, 1], [2, 12, 30, 56, 90], [2.0, 1.3333333333333333, 1.2, 1.1428571428571428, 1.1111111111111112]], actual)
