package main

import (
	"image/color"
)

type Sphere struct {
	center *Vector
	radius float64
	rgba   *color.RGBA
}

func NewSphere(center *Vector, radius float64, rgba *color.RGBA) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		rgba:   rgba,
	}
}

func (o *Sphere) GetCenter() *Vector {
	return o.center
}

func (o *Sphere) GetRadius() float64 {
	return o.radius
}

func (o *Sphere) GetRGBA() *color.RGBA {
	return o.rgba
}
