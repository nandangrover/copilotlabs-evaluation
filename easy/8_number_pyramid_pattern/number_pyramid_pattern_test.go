package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	pyramidPattern(5)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	s := strings.TrimSpace(string(out))

	expected := strings.TrimSpace("1\n   22\n  333\n 4444\n55555")
	if s != expected {
		t.Errorf("Expected %s, got %s", expected, s)
	}
}
