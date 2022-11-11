import unittest
from tournament_winner import tournamentWinner

class TestTournamentWinner(unittest.TestCase):
    def test_tournament_winner(self):
        competitions = [
            ["HTML", "C#"],
            ["C#", "Python"],
            ["Python", "HTML"],
        ]
        results = [0, 0, 1]
        expected = "Python"
        actual = tournamentWinner(competitions, results)
        self.assertEqual(expected, actual)
