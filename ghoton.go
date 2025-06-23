package main

import (
	"fmt"
	"github.com/doconnell2020/ghoton/spectra"
	"reflect"
	"strconv"
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
	v := reflect.ValueOf(s).Elem()

	// Iterate over the remaining fields to populate spectrum
	for i := 3; i < v.NumField() && i < len(arr_datum); i++ {
		// TODO Capture all items into type
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}
		switch field.Type().Kind() {
		case reflect.Float32:
			if floatVal, err := strconv.ParseFloat(arr_datum[i], 32); err == nil {
				field.SetFloat(floatVal)
			} else {
				fmt.Println("Error parsing float for field %d %v\n", i, err)
			}
		case reflect.String:
			field.SetString(arr_datum[i])
		case reflect.Int:
			intVal, err := strconv.Atoi(arr_datum[i])
			if err == nil {
				field.SetInt(int64(intVal))
			} else {
				fmt.Println("Error parsing int for field %d %v\n", i, err)

			}
		default:
			panic("unhandled default case")
		}

	}

	fmt.Println(
		"Content is :\n ",
		*s,
	)

	//data = spectra.Spectra{}
}
