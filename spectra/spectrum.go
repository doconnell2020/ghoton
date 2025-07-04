package spectra

import (
	"fmt"
	"reflect"
	"strings"
)

type Spectra struct {
	data []Spectrum
}

func (sp *Spectrum) String() string {
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
	v := reflect.ValueOf(s).Elem()

	// Iterate over the remaining fields to populate spectrum
	for i := 0; i < v.NumField() && i < len(data); i++ {
		// TODO Capture all items into type
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}
		dataValue := reflect.ValueOf(data[i])
		if dataValue.Type().ConvertibleTo(field.Type()) {
			field.Set(dataValue.Convert(field.Type()))
		}

	}
	return s, nil
}

type Spectrum struct {
	Well     string
	Name     string
	Dilution int
	Data     []float32
}
