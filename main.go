package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

const (
	Width                       = 1024
	Height                      = 768
	DistanceToViewport          = 500
	DefaultViewDistanceMultiple = 10
)

var (
	BackgroundColor = &color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

type Viewport struct {
	width  int
	height int
}

type Camera struct {
	origin               *Vector
	viewport             *Viewport
	distanceToViewport   int
	viewDistanceMultiple int
}

type MyGame struct {
	framebuffer []*color.RGBA
	spheres     []*Sphere
	camera      *Camera
}

func NewMyGame(origin *Vector, width, height, distanceToViewport, viewDistanceMultiple int) *MyGame {
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
			viewDistanceMultiple,
		},
	}
}

func (g *MyGame) AddSphere(sphere *Sphere) {
	g.spheres = append(g.spheres, sphere)
}

// IntersectRaySphere
// return isMet, t1, t2
func (g *MyGame) IntersectRaySphere(O, D *Vector, sphere *Sphere) (bool, float64, float64) {
	radius := sphere.GetRadius()
	CO := Subtract(O, sphere.GetCenter())

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
	tMax := g.camera.viewDistanceMultiple
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
	return closestSphere.GetRGBA()
}

func (g *MyGame) UpdateFramebuffer() {
	vw := g.camera.viewport.width
	vh := g.camera.viewport.height

	for l := 0; l < vh; l++ {
		for k := 0; k < vw; k++ {
			vx := float64(k - vw/2)
			vy := -float64(l - vh/2)
			vz := float64(g.camera.distanceToViewport)
			V := NewVector(vx, vy, vz)
			O := g.camera.origin
			D := Subtract(V, O)
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

	for l := 0; l < vh; l++ {
		for k := 0; k < vw; k++ {
			i := l*vw + k
			rgba := g.framebuffer[i]
			screen.Set(k, l, rgba)
		}
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Basic Raytracing")

	origin := NewVector(0, 0, 0)
	myGame := NewMyGame(origin, Width, Height, DistanceToViewport, DefaultViewDistanceMultiple)

	sphere := NewSphere(
		NewVector(0, 0, 1000),
		300,
		&color.RGBA{R: 0xff, A: 0xff},
	)
	myGame.AddSphere(sphere)
	sphere1 := NewSphere(
		NewVector(300, 100, 600),
		200,
		&color.RGBA{G: 0xff, A: 0xff},
	)
	myGame.AddSphere(sphere1)

	sphere2 := NewSphere(
		NewVector(-500, -200, 1200),
		500,
		&color.RGBA{B: 0xff, A: 0xff},
	)
	myGame.AddSphere(sphere2)

	go func() {
		for {
			myGame.UpdateFramebuffer()
		}
	}()

	if err := ebiten.RunGame(myGame); err != nil {
		panic(err)
	}
}
