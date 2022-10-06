package object

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type object struct {
	center   *vector.Vector
	color    *color.Color
	specular float64 // 1 ~ 1000
}

func newObject(center *vector.Vector, color *color.Color, specular float64) *object {
	if specular < 0 || 1000 < specular {
		// TODO
	}
	return &object{
		center:   center,
		color:    color,
		specular: specular,
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
