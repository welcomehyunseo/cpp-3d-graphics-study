package light

import (
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

type Light interface {
	GetIntensity() float64
}

type light struct {
	intensity float64
}

func (l *light) GetIntensity() float64 {
	return l.intensity
}

type AmbientLight struct {
	light
}

func NewAmbientLight(intensity float64) *AmbientLight {
	return &AmbientLight{
		light: light{
			intensity: intensity,
		},
	}
}

type PointLight struct {
	light
	position *vector.Vector
}

func NewPointLight(intensity float64, position *vector.Vector) *PointLight {
	return &PointLight{
		light: light{
			intensity: intensity,
		},
		position: position,
	}
}

func (l *PointLight) GetPosition() *vector.Vector {
	return l.position
}

type DirectionalLight struct {
	light
	direction *vector.Vector
}

func NewDirectionalLight(intensity float64, direction *vector.Vector) *DirectionalLight {
	return &DirectionalLight{
		light: light{
			intensity: intensity,
		},
		direction: direction,
	}
}

func (l *DirectionalLight) GetDirection() *vector.Vector {
	return l.direction
}
