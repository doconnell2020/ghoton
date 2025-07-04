package main

import (
	"fmt"
	"github.com/doconnell2020/ghoton/spectra"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

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
	arr := strings.Split(string(content), string("\n"))
	datum := arr[11]
	arr_datum := strings.Split(string(datum), string(","))
	s := &spectra.Spectrum{}
	s.Well = arr_datum[0]
	s.Name = arr_datum[1]
	atoi, err := strconv.Atoi(arr_datum[2])
	if err != nil {
		fmt.Println("Error while converting dilution: ", arr_datum[1])
		os.Exit(1)
	}
	s.Dilution = atoi
	for _, i := range arr_datum[3:] {
		j, err := strconv.ParseFloat(i, 32)
		if err != nil {
			fmt.Println("Error while converting wavelength: ", i)
		}
		s.Data = append(s.Data, float32(j))
	}

	fmt.Println(
		"Content is :\n ",
		*s,
	)

}
