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
	object := &Object{
		image:      &img,
		pos:        &ScreenCoordinate{x: float64(rand.Intn(screenWidth - w)), y: float64(rand.Intn(screenHeight - h))},
		width:      float64(w),
		height:     float64(h),
		collidable: collidable,
	}

	return object
}
