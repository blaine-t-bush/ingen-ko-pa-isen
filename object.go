package main

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	id          string
	image       *ebiten.Image
	pos         *ScreenCoordinate
	width       float64
	height      float64
	collidable  bool
	aboveActors bool
	belongsTo   string
	createdAt   int64
}

func (o *Object) BoundingBox() BoundingBox {
	return BoundingBox{pos: *o.pos, width: o.width, height: o.height}
}

func (g *Game) CreateRandomObject(img *ebiten.Image, collidable bool, aboveActors bool) *Object {
	w, h := img.Size()
	boundingBox := &BoundingBox{
		pos:    ScreenCoordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
		width:  float64(w),
		height: float64(h),
	}

	for {
		if g.CheckCollision(*boundingBox) {
			boundingBox.pos = ScreenCoordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))}
		} else {
			break
		}
	}

	return &Object{
		id:          uuid.NewString(),
		image:       img,
		pos:         &boundingBox.pos,
		width:       boundingBox.width,
		height:      boundingBox.height,
		collidable:  collidable,
		aboveActors: aboveActors,
	}
}

func (g *Game) CreateRandomTree() []*Object {
	w, h := treeTrunkImage.Size()
	boundingBox := &BoundingBox{
		pos:    g.RandomCoordinateOfTerrainType(TerrainTypeSnow),
		width:  float64(w),
		height: float64(h),
	}

	for {
		if g.CheckCollision(*boundingBox) {
			boundingBox.pos = g.RandomCoordinateOfTerrainType(TerrainTypeSnow)
		} else {
			break
		}
	}

	trunk := &Object{
		id:          uuid.NewString(),
		image:       treeTrunkImage,
		pos:         &boundingBox.pos,
		width:       boundingBox.width,
		height:      boundingBox.height,
		collidable:  true,
		aboveActors: false,
	}

	wCanopy, hCanopy := treeCanopyImage.Size()

	canopy := &Object{
		id:          uuid.NewString(),
		image:       treeCanopyImage,
		pos:         &ScreenCoordinate{trunk.pos.x - (float64(wCanopy)-trunk.width)/2, trunk.pos.y - float64(hCanopy)},
		width:       float64(wCanopy),
		height:      float64(hCanopy),
		collidable:  false,
		aboveActors: true,
	}

	return []*Object{trunk, canopy}
}

func (g *Game) CreateRandomCowPie() *Object {
	return g.CreateRandomObject(cowPieImage, false, false)
}
