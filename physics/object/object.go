package object

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type object struct {
	center     *vector.Vector
	color      *color.Color
	specular   float64 // 0 < n <= 1000
	reflective float64 // 0 <= n < 1
}

func newObject(center *vector.Vector, color *color.Color, specular float64, reflective float64) *object {
	if specular < 0 || 1000 < specular {
		// TODO
	}
	return &object{
		center:     center,
		color:      color,
		specular:   specular,
		reflective: reflective,
	}
}

func (o *object) GetCenter() *vector.Vector {
	return o.center
}

func (o *object) GetColor() *color.Color {
	return o.color
}

func (o *object) GetSpecular() float64 {
	return o.specular
}

func (o *object) GetReflective() float64 {
	return o.reflective
}
