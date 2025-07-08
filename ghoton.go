package main

import (
	"fmt"
	"github.com/doconnell2020/ghoton/plot"
	"github.com/doconnell2020/ghoton/spectra"
	"os"
	"strconv"
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
	s := &spectra.Spectrum{}
	s.Well = arrDatum[0]
	s.Name = arrDatum[1]
	atoi, err := strconv.Atoi(arrDatum[2])
	if err != nil {
		fmt.Println("Error while converting dilution: ", arrDatum[1])
		os.Exit(1)
	}
	s.Dilution = atoi
	for _, i := range arrDatum[3:] {
		j, err := strconv.ParseFloat(i, 32)
		if err != nil {
			fmt.Println("Error while converting wavelength: ", i)
		}
		s.Data = append(
			s.Data,
			j,
		)
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
