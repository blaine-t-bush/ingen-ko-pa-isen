package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	image      *ebiten.Image
	pos        *ScreenCoordinate
	width      float64
	height     float64
	collidable bool
}

func (o *Object) BoundingBox() BoundingBox {
	return BoundingBox{pos: *o.pos, width: o.width, height: o.height}
}

func (g *Game) CreateRandomObject(img ebiten.Image, collidable bool) *Object {
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
		image:      &img,
		pos:        &boundingBox.pos,
		width:      boundingBox.width,
		height:     boundingBox.height,
		collidable: collidable,
	}
}

func (g *Game) CreateRandomTree(trunkImg *ebiten.Image, canopyImg *ebiten.Image) []*Object {
	wTrunk, hTrunk := trunkImg.Size()
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
		image:      trunkImg,
		pos:        &boundingBox.pos,
		width:      float64(wTrunk),
		height:     float64(hTrunk),
		collidable: true,
	}

	wCanopy, hCanopy := canopyImg.Size()

	canopy := &Object{
		image: canopyImg,
		pos:   &ScreenCoordinate{trunk.pos.x - (float64(wCanopy)-trunk.width)/2, trunk.pos.y - float64(hCanopy)},
	}

	return []*Object{trunk, canopy}
}
