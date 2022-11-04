package main

type Actor struct {
	sprite   *Sprite
	velocity *Vector
}

func (a *Actor) Shunt() {
	// moves actor to just inside borders if they are at or outside border
	// TODO also check for overlap with collidable objects

	if a.sprite.pos.x <= 0 {
		a.sprite.pos.x = 1
	} else if a.sprite.pos.x+a.sprite.imageWidth >= screenWidth {
		a.sprite.pos.x = screenWidth - a.sprite.imageWidth - 1
	}

	if a.sprite.pos.y <= 0 {
		a.sprite.pos.y = 1
	} else if a.sprite.pos.y+a.sprite.imageHeight >= screenHeight {
		a.sprite.pos.y = screenHeight - a.sprite.imageHeight - 1
	}
}

func (g *Game) MoveActor(a Actor, v Vector, speedMultiplier float64) {
	v.Normalize()
	v.Scale(speedMultiplier)
	v = g.CheckMovementActor(a, v)
	a.sprite.Move(v)
}

func (g *Game) CheckMovementActor(a Actor, v Vector) Vector {
	// check if move would result in collision
	newSpriteX := &Sprite{
		imageWidth:  a.sprite.imageWidth,
		imageHeight: a.sprite.imageHeight,
		pos:         &Coordinate{a.sprite.pos.x + v.X(), a.sprite.pos.y},
	}
	newSpriteY := &Sprite{
		imageWidth:  a.sprite.imageWidth,
		imageHeight: a.sprite.imageHeight,
		pos:         &Coordinate{a.sprite.pos.x, a.sprite.pos.y + v.Y()},
	}
	newSpriteXY := &Sprite{
		imageWidth:  a.sprite.imageWidth,
		imageHeight: a.sprite.imageHeight,
		pos:         &Coordinate{a.sprite.pos.x + v.X(), a.sprite.pos.y + v.Y()},
	}

	validXY := !g.CheckCollision(newSpriteXY)
	validX := !g.CheckCollision(newSpriteX)
	validY := !g.CheckCollision(newSpriteY)

	if !validXY {
		if validX && !validY {
			v.SetXY(Coordinate{v.X(), 0})
		} else if validY && !validX {
			v.SetXY(Coordinate{0, v.Y()})
		} else {
			v = Vector{0, 0}
		}
	}

	return v
}
