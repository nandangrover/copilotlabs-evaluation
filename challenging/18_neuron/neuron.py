import numpy as np
import math


class Neuron:
    def __init__(self, examples):
        np.random.seed(42)
        # Three weights: one for each feature and one more for the bias.
        self.weights = np.random.normal(0, 1, 3 + 1)
        self.examples = examples
        self.train()

    def train(self, learning_rate=0.01, batch_size=10, epochs=200):
        for _ in range(epochs):
            for batch_window in range(len(self.examples) // batch_size):
                mini_batch = self.examples[0 + (batch_size * batch_window) : batch_size + (batch_size * batch_window)]
                predictions_labels = [
                    {"prediction": self.predict(example["features"]), "label": example["label"]}
                    for example in mini_batch
                ]
                gradients = self.__get_gradients(mini_batch, predictions_labels)
                self.weights = self.weights - [learning_rate * gradient for gradient in gradients]

    # Return the probability, not the corresponding 0 or 1 label.
    def predict(self, features):
        model_inputs = features + [1]
        wTx = 0
        for i, model_input in enumerate(model_inputs):
            wTx = wTx + self.weights[i] * model_input
        return 1 / (1 + math.exp(-wTx))

    def __get_gradients(self, batch, predictions_labels):
        errors = [
            predictions_label["prediction"] - predictions_label["label"] for predictions_label in predictions_labels
        ]
        gradients = [0] * len(self.weights)
        for example_i, example in enumerate(batch):
            features = example["features"] + [1]
            for feature_i, feature in enumerate(features):
                gradients[feature_i] = gradients[feature_i] + errors[example_i] * feature
        gradients = [gradient / len(batch) for gradient in gradients]
        return gradients

examples = [
		{"features": [0, 0], "label": 0.0},
		{"features": [0, 1], "label": 1.0},
		{"features": [1, 0], "label": 1.0},
		{"features": [1, 1], "label": 0.0},
	]
neuron = Neuron(examples)
print(neuron.predict([0.79, 0.89, 0.777]))
