package spectra

import (
	"fmt"
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
	var builder strings.Builder
	for i, datum := range sp.Data {
		wavelength := 220 + i*2
		builder.WriteString(fmt.Sprintf("\n\u03BB %d: %v", wavelength, datum))
	}
	return fmt.Sprintf(
		"\nWell: %s\nName: %s\nDilution: %v\nWavelength (\u03BB): %s",
		sp.Well,
		sp.Name,
		sp.Dilution,
		builder.String(),
	)
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
