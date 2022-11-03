package main

import (
	"math"
	"math/rand"
)

const (
	CowSpeedMultiplier = 2
)

type Cow struct {
	sprite   *Sprite
	velocity *Vector
}

func (c *Cow) Update() {
	// Update position.
	c.Move(*c.velocity)

	// Small chance to choose new direction.
	if rand.Float64() >= 0.95 {
		c.velocity.Rotate((math.Pi / 2) * (0.5 - rand.Float64()))
	}

	c.sprite.Update()
}

func (c *Cow) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(CowSpeedMultiplier)
	c.sprite.Move(velocity)
}
