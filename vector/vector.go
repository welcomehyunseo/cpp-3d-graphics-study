package vector

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func Add(left, right *Vector) *Vector {
	return &Vector{left.x + right.x, left.y + right.y, left.z + right.z}
}

func Subtract(left, right *Vector) *Vector {
	return &Vector{left.x - right.x, left.y - right.y, left.z - right.z}
}

// Dot - dot product
func Dot(left, right *Vector) float64 {
	return left.x*right.x + left.y*right.y + left.z*right.z
}

// Scalar - scalar product
func Scalar(num float64, right *Vector) float64 {
	return right.x*num + right.y*num + right.z*num
}
