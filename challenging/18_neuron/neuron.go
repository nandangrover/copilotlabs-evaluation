package main

import (
	"fmt"
	"math"
)

type Neuron struct {
	weights   []float64
	Examples  []Example
	Gradients []float64
}

func (n *Neuron) Train(learningRate float64, batchSize int, epochs int) {
	for i := 0; i < epochs; i++ {
		for j := 0; j < len(n.Examples)/batchSize; j++ {
			miniBatch := n.Examples[0+(batchSize*j) : batchSize+(batchSize*j)]
			predictionsLabels := []PredictionLabel{}
			for _, example := range miniBatch {
				predictionsLabels = append(predictionsLabels, PredictionLabel{
					Prediction: n.Predict(example.Features),
					Label:      example.Label,
				})
			}
			n.Gradients = n.GetGradients(miniBatch, predictionsLabels)
			for k, gradient := range n.Gradients {
				n.weights[k] = n.weights[k] - (learningRate * gradient)
			}
		}
	}
}

func (n *Neuron) Predict(features []float64) float64 {
	modelInputs := append(features, 1.0)
	wTx := 0.0
	for i, modelInput := range modelInputs {
		wTx += n.weights[i] * modelInput
	}
	return 1.0 / (1.0 + math.Exp(-wTx))
}

func (n *Neuron) GetGradients(batch []Example, predictionsLabels []PredictionLabel) []float64 {
	errors := []float64{}
	for _, predictionLabel := range predictionsLabels {
		errors = append(errors, predictionLabel.Prediction-predictionLabel.Label)
	}
	gradients := []float64{0.0, 0.0, 0.0, 0.0}
	for exampleI, example := range batch {
		features := append(example.Features, 1.0)
		for featureI, feature := range features {
			gradients[featureI] += errors[exampleI] * feature
		}
	}
	for i, gradient := range gradients {
		gradients[i] = gradient / float64(len(batch))
	}
	return gradients
}

type PredictionLabel struct {
	Prediction float64
	Label      float64
}

type Example struct {
	Features []float64
	Label    float64
}

func main() {
	examples := []Example{
		Example{
			Features: []float64{0.0, 0.0},
			Label:    0.0,
		},
		Example{
			Features: []float64{0.0, 1.0},
			Label:    1.0,
		},
		Example{
			Features: []float64{1.0, 0.0},
			Label:    1.0,
		},
		Example{
			Features: []float64{1.0, 1.0},
			Label:    0.0,
		},
	}
	neuron := Neuron{
		weights:  []float64{0.49671415, 0.64768854, 0.52302986, 0.64589411},
		Examples: examples,
	}
	neuron.Train(0.01, 10, 200)
	fmt.Println(neuron.Predict([]float64{0.79, 0.89, 0.777}))
}
