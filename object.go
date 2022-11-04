package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	sprite     *Sprite
	collidable bool
}

func (g *Game) CreateRandomObject(img ebiten.Image, collidable bool) *Object {
	w, h := img.Size()
	potentialSprite := &Sprite{
		image:       &img,
		imageWidth:  float64(w),
		imageHeight: float64(h),
		pos:         &Coordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
	}

	for {
		if g.CheckCollision(potentialSprite) {
			potentialSprite = &Sprite{
				image:       &img,
				imageWidth:  float64(w),
				imageHeight: float64(h),
				pos:         &Coordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
			}
		} else {
			break
		}
	}

	return &Object{
		sprite:     potentialSprite,
		collidable: collidable,
	}
}
