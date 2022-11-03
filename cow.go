package main

import (
	"math"
	"math/rand"
)

const (
	CowSpeedMultiplier    = 2
	FarmerDetectionRadius = 100
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

	// Flees from farmer.
	if c.sprite.pos.WithinRadius(f.sprite.Center(), FarmerDetectionRadius) {
		dir = VectorFromPoints(f.sprite.Center(), c.sprite.Center()).dir
	}

	// Bounded at edges.
	// Boundaries at left and right borders.
	if c.sprite.pos.x < 0 {
		dir = 0
	} else if mx := screenWidth - c.sprite.imageWidth; mx <= c.sprite.pos.x {
		dir = math.Pi
	}

	// Boundaries at top and bottom borders.
	if c.sprite.pos.y < 0 {
		dir = math.Pi / 2
	} else if my := screenHeight - c.sprite.imageHeight; my <= c.sprite.pos.y {
		dir = -math.Pi / 2
	}

	return dir
}

func (c *Cow) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(CowSpeedMultiplier)
	c.sprite.Move(velocity)
}
