package main

type Actor struct {
	sprite   *Sprite
	velocity *Vector
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
