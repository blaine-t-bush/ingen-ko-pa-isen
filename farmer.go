package main

const (
	FarmerSpeedMultiplier = 4
)

type Farmer struct {
	sprite *Sprite
}

func (f *Farmer) Update() {
	f.sprite.Update()
}

func (f *Farmer) Move(dir Direction) {
	f.sprite.Move(dir, FarmerSpeedMultiplier)
}
