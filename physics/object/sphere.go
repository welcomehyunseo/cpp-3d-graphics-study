package object

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type Sphere struct {
	object
	radius float64
}

func NewSphere(center *vector.Vector, color *color.Color, specular, reflective, radius float64) *Sphere {
	return &Sphere{
		radius: radius,
		object: *newObject(center, color, specular, reflective),
	}
}

func (o *Sphere) GetRadius() float64 {
	return o.radius
}
