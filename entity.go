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

func (g *Game) CreateEntity(img *ebiten.Image, x, y, w, h float64, isCollidable bool) error {
	tags := []string{"entity"}
	if isCollidable {
		tags = append(tags, "collidable")
	}

	// Create entity struct
	entity := &Entity{
		Object: resolv.NewObject(x, y, w, h, tags...),
		Image:  img,
	}
	g.entities = append(g.entities, entity)

	// Add entity object to space
	g.space.Add(entity.Object)

	return nil
}

func (g *Game) CreateTree() error {
	// Create tree trunk (collidable)
	wTrunk, hTrunk := imageTreeTrunk.Size()
	xTrunk, yTrunk := GetRandomCoordinate()
	g.CreateEntity(imageTreeTrunk, float64(xTrunk), float64(yTrunk), float64(wTrunk), float64(hTrunk), true)

	// Create tree trunk (not collidable)
	wCanopy, hCanopy := imageTreeCanopy.Size()
	xCanopy := xTrunk + (wTrunk-wCanopy)/2
	yCanopy := yTrunk - hCanopy
	g.CreateEntity(imageTreeCanopy, float64(xCanopy), float64(yCanopy), float64(wCanopy), float64(hCanopy), true)

	return nil
}

func GetRandomCoordinate() (int, int) {
	return rand.Intn(ScreenWidth), rand.Intn(ScreenHeight)
}
