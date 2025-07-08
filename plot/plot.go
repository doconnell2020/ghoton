package plot

import (
	"fmt"
	"gonum.org/v1/plot"
	_ "gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

const (
	startWavelength = 200.0
	stepSize        = 2.0
)

func Spectrum(data []float64, filename string) error {
	p := plot.New()
	p.Title.Text = "Spectrum"
	p.X.Label.Text = "Wavelength"
	p.Y.Label.Text = "Absorbance"

	pts := make(plotter.XYs, len(data))
	for i, absorbance := range data {
		wavelength := startWavelength + float64(i)*stepSize
		pts[i].X = wavelength
		pts[i].Y = absorbance
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.Add(line)
	p.Legend.Add("Sample", line)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, filename); err != nil {
		return err
	}
	return nil
}
