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

func (f *Farmer) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(FarmerSpeedMultiplier)
	f.sprite.Move(velocity)
}
