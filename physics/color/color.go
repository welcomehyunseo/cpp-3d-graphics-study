package color

import (
	"image/color"
)

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

func NewColor(red uint8, green uint8, blue uint8) *Color {
	return &Color{
		red:   red,
		green: green,
		blue:  blue,
	}
}

func (c *Color) GetRed() uint8 {
	return c.red
}

func (c *Color) GetGreen() uint8 {
	return c.green
}

func (c *Color) GetBlue() uint8 {
	return c.blue
}

func (c *Color) ToRGBA() *color.RGBA {
	return &color.RGBA{R: c.red, G: c.green, B: c.blue, A: 0xff}
}

func (c *Color) Add(c1 *Color) *Color {
	f0 := float64(c.red) + float64(c1.red)
	f1 := float64(c.green) + float64(c1.green)
	f2 := float64(c.blue) + float64(c1.blue)
	var red uint8
	var green uint8
	var blue uint8
	if f0 < 0 {
		red = 0
	} else if f0 > 0xff {
		red = 0xff
	} else {
		red = uint8(f0)
	}
	if f1 < 0 {
		green = 0
	} else if f1 > 0xff {
		green = 0xff
	} else {
		green = uint8(f1)
	}
	if f2 < 0 {
		blue = 0
	} else if f2 > 0xff {
		blue = 0xff
	} else {
		blue = uint8(f2)
	}
	return NewColor(red, green, blue)
}

func (c *Color) ApplyIntensity(i float64) *Color {
	f0 := float64(c.red) * i
	f1 := float64(c.green) * i
	f2 := float64(c.blue) * i
	var red uint8
	var green uint8
	var blue uint8
	if f0 < 0 {
		red = 0
	} else if f0 > 0xff {
		red = 0xff
	} else {
		red = uint8(f0)
	}
	if f1 < 0 {
		green = 0
	} else if f1 > 0xff {
		green = 0xff
	} else {
		green = uint8(f1)
	}
	if f2 < 0 {
		blue = 0
	} else if f2 > 0xff {
		blue = 0xff
	} else {
		blue = uint8(f2)
	}
	return NewColor(red, green, blue)
}
