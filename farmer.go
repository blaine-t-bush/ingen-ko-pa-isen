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

func (g *Game) MoveFarmer(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(FarmerSpeedMultiplier)
	// check if move would result in collision
	newSpriteX := &Sprite{
		imageWidth:  g.farmer.sprite.imageWidth,
		imageHeight: g.farmer.sprite.imageHeight,
		pos:         &Coordinate{g.farmer.sprite.pos.x + velocity.X(), g.farmer.sprite.pos.y},
	}
	newSpriteY := &Sprite{
		imageWidth:  g.farmer.sprite.imageWidth,
		imageHeight: g.farmer.sprite.imageHeight,
		pos:         &Coordinate{g.farmer.sprite.pos.x, g.farmer.sprite.pos.y + velocity.Y()},
	}
	newSpriteXY := &Sprite{
		imageWidth:  g.farmer.sprite.imageWidth,
		imageHeight: g.farmer.sprite.imageHeight,
		pos:         &Coordinate{g.farmer.sprite.pos.x + velocity.X(), g.farmer.sprite.pos.y + velocity.Y()},
	}

	validXY := !g.CheckCollision(newSpriteXY)
	validX := !g.CheckCollision(newSpriteX)
	validY := !g.CheckCollision(newSpriteY)

	if !validXY {
		if validX && !validY {
			velocity.SetXY(Coordinate{velocity.X(), 0})
		} else if validY && !validX {
			velocity.SetXY(Coordinate{0, velocity.Y()})
		} else {
			velocity = Vector{0, 0}
		}
	}

	g.farmer.sprite.Move(velocity)
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

func (g *Game) CheckCollision(s *Sprite) bool {
	collides := false
	// loop over objects
	for _, object := range g.objects {
		// if object is collidable and rectangle is within boundaries
		if object.collidable {
			collidesTop, _ := s.CollidesWithTopOf(*object.sprite)
			collidesBottom, _ := s.CollidesWithBottomOf(*object.sprite)
			collidesLeft, _ := s.CollidesWithLeftOf(*object.sprite)
			collidesRight, _ := s.CollidesWithRightOf(*object.sprite)

			// if collidesTop {
			// 	if overlapTop >= overlapLeft && overlapTop >= overlapRight {
			// 		g.farmer.sprite.pos.y -= overlapTop
			// 	}
			// } else if collidesBottom {
			// 	if overlapBottom >= overlapLeft && overlapBottom >= overlapRight {
			// 		g.farmer.sprite.pos.y += overlapBottom
			// 	}
			// }

			// if collidesLeft {
			// 	if overlapLeft >= overlapTop && overlapLeft >= overlapBottom {
			// 		g.farmer.sprite.pos.x -= overlapLeft
			// 	}
			// } else if collidesRight {
			// 	if overlapRight >= overlapTop && overlapRight >= overlapBottom {
			// 		g.farmer.sprite.pos.x += overlapRight
			// 	}
			// }

			if collidesTop || collidesBottom || collidesLeft || collidesRight {
				collides = true
			}
		}
	}

	return collides
}
