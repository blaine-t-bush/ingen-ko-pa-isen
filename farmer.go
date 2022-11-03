package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FarmerSpeedMultiplier = 4
)

type Farmer struct {
	sprite *Sprite
}

func (f *Farmer) Update() {
	f.sprite.Update()
}

func (f *Farmer) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(FarmerSpeedMultiplier)
	f.sprite.Move(velocity)
}

func CreateFarmer(img ebiten.Image) *Farmer {
	w, h := img.Size()

	return &Farmer{
		sprite: &Sprite{
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(screenWidth / 2), float64(screenHeight / 2)},
		},
	}
}

func (g *Game) CheckFarmerCollision() {
	// loop over objects
	for _, object := range g.objects {
		// if object is collidable and rectangle is within boundaries
		if object.collidable {
			collidesTop, overlapTop := g.farmer.sprite.CollidesWithTopOf(*object.sprite)
			collidesBottom, overlapBottom := g.farmer.sprite.CollidesWithBottomOf(*object.sprite)
			collidesLeft, overlapLeft := g.farmer.sprite.CollidesWithLeftOf(*object.sprite)
			collidesRight, overlapRight := g.farmer.sprite.CollidesWithRightOf(*object.sprite)

			if collidesTop {
				g.farmer.sprite.pos.y -= overlapTop
			} else if collidesBottom {
				g.farmer.sprite.pos.y += overlapBottom
			} else if collidesLeft {
				g.farmer.sprite.pos.x -= overlapLeft
			} else if collidesRight {
				g.farmer.sprite.pos.x += overlapRight
			}
		}
	}
}
