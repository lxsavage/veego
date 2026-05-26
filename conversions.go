package veego

type colorRGB int
type colorTemp int
type percentage int

// Get an RGB value from its components
func RGB(r, g, b uint8) colorRGB {
	return colorRGB((int(r) << 16) | (int(g) << 8) | int(b))
}

// Get a number as a Kelvin temperature
func K(k int) colorTemp {
	if k < 2000 {
		return 2000
	}
	if k > 9000 {
		return 9000
	}
	return colorTemp(k)
}

// Get a number as a percentage
func Percentage(p int) percentage {
	if p < 0 {
		return 0
	}
	if p > 100 {
		return 100
	}
	return percentage(p)
}
