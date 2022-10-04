package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	Width              = 1024
	Height             = 768
	DistanceToViewport = 500

	DefaultViewDistance = 10

	MaxSpheresNum = 3
)

var (
	BackgroundColor = &color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

type Vector struct {
	x, y, z float64
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

func Sub(left, right *Vector) *Vector {
	return &Vector{left.x - right.x, left.y - right.y, left.z - right.z}
}

type Sphere struct {
	center *Vector
	radius float64
	rgba   *color.RGBA
}

type Viewport struct {
	width  uint
	height uint
}

type Camera struct {
	origin             *Vector
	viewport           *Viewport
	distanceToViewport uint
	viewDistance       uint
}

type MyGame struct {
	framebuffer   []*color.RGBA
	maxSpheresNum uint8
	spheres       []*Sphere
	camera        *Camera
}

func NewMyGame(maxSpheresNum uint8, origin *Vector, width, height, distanceToViewport, viewDistance uint) *MyGame {
	framebuffer := make([]*color.RGBA, width*height)
	for i, _ := range framebuffer {
		framebuffer[i] = BackgroundColor
	}

	return &MyGame{
		framebuffer:   framebuffer,
		maxSpheresNum: maxSpheresNum,
		spheres:       make([]*Sphere, maxSpheresNum),
		camera: &Camera{
			origin,
			&Viewport{width, height},
			distanceToViewport,
			viewDistance,
		},
	}
}

// IntersectRaySphere
// return isMet, t1, t2
func (g *MyGame) IntersectRaySphere(O, D *Vector, sphere *Sphere) (bool, float64, float64) {
	radius := sphere.radius
	CO := Sub(O, sphere.center)

	a := Dot(D, D)
	b := Dot(CO, D)
	c := Dot(CO, CO) - math.Pow(radius, 2)

	discriminant := math.Pow(b, 2) - a*c
	if discriminant < 0 {
		return false, 0, 0
	}
	t1 := (-b + math.Sqrt(discriminant)) / a
	t2 := (-b - math.Sqrt(discriminant)) / a
	return true, t1, t2
}

func (g *MyGame) TraceRay(O, D *Vector) *color.RGBA {
	tMin := 1
	tMax := g.camera.viewDistance
	closestT := math.MaxFloat64
	var closestSphere *Sphere = nil
	for _, sphere := range g.spheres {
		isMet, t1, t2 := g.IntersectRaySphere(O, D, sphere)
		if !isMet {
			continue
		}
		if float64(tMin) <= t1 && t1 <= float64(tMax) && t1 < closestT {
			closestT = t1
			closestSphere = sphere
		}
		if float64(tMin) <= t2 && t2 <= float64(tMax) && t2 < closestT {
			closestT = t2
			closestSphere = sphere
		}
	}
	if closestSphere == nil {
		return BackgroundColor
	}
	return closestSphere.rgba
}

func (g *MyGame) UpdateFramebuffer() {
	vh := g.camera.viewport.height
	vw := g.camera.viewport.width

	for l := uint(0); l < vh; l++ {
		for k := uint(0); k < vw; k++ {
			vx := float64(k) - float64(vw/2)
			vy := float64(l) - float64(vh/2)
			vz := float64(g.camera.distanceToViewport)
			V := &Vector{vx, vy, vz}
			O := g.camera.origin
			D := Sub(V, O)
			rgba := g.TraceRay(O, D)
			i := l*vw + k
			g.framebuffer[i] = rgba
		}
	}
}

func (g *MyGame) Update(screen *ebiten.Image) error {

	vh := g.camera.viewport.height
	vw := g.camera.viewport.width

	for l := uint(0); l < vh; l++ {
		for k := uint(0); k < vw; k++ {
			i := l*vw + k
			rgba := g.framebuffer[i]
			screen.Set(int(k), int(l), rgba)
		}
	}
	return nil
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Basic Raytracing")

	origin := &Vector{0, 0, 0}
	myGame := NewMyGame(MaxSpheresNum, origin, Width, Height, DistanceToViewport, DefaultViewDistance)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < MaxSpheresNum; i++ {
		sphere := &Sphere{
			center: &Vector{rand.Float64() * 100, rand.Float64() * 100, rand.Float64() * 1000},
			radius: rand.Float64() * 1000,
			rgba:   &color.RGBA{R: uint8(rand.Intn(0xff)), G: uint8(rand.Intn(0xff)), B: uint8(rand.Intn(0xff)), A: uint8(rand.Intn(0xff))},
		}
		myGame.spheres[i] = sphere
	}

	go func() {
		for {
			myGame.UpdateFramebuffer()
		}
	}()

	if err := ebiten.RunGame(myGame); err != nil {
		panic(err)
	}
}
