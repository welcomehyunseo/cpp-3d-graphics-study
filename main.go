package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/welcomehyunseo/golang-3d-graphics-example/geometry"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	Width               = 1024
	Height              = 768
	DistanceToViewport  = 500
	DefaultViewDistance = 10
)

var (
	BackgroundColor = &color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

type Sphere struct {
	center *geometry.Vector
	radius float64
	rgba   *color.RGBA
}

type Viewport struct {
	width  uint
	height uint
}

type Camera struct {
	origin             *geometry.Vector
	viewport           *Viewport
	distanceToViewport uint
	viewDistance       uint
}

type MyGame struct {
	framebuffer []*color.RGBA
	spheres     []*Sphere
	camera      *Camera
}

func NewMyGame(origin *geometry.Vector, width, height, distanceToViewport, viewDistance uint) *MyGame {
	framebuffer := make([]*color.RGBA, width*height)
	for i, _ := range framebuffer {
		framebuffer[i] = BackgroundColor
	}

	return &MyGame{
		framebuffer: framebuffer,
		spheres:     []*Sphere{},
		camera: &Camera{
			origin,
			&Viewport{width, height},
			distanceToViewport,
			viewDistance,
		},
	}
}

func (g *MyGame) AddSphere(sphere *Sphere) {
	g.spheres = append(g.spheres, sphere)
}

// IntersectRaySphere
// return isMet, t1, t2
func (g *MyGame) IntersectRaySphere(O, D *geometry.Vector, sphere *Sphere) (bool, float64, float64) {
	radius := sphere.radius
	CO := geometry.Subtract(O, sphere.center)

	a := geometry.Dot(D, D)
	b := geometry.Dot(CO, D)
	c := geometry.Dot(CO, CO) - math.Pow(radius, 2)

	discriminant := math.Pow(b, 2) - a*c
	if discriminant < 0 {
		return false, 0, 0
	}
	t1 := (-b + math.Sqrt(discriminant)) / a
	t2 := (-b - math.Sqrt(discriminant)) / a
	return true, t1, t2
}

func (g *MyGame) TraceRay(O, D *geometry.Vector) *color.RGBA {
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
			V := geometry.NewVector(vx, vy, vz)
			O := g.camera.origin
			D := geometry.Subtract(V, O)
			rgba := g.TraceRay(O, D)
			i := l*vw + k
			g.framebuffer[i] = rgba
		}
	}
}

func (g *MyGame) Update() error {
	return nil
}

func (g *MyGame) Draw(screen *ebiten.Image) {

	vh := g.camera.viewport.height
	vw := g.camera.viewport.width

	for l := uint(0); l < vh; l++ {
		for k := uint(0); k < vw; k++ {
			i := l*vw + k
			rgba := g.framebuffer[i]
			screen.Set(int(k), int(l), rgba)
		}
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func RandomRGBA() *color.RGBA {
	return &color.RGBA{
		R: uint8(rand.Intn(0xff)),
		G: uint8(rand.Intn(0xff)),
		B: uint8(rand.Intn(0xff)),
		A: uint8(rand.Intn(0xff)),
	}
}

func RandomSphere() *Sphere {
	rand.Seed(time.Now().UnixNano())
	return &Sphere{
		center: geometry.NewVector(rand.Float64()*100, rand.Float64()*100, rand.Float64()*1000),
		radius: rand.Float64() * 1000,
		rgba:   RandomRGBA(),
	}
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Basic Raytracing")

	origin := geometry.NewVector(0, 0, 0)
	myGame := NewMyGame(origin, Width, Height, DistanceToViewport, DefaultViewDistance)

	for i := 0; i < 10; i++ {
		sphere := RandomSphere()
		myGame.AddSphere(sphere)
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
