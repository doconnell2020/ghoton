package plot

const (
	width, height = 600, 320 // canvas size in pixels
	cells         = 100      // number of grid cells
	xyrange       = 30.0     // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2
	zscale        = height * 0.4 // pixels per z unit
)

func plot()
