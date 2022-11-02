package main

import "math/rand"

const (
	CowSpeedMultiplier = 2
)

type Cow struct {
	sprite *Sprite
	vx     int
	vy     int
}

func (c *Cow) Update() {
	// Update position.
	if c.vx < 0 && c.vy < 0 {
		c.sprite.Move(DirectionLeftUp, CowSpeedMultiplier)
	} else if c.vx < 0 && c.vy > 0 {
		c.sprite.Move(DirectionLeftDown, CowSpeedMultiplier)
	} else if c.vx < 0 && c.vy == 0 {
		c.sprite.Move(DirectionLeft, CowSpeedMultiplier)
	} else if c.vx > 0 && c.vy < 0 {
		c.sprite.Move(DirectionRightUp, CowSpeedMultiplier)
	} else if c.vx > 0 && c.vy > 0 {
		c.sprite.Move(DirectionRightDown, CowSpeedMultiplier)
	} else if c.vx > 0 && c.vy == 0 {
		c.sprite.Move(DirectionRight, CowSpeedMultiplier)
	} else if c.vy < 0 {
		c.sprite.Move(DirectionUp, CowSpeedMultiplier)
	} else if c.vy > 0 {
		c.sprite.Move(DirectionDown, CowSpeedMultiplier)
	}

	// Small chance to choose new direction.
	if rand.Float64() >= 0.995 {
		c.vx = 1 - rand.Intn(3)
		c.vy = 1 - rand.Intn(3)
	}

	c.sprite.Update()
}

func (c *Cow) Move(dir Direction) {
	c.sprite.Move(dir, CowSpeedMultiplier)
}
