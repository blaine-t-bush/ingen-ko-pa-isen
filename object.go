package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	image       *ebiten.Image
	pos         *ScreenCoordinate
	width       float64
	height      float64
	collidable  bool
	aboveActors bool
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
		image:       img,
		pos:         &boundingBox.pos,
		width:       boundingBox.width,
		height:      boundingBox.height,
		collidable:  collidable,
		aboveActors: aboveActors,
	}
}

func (g *Game) CreateRandomTree() []*Object {
	wTrunk, hTrunk := treeTrunkImage.Size()
	boundingBox := &BoundingBox{
		pos:    ScreenCoordinate{float64(rand.Intn(screenWidth - wTrunk)), float64(rand.Intn(screenHeight - hTrunk))},
		width:  float64(wTrunk),
		height: float64(hTrunk),
	}

	for {
		if g.CheckCollision(*boundingBox) {
			boundingBox.pos = ScreenCoordinate{float64(rand.Intn(screenWidth - wTrunk)), float64(rand.Intn(screenHeight - hTrunk))}
		} else {
			break
		}
	}

	trunk := &Object{
		image:       treeTrunkImage,
		pos:         &boundingBox.pos,
		width:       float64(wTrunk),
		height:      float64(hTrunk),
		collidable:  true,
		aboveActors: false,
	}

	wCanopy, hCanopy := treeCanopyImage.Size()

	canopy := &Object{
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
