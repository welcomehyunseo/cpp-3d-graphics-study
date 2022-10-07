package physics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/light"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/object"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
	"math"
)

var (
	BackgroundColor = color.NewColor(255, 255, 255)
)

type Viewport struct {
	width  int
	height int
}

type Camera struct {
	center               *vector.Vector
	viewport             *Viewport
	distanceToViewport   float64
	viewDistanceMultiple float64
}

type MyGame struct {
	framebuffer []*color.Color
	lights      []light.Light
	spheres     []*object.Sphere
	camera      *Camera
}

func NewMyGame(cameraCenter *vector.Vector, width, height int, distanceToViewport, viewDistanceMultiple float64) *MyGame {
	framebuffer := make([]*color.Color, width*height)
	for i, _ := range framebuffer {
		framebuffer[i] = BackgroundColor
	}

	return &MyGame{
		framebuffer: framebuffer,
		spheres:     []*object.Sphere{},
		camera: &Camera{
			cameraCenter,
			&Viewport{width, height},
			distanceToViewport,
			viewDistanceMultiple,
		},
	}
}

func (g *MyGame) AddLight(light light.Light) {
	g.lights = append(g.lights, light)
}

func (g *MyGame) AddSphere(sphere *object.Sphere) {
	g.spheres = append(g.spheres, sphere)
}

// IntersectRaySphere
// params D = V - C (D is direction, V on viewport point, C is camera center)
// return isMet, t1, t2
func (g *MyGame) IntersectRaySphere(O, D *vector.Vector, sphere *object.Sphere) (bool, float64, float64) {

	radius := sphere.GetRadius()
	A := O.Subtract(sphere.GetCenter())

	a := D.Dot(D)
	b := A.Dot(D)
	c := A.Dot(A) - math.Pow(radius, 2)

	discriminant := math.Pow(b, 2) - a*c
	if discriminant < 0 {
		return false, 0, 0
	}
	t1 := (-b + math.Sqrt(discriminant)) / a
	t2 := (-b - math.Sqrt(discriminant)) / a
	return true, t1, t2
}

func (g *MyGame) ClosestSphere(O, D *vector.Vector, tMin, tMax float64) (float64, *object.Sphere) {
	closestT := math.MaxFloat64
	var closestSphere *object.Sphere = nil
	for _, sphere := range g.spheres {
		isMet, t1, t2 := g.IntersectRaySphere(O, D, sphere)
		if !isMet {
			continue
		}
		if tMin <= t1 && t1 <= tMax && t1 < closestT {
			closestT = t1
			closestSphere = sphere
		}
		if tMin <= t2 && t2 <= tMax && t2 < closestT {
			closestT = t2
			closestSphere = sphere
		}
	}
	return closestT, closestSphere
}

// ComputeLightIntensity
// params point at object surface, normal vector, D, specular
func (g *MyGame) ComputeLightIntensity(P, N, D *vector.Vector, s float64) float64 {
	var intensity float64 = 0

	for _, l := range g.lights {
		switch l.(type) {
		case *light.AmbientLight:
			intensity += l.GetIntensity()
			continue
		}

		var L *vector.Vector = nil
		var tMax float64
		switch l.(type) {
		default:
			continue
		case *light.PointLight:
			position := (l.(*light.PointLight)).GetPosition()
			L = position.Subtract(P)
			tMax = 1
			break
		case *light.DirectionalLight:
			direction := l.(*light.DirectionalLight).GetDirection()
			L = direction.Multiply(-1) // reverse because inner product (a*b=length(a)*length(b)*cos(angle))
			tMax = math.MaxFloat64
			break
		}
		// shadow check
		_, shadowSphere := g.ClosestSphere(P, L, 0.001, tMax)
		if shadowSphere != nil {
			continue
		}

		// diffuse reflection
		var0 := N.Dot(L)
		if var0 > 0 {
			intensity += l.GetIntensity() * (var0 / (N.GetLength() * L.GetLength()))
		}

		// specular reflection
		R := N.Multiply(2).Multiply(N.Dot(L)).Subtract(L)
		V := D.Multiply(-1)
		var0 = R.Dot(V)
		if var0 > 0 {
			intensity += l.GetIntensity() * math.Pow(var0/(R.GetLength()*V.GetLength()), s)
		}
	}
	return intensity
}

func (g *MyGame) TraceRay(O, D *vector.Vector) *color.Color {
	tMin := float64(1)
	tMax := g.camera.viewDistanceMultiple
	closestT, closestSphere := g.ClosestSphere(O, D, tMin, tMax)
	if closestSphere == nil {
		return BackgroundColor
	}

	A := g.camera.center
	P := A.Add(D.Multiply(closestT))            // origin to point
	CP := P.Subtract(closestSphere.GetCenter()) // sphere center to point
	N := CP.Normalize()

	intensity := g.ComputeLightIntensity(P, N, D, closestSphere.GetSpecular())

	return closestSphere.GetColor().ApplyIntensity(intensity)
}

func (g *MyGame) UpdateFramebuffer() {
	vw := g.camera.viewport.width
	vh := g.camera.viewport.height

	for l := 0; l < vh; l++ {
		for k := 0; k < vw; k++ {
			vx := float64(k - vw/2)
			vy := -float64(l - vh/2)
			vz := g.camera.distanceToViewport
			V := vector.NewVector(vx, vy, vz)
			C := g.camera.center
			D := V.Subtract(C)
			i := l*vw + k
			g.framebuffer[i] = g.TraceRay(C, D)
		}
	}
}

func (g *MyGame) Update() error {
	return nil
}

func (g *MyGame) Draw(screen *ebiten.Image) {

	vh := g.camera.viewport.height
	vw := g.camera.viewport.width

	for l := 0; l < vh; l++ {
		for k := 0; k < vw; k++ {
			i := l*vw + k
			c := g.framebuffer[i]
			screen.Set(k, l, c.ToRGBA())
		}
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
