package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var files []string
	err := filepath.Walk("../../../copilotlabs-evaluation", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(files))
}
