package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Centroid struct {
	location      []float64
	closest_users []string
}

func get_k_means(user_feature_map map[string][]float64, num_features_per_user int, k int) [][]float64 {
	// Don't change the following two lines of code.
	rand.Seed(42)
	// Gets the inital users, to be used as centroids.
	inital_centroid_users := make([]string, 0)
	for user, _ := range user_feature_map {
		inital_centroid_users = append(inital_centroid_users, user)
	}
	sort.Strings(inital_centroid_users)
	inital_centroid_users = inital_centroid_users[:k]

	centroids := make([]Centroid, 0)
	for _, inital_centroid_user := range inital_centroid_users {
		centroids = append(centroids, Centroid{location: user_feature_map[inital_centroid_user]})
	}
	for i := 0; i < 10; i++ {
		for uid, features := range user_feature_map {
			closest_centroid_distance := math.Inf(1)
			var closest_centroid *Centroid = nil
			for _, centroid := range centroids {
				features_to_centroid_distance := get_manhattan_distance(features, centroid.location)
				if features_to_centroid_distance < closest_centroid_distance {
					closest_centroid_distance = features_to_centroid_distance
					closest_centroid = &centroid
				}
			}
			closest_centroid.closest_users = append(closest_centroid.closest_users, uid)
		}

		for _, centroid := range centroids {
			centroid.location = get_centroid_average(centroid, user_feature_map, num_features_per_user)
			centroid.closest_users = make([]string, 0)
		}
	}
	centroids_array := make([][]float64, 0)
	for _, centroid := range centroids {
		centroids_array = append(centroids_array, centroid.location)
	}
	return centroids_array
}

func get_manhattan_distance(features []float64, other_features []float64) float64 {
	absolute_differences := make([]float64, 0)
	for i := 0; i < len(features); i++ {
		absolute_differences = append(absolute_differences, math.Abs(features[i]-other_features[i]))
	}
	return sum(absolute_differences)
}

func sum(array []float64) float64 {
	sum := 0.0
	for _, element := range array {
		sum += element
	}
	return sum
}

func get_centroid_average(centroid Centroid, user_feature_map map[string][]float64, num_features_per_user int) []float64 {
	centroid_average := make([]float64, num_features_per_user)
	for i := 0; i < num_features_per_user; i++ {
		for _, user := range centroid.closest_users {
			centroid_average[i] = centroid_average[i] + user_feature_map[user][i]
		}
	}
	for i := 0; i < num_features_per_user; i++ {
		centroid_average[i] = centroid_average[i] / float64(len(centroid.closest_users))
	}
	return centroid_average
}

func main() {
	user_feature_map := map[string][]float64{
		"user1":  {1, 1, 1},
		"user2":  {2, 2, 2},
		"user3":  {3, 3, 3},
		"user4":  {4, 4, 4},
		"user5":  {5, 5, 5},
		"user6":  {6, 6, 6},
		"user7":  {7, 7, 7},
		"user8":  {8, 8, 8},
		"user9":  {9, 9, 9},
		"user10": {10, 10, 10},
	}
	num_features_per_user := 3
	k := 2
	centroids := get_k_means(user_feature_map, num_features_per_user, k)
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}
