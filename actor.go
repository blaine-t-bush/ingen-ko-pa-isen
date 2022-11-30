package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Actor struct {
	Object *resolv.Object
	Image  *ebiten.Image
	SpeedX float64
	SpeedY float64
	NoiseX Noise
	NoiseY Noise
}

func (g *Game) CreateActor(img *ebiten.Image, x, y, w, h float64) error {
	// Create player struct
	actor := &Actor{
		Object: resolv.NewObject(x, y, w, h, "actor", "collidable"),
		Image:  img,
		SpeedX: 0,
		SpeedY: 0,
		NoiseX: GenerateNoise(50, 50, -1, 1),
		NoiseY: GenerateNoise(50, 50, -1, 1),
	}
	g.actors = append(g.actors, actor)

	// Use noise to determine speed.
	actor.UpdateSpeed()

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
		actor.UpdateSpeed()
		g.MoveActor(actor, actor.SpeedX, actor.SpeedY)
	}
}

func (a *Actor) UpdateSpeed() {
	a.SpeedX = a.NoiseX.UpdateAndGetValue()
	a.SpeedY = a.NoiseY.UpdateAndGetValue()
}
