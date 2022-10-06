package vector

import "math"

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v *Vector) Normalize() *Vector {
	return v.Divide(v.GetLength())
}

func (v *Vector) GetLength() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

func (v *Vector) Add(right *Vector) *Vector {
	return NewVector(v.x+right.x, v.y+right.y, v.z+right.z)
}

func (v *Vector) Subtract(right *Vector) *Vector {
	return NewVector(v.x-right.x, v.y-right.y, v.z-right.z)
}

// Dot - dot product
func (v *Vector) Dot(right *Vector) float64 {
	return v.x*right.x + v.y*right.y + v.z*right.z
}

func (v *Vector) Multiply(num float64) *Vector {
	return NewVector(num*v.x, num*v.y, num*v.z)
}

func (v *Vector) Divide(num float64) *Vector {
	return NewVector(v.x/num, v.y/num, v.z/num)
}
