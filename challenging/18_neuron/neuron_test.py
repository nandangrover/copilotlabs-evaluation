import unittest
from neuron import Neuron
import numpy as np

class TestTournamentWinner(unittest.TestCase):
    def test_tournament_winner(self):
        examples = [
        		{"features": [0, 0], "label": 0.0},
        		{"features": [0, 1], "label": 1.0},
        		{"features": [1, 0], "label": 1.0},
        		{"features": [1, 1], "label": 0.0},
        	]
        neuron = Neuron(examples)
        actual = np.round(neuron.predict([0.79, 0.89, 0.777]))
        expected = 1.0
        self.assertEqual(actual, expected)
