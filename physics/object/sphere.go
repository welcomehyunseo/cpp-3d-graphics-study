package object

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type Sphere struct {
	object
	radius float64
}

func NewSphere(center *vector.Vector, color *color.Color, specular float64, radius float64) *Sphere {
	return &Sphere{
		radius: radius,
		object: *newObject(center, color, specular),
	}
}

func (o *Sphere) GetRadius() float64 {
	return o.radius
}
