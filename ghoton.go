package main

import (
	"fmt"
	"github.com/doconnell2020/ghoton/plot"
	"github.com/doconnell2020/ghoton/spectra"
	"os"
	"strings"
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
	arr := strings.Split(string(content), string("\n"))
	datum := arr[11]
	arrDatum := strings.Split(string(datum), string(","))

	s, err := spectra.NewSpectrumFromArray(arrDatum)
	if err != nil {
		fmt.Println("Error creating spectrum: ", err)
	}
	fmt.Println(
		"Content is :\n ",
		*s,
	)
	err = plot.Spectrum(s.Data, "output/points.png")
	if err != nil {
		return
	}
}
