package main

import (
	"fmt"
	"math"
)

type MultinomialNB struct {
	alpha                   float64
	priorsPerTag            map[string]float64
	likelihoodPerWordPerTag map[string]map[string]float64
	articlesPerTag          map[string][][]string
	tags                    []string
}

func NewMultinomialNB(articlesPerTag map[string][][]string) *MultinomialNB {
	return &MultinomialNB{
		alpha:                   1,
		priorsPerTag:            make(map[string]float64),
		likelihoodPerWordPerTag: make(map[string]map[string]float64),
		articlesPerTag:          articlesPerTag,
		tags:                    make([]string, 0),
	}
}

func (nb *MultinomialNB) Train() {
	tagCounts := make(map[string]int)
	for tag, articles := range nb.articlesPerTag {
		tagCounts[tag] = len(articles)
		nb.tags = append(nb.tags, tag)
	}

	for tag, count := range tagCounts {
		nb.priorsPerTag[tag] = float64(count) / float64(sum(tagCounts))
	}

	nb.likelihoodPerWordPerTag = nb.getWordLikelihoodsPerTag()
}

func (nb *MultinomialNB) Predict(article []string) map[string]float64 {
	posteriorsPerTag := make(map[string]float64)
	for _, tag := range nb.tags {
		posteriorsPerTag[tag] = math.Log(nb.priorsPerTag[tag])
	}

	for _, word := range article {
		for _, tag := range nb.tags {
			posteriorsPerTag[tag] = posteriorsPerTag[tag] + math.Log(nb.likelihoodPerWordPerTag[word][tag])
		}
	}

	return posteriorsPerTag
}

func (nb *MultinomialNB) getWordLikelihoodsPerTag() map[string]map[string]float64 {
	wordFrequenciesPerTag := make(map[string]map[string]int)
	totalWordCountPerTag := make(map[string]int)

	for _, tag := range nb.tags {
		for _, article := range nb.articlesPerTag[tag] {
			for _, word := range article {
				if _, ok := wordFrequenciesPerTag[word]; !ok {
					wordFrequenciesPerTag[word] = make(map[string]int)
				}
				wordFrequenciesPerTag[word][tag] += 1
				totalWordCountPerTag[tag] += 1
			}
		}
	}

	wordLikelihoodsPerTag := make(map[string]map[string]float64)
	for word, tagsMap := range wordFrequenciesPerTag {
		wordLikelihoodsPerTag[word] = make(map[string]float64)
		for _, tag := range nb.tags {
			wordLikelihoodsPerTag[word][tag] = (float64(wordFrequenciesPerTag[word][tag]) + 1*nb.alpha) / (float64(totalWordCountPerTag[tag]) + 2*nb.alpha)
		}
	}

	return wordLikelihoodsPerTag
}

func sum(values map[string]int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	articlesPerTag := map[string][][]string{
		"politics": {
			{
				"trump impeachment",
				"hillary clinton",
				"barack obama",
				"donald trump",
				"hillary clinton",
				"barack obama",
				"donald trump",
				"hillary clinton",
			},
		},
		"technology": {
			{
				"artificial intelligence",
				"machine learning",
				"deep learning",
				"computer science",
				"artificial intelligence",
				"machine learning",
				"deep learning",
				"computer science",
			},
		},
		"science": {
			{
				"global warming",
				"climate change",
				"black hole",
				"dark matter",
				"global warming",
				"climate change",
				"black hole",
				"dark matter",
			},
		},
	}

	multinomialNB := NewMultinomialNB(articlesPerTag)
	multinomialNB.Train()
	article := []string{
		"donald trump",
		"hillary clinton",
		"barack obama",
		"donald trump",
		"hillary clinton",
	}
	posteriorsPerTag := multinomialNB.Predict(article)
	fmt.Println(posteriorsPerTag)
}
