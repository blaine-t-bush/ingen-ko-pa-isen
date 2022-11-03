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

	// Bounded at edgec.
	// Boundaries at left and right borderc.
	if c.sprite.pos.x < 0 {
		c.sprite.pos.x = -c.sprite.pos.x
		c.velocity.dir = 0
	} else if mx := screenWidth - c.sprite.imageWidth; mx <= c.sprite.pos.x {
		c.sprite.pos.x = 2*mx - c.sprite.pos.x
		c.velocity.dir = math.Pi
	}

	// Boundaries at top and bottom borderc.
	if c.sprite.pos.y < 0 {
		c.sprite.pos.y = -c.sprite.pos.y
		c.velocity.dir = math.Pi / 2
	} else if my := screenHeight - c.sprite.imageHeight; my <= c.sprite.pos.y {
		c.sprite.pos.y = 2*my - c.sprite.pos.y
		c.velocity.dir = -math.Pi / 2
	}
}

func (c *Cow) Move(velocity Vector) {
	velocity.Normalize()
	velocity.Scale(CowSpeedMultiplier)
	c.sprite.Move(velocity)
}
