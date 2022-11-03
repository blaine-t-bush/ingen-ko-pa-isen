package main

import "github.com/hajimehoshi/ebiten/v2"

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
