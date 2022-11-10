package main

import (
	"testing"
	"os"
	"io/ioutil"
    "strings"
)

func TestPrint(t *testing.T) {
    rescueStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    alphabetTriangle()

    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = rescueStdout

    s := strings.TrimSpace(string(out))
	if s != "A \nA B \nA B C \nA B C D \nA B C D E \nA B C D E F \nA B C D E F G" {
        t.Errorf("Expected %s, got %s", "A \nA B \nA B C \nA B C D \nA B C D E \nA B C D E F \nA B C D E F G", out)
    }
}
