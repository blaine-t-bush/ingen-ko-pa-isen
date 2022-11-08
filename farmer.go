package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FarmerSpeedMultiplier = 4
)

func (g *Game) CreateFarmer(img ebiten.Image) *Actor {
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

	return &Actor{
		image:  &img,
		pos:    &boundingBox.pos,
		width:  boundingBox.width,
		height: boundingBox.height,
	}
}
