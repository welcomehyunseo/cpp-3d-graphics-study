package geometry

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func Dot(left, right *Vector) float64 {
	return left.x*right.x + left.y*right.y + left.z*right.z
}

func Multiply(left *Vector, num float64) float64 {
	return left.x*num + left.y*num + left.z*num
}

func Add(left, right *Vector) *Vector {
	return &Vector{left.x + right.x, left.y + right.y, left.z + right.z}
}

func Subtract(left, right *Vector) *Vector {
	return &Vector{left.x - right.x, left.y - right.y, left.z - right.z}
}
