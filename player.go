package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	PlayerSpeed = 2
)

func (g *Game) CreatePlayer(img *ebiten.Image) error {
	// Create a new player struct and add it to the resolv space.
	g.player = &Actor{
		Object: resolv.NewObject(ScreenWidth/2, ScreenHeight/2, PlayerSize, PlayerSize, "actor", "player", "collidable"),
		Image:  img,
		SpeedX: 0,
		SpeedY: 0,
	}

	g.space.Add(g.player.Object)

	return nil
}

func (g *Game) MovePlayer(dx, dy float64) {
	// Check X and Y collision separately, starting with X.
	// Checking both simultaneously can result in strange shunting behavior.

	if collision := g.player.Object.Check(dx, 0, "collidable"); collision != nil {
		dx = collision.ContactWithObject(collision.Objects[0]).X()
	}

	g.player.Object.X += dx
	g.player.Object.Update()

	if collision := g.player.Object.Check(0, dy, "collidable"); collision != nil {
		dy = collision.ContactWithObject(collision.Objects[0]).Y()
	}

	g.player.Object.Y += dy
	g.player.Object.Update()
}
