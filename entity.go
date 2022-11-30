package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Entity struct {
	Object *resolv.Object
	Image  *ebiten.Image
}

var (
	imageTreeTrunk  = PrepareImage("./assets/sprites/tree_trunk.png", &ebiten.DrawImageOptions{})
	imageTreeCanopy = PrepareImage("./assets/sprites/tree_canopy.png", &ebiten.DrawImageOptions{})
)

func (g *Game) CreateBorders() error {
	g.space.Add(
		resolv.NewObject(-TileSize, -TileSize, ScreenWidth+2*TileSize, TileSize+1, "collidable"),      // top
		resolv.NewObject(-TileSize, ScreenHeight-1, ScreenWidth+2*TileSize, TileSize+1, "collidable"), // bottom
		resolv.NewObject(-TileSize, -TileSize, TileSize+1, ScreenHeight+2*TileSize, "collidable"),     // left
		resolv.NewObject(ScreenWidth-1, -TileSize, TileSize+1, ScreenHeight+2*TileSize, "collidable"), // right
	)

	return nil
}

func (g *Game) CreateEntity(img *ebiten.Image, x, y, w, h float64, isCollidable bool) error {
	// Create a new entity struct and add its object to resolv space, including appropriate tags.
	tags := []string{"entity"}
	if isCollidable {
		tags = append(tags, "collidable")
	}

	entity := &Entity{
		Object: resolv.NewObject(x, y, w, h, tags...),
		Image:  img,
	}

	g.entities = append(g.entities, entity)

	g.space.Add(entity.Object)

	return nil
}

func (g *Game) CreateTree(bottom, center int) error {
	// Create tree trunk (collidable)
	wTrunk, hTrunk := imageTreeTrunk.Size()
	xTrunk := center - wTrunk/2
	yTrunk := bottom - hTrunk
	g.CreateEntity(imageTreeTrunk, float64(xTrunk), float64(yTrunk), float64(wTrunk), float64(hTrunk), true)

	// Create tree trunk (not collidable)
	wCanopy, hCanopy := imageTreeCanopy.Size()
	xCanopy := center - wCanopy/2
	yCanopy := yTrunk - hCanopy
	g.CreateEntity(imageTreeCanopy, float64(xCanopy), float64(yCanopy), float64(wCanopy), float64(hCanopy), false)

	return nil
}

func GetRandomCoordinate() (int, int) {
	return rand.Intn(ScreenWidth), rand.Intn(ScreenHeight)
}
