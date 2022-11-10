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

	invertedStar(5)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	s := strings.TrimSpace(string(out))

	expected := strings.TrimSpace("    *\n   **\n  ***\n ****\n*****")
	if s != expected {
		t.Errorf("Expected %s, got %s", expected, s)
	}
}
