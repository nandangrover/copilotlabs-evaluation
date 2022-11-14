from far_from_land import max_distance
import unittest
class TestTournamentWinner(unittest.TestCase):
    def test_max_distance(self):
        tcs = [
            {
                'grid': [
                    [1, 1, 1],
                    [1, 1, 1],
                    [1, 1, 1],
                ],
                'ans': -1,
            },
    
            {
                'grid': [
                    [1, 0, 1],
                    [0, 0, 0],
                    [1, 0, 1],
                ],
                'ans': 2,
            },
    
            {
                'grid': [
                    [1, 0, 0],
                    [0, 0, 0],
                    [0, 0, 0],
                ],
                'ans': 4,
            },
        ]
    
        for item in tcs:
            assert item['ans'] == max_distance(item['grid'])
