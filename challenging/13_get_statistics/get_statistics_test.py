import unittest
from get_statistics import get_statistics

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        input = [2, 1, 3, 4, 4, 5, 6, 7]
        expected = {
            "mean": 4.0,
            "median": 4.0,
            "mode": 4.0,
            "sample_variance": 4.0,
            "sample_standard_deviation": 2.0,
            "mean_confidence_interval": [2.6141, 5.3859],
        }
        actual = get_statistics(input)
        self.assertEqual(self.round_output_values(actual), expected)

    def round_output_values(self, stats):
        if type(stats) is not dict:
        # Bad output; let tests fail.
            return stats

        new_output = {}
        for key in stats.keys():
            new_output[key] = stats[key]

            if type(stats[key]) is list:
                if len(stats[key]) >= 1:
                    new_output[key][0] = self.round_to_4(stats[key][0])
                if len(stats[key]) >= 2:
                    new_output[key][1] = self.round_to_4(stats[key][1])
            else:
                new_output[key] = self.round_to_4(stats[key])

        return new_output

    def round_to_4(self, number):
        if not self.is_number(number):
        # Bad output; let tests fail.
            return number

        return round(number, 4)

    def is_number(self, element):
        return type(element) in (int, float)
