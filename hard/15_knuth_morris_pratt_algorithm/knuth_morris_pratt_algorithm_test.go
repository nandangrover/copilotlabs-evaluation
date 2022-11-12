package main

import (
	"reflect"
	"testing"
)

func TestKnuthMorrisPrattAlgorithm(t *testing.T) {
	testCases := []struct {
		str    string
		substr string
		want   bool
	}{
		{"aefoaefcdaefcdaed", "aefcdaed", true},
	}
	for _, tc := range testCases {
		got := knuthMorrisPrattAlgorithm(tc.str, tc.substr)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("knuthMorrisPrattAlgorithm(%v, %v) = %v; want %v", tc.str, tc.substr, got, tc.want)
		}
	}
}
