package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	sprite     *Sprite
	collidable bool
}

func CreateRandomRock(img ebiten.Image) *Object {
	w, h := img.Size()
	return &Object{
		sprite: &Sprite{
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
		},
		collidable: true,
	}
}
