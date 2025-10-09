package spectra

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Spectrum struct {
	Well     string
	Name     string
	Dilution int
	Data     [291]float64
}

type Spectra struct {
	data []Spectrum
}

// Provide a cleaner string representation for print functions
func (sp Spectrum) String() string {
	v := reflect.ValueOf(sp).Elem()

	var builder strings.Builder
	builder.WriteString("Spectrum{\n")

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := field.Type()

		var valueStr string
		switch field.Kind() {
		case reflect.Float32, reflect.Float64:
			valueStr = fmt.Sprintf("%.2f", field.Float())
		case reflect.String:
			valueStr = fmt.Sprintf("\"%s\"", field.String())
		default:
			valueStr = fmt.Sprintf("%v", field.Interface())
		}

		builder.WriteString(fmt.Sprintf("  %s: %s\n", fieldType.Name(), valueStr))
	}

	builder.WriteString("}")
	return builder.String()
}

func NewSpectrumFromArray(data []string) (*Spectrum, error) {
	if len(data) < 292 {
		return nil, fmt.Errorf("insufficient data: got %d, need 292", len(data))
	}
	s := &Spectrum{}
	s.Well = data[0]
	s.Name = data[1]
	atoi, err := strconv.Atoi(data[2])
	if err != nil {
		fmt.Println("Error while converting dilution: ", data[1])
		return nil, err
	}
	s.Dilution = atoi
	for idx, i := range data[3:] {
		j, err := strconv.ParseFloat(i, 32)
		if err != nil {
			fmt.Println("Error while converting wavelength: ", i)
			return nil, err
		}
		s.Data[idx] = j
	}

	return s, nil

}
