package main

import (
	"bytes"
	"fmt"
	"github.com/doconnell2020/ghoton/spectra"
	"strings"

	//"github.com/doconnell2020/ghoton/spectra"
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
	arr := strings.Split(content, []string("\n"))
	datum := arr[11]
	spec, err := spectra.NewSpectrumFromArray(string(datum))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Content is :\n ", spec)

	//data = spectra.Spectra{}
}
