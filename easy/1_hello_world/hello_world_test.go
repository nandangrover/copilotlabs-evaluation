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

    helloworld()

    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = rescueStdout

    s := strings.TrimSpace(string(out))
    if s != "Hello World" {
        t.Errorf("Expected %s, got %s", "Hello World", out)
    }
}
