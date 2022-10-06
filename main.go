package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/light"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/object"
	"github.com/welcomehyunseo/golang-3d-graphics-example/physics/vector"
)

const (
	Width                       = 1024
	Height                      = 768
	DistanceToViewport          = 500
	DefaultViewDistanceMultiple = 10
)

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Basic Raytracing")

	cameraCenter := vector.NewVector(0, 0, 0)
	myGame := physics.NewMyGame(cameraCenter, Width, Height, DistanceToViewport, DefaultViewDistanceMultiple)

	s0 := object.NewSphere(
		vector.NewVector(-2000, 0, 4500),
		2500,
		color.NewColor(0xff, 0, 0),
	)
	myGame.AddSphere(s0)
	s1 := object.NewSphere(
		vector.NewVector(0, 0, 4000),
		1500,
		color.NewColor(0, 0xff, 0),
	)
	myGame.AddSphere(s1)
	s2 := object.NewSphere(
		vector.NewVector(1000, -1000, 3000),
		1000,
		color.NewColor(0, 0, 0xff),
	)
	myGame.AddSphere(s2)

	l0 := light.NewAmbientLight(0.2)
	myGame.AddLight(l0)
	l1 := light.NewPointLight(0.4, vector.NewVector(10000, 0, 0))
	myGame.AddLight(l1)
	l2 := light.NewDirectionalLight(0.6, vector.NewVector(0, -1, 0))
	myGame.AddLight(l2)

	go func() {
		for {
			myGame.UpdateFramebuffer()
		}
	}()

	if err := ebiten.RunGame(myGame); err != nil {
		panic(err)
	}
}
