package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFiles(".", "go")

	new_num := Generator(1)
	for i := 0; i < 10; i += 1 {
		fmt.Println(new_num())
	}
}

func PrintAllFiles(path string, filter string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}
		if f.IsDir() {
			PrintAllFiles(filename, filter)
		}
	}
}

func Generator(iter int) func() int {
	return func() int {
		iter += 2
		return iter
	}
}
