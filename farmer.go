package main

type Farmer struct {
	sprite *Sprite
}

const (
	MaxFarmerSpeed = 3
)

func (f *Farmer) UpdateSpeed(deltaVX int, deltaVY int) {
	f.sprite.vx += deltaVX
	f.sprite.vy += deltaVY

	if f.sprite.vx > MaxFarmerSpeed {
		f.sprite.vx = MaxFarmerSpeed
	} else if f.sprite.vx < -MaxFarmerSpeed {
		f.sprite.vx = -MaxFarmerSpeed
	}

	if f.sprite.vy > MaxFarmerSpeed {
		f.sprite.vy = MaxFarmerSpeed
	} else if f.sprite.vy < -MaxFarmerSpeed {
		f.sprite.vy = -MaxFarmerSpeed
	}
}
