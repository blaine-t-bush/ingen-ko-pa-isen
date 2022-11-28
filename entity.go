package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Entity struct {
	Object *resolv.Object
	Image  *ebiten.Image
}

func (g *Game) CreateEntity(img *ebiten.Image, x, y, w, h float64) error {
	// Create entity struct
	entity := &Entity{
		Object: resolv.NewObject(x, y, w, h, "entity", "collidable"),
		Image:  img,
	}
	g.entities = append(g.entities, entity)

	// Add entity object to space
	g.space.Add(entity.Object)

	return nil
}
