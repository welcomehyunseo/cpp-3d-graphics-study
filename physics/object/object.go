package object

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type Sphere struct {
	center *vector.Vector
	radius float64
	rgba   *color.Color
}

func NewSphere(center *vector.Vector, radius float64, rgba *color.Color) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		rgba:   rgba,
	}
}

func (o *Sphere) GetCenter() *vector.Vector {
	return o.center
}

func (o *Sphere) GetRadius() float64 {
	return o.radius
}

func (o *Sphere) GetColor() *color.Color {
	return o.rgba
}
