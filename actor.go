package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Actor struct {
	Object *resolv.Object
	Image  *ebiten.Image
	SpeedX float64
	SpeedY float64
	Noise  Noise
}

func (g *Game) CreateActor(img *ebiten.Image, x, y, w, h float64) error {
	// Create player struct
	actor := &Actor{
		Object: resolv.NewObject(x, y, w, h, "actor", "collidable"),
		Image:  img,
		SpeedX: GetRandomSpeed(),
		SpeedY: GetRandomSpeed(),
		Noise:  GenerateNoise(100, 100),
	}
	g.actors = append(g.actors, actor)

	// Add player object to space
	g.space.Add(actor.Object)

	return nil
}

func (g *Game) MoveActor(a *Actor, dx, dy float64) {
	if collision := a.Object.Check(dx, 0, "collidable"); collision != nil {
		dx = collision.ContactWithObject(collision.Objects[0]).X()
	}

	a.Object.X += dx
	a.Object.Update()

	if collision := a.Object.Check(0, dy, "collidable"); collision != nil {
		dy = collision.ContactWithObject(collision.Objects[0]).Y()
	}

	a.Object.Y += dy
	a.Object.Update()
}

func (g *Game) MoveActors() {
	for _, actor := range g.actors {
		if rand.Float64() <= 0.025 {
			actor.SpeedX = GetRandomSpeed()
		}

		if rand.Float64() <= 0.025 {
			actor.SpeedY = GetRandomSpeed()
		}

		g.MoveActor(actor, actor.SpeedX, actor.SpeedY)
	}
}

func GetRandomSpeed() float64 {
	return 0.5 * (1 - 2*rand.Float64())
}
