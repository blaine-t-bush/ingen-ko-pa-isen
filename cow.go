package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CowSpeedMultiplier    = 2
	FarmerDetectionRadius = 120
)

type Cow struct {
	sprite   *Sprite
	velocity *Vector
}

func (c *Cow) Update(f *Farmer) {
	// Choose a new wander direction, possibly.
	c.velocity.dir = c.ChooseDirection(f)

	// Update position.
	c.Move(*c.velocity)
}

func (c *Cow) ChooseDirection(f *Farmer) float64 {
	dir := c.velocity.dir

	// Small chance to choose new direction.
	if rand.Float64() >= 0.95 {
		dir = c.velocity.dir + (math.Pi/2)*(0.5-rand.Float64())
	}

	// Flees directly away from farmer.
	if c.sprite.pos.WithinRadius(f.sprite.Center(), FarmerDetectionRadius) {
		dir = VectorFromPoints(f.sprite.Center(), c.sprite.Center()).dir
	}

	// Bounces off edges.
	if c.sprite.pos.x <= 0 {
		dir = 0
	} else if c.sprite.pos.x+c.sprite.imageWidth >= screenWidth {
		dir = math.Pi
	}

	if c.sprite.pos.y <= 0 {
		dir = math.Pi / 2
	} else if c.sprite.pos.y+c.sprite.imageHeight >= screenHeight {
		dir = -math.Pi / 2
	}

	return dir
}

func (c *Cow) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(CowSpeedMultiplier)
	c.sprite.Move(velocity)
}

func CreateRandomCow(img ebiten.Image) *Cow {
	w, h := img.Size()
	return &Cow{
		sprite: &Sprite{
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(rand.Intn(screenWidth)), float64(rand.Intn(screenHeight))},
		},
		velocity: &Vector{
			dir: 2 * math.Pi * rand.Float64(),
			len: 1,
		},
	}
}
