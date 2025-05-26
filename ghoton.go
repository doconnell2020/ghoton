package main

import (
	"fmt"
	"os"
)

func readCsv(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading file: ", path)
		os.Exit(1)
	}
	return content
}

func main() {
	path := os.Args[1]
	content := readCsv(path)
	fmt.Println("Content is :\n ", string(content))
}
