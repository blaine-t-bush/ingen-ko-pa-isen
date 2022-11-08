package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	image      *ebiten.Image
	pos        *ScreenCoordinate
	width      float64
	height     float64
	velocity   *Vector
	noiseSpeed *Noise
	noiseDir   *Noise
}

func (a Actor) BoundingBox() BoundingBox {
	return BoundingBox{pos: *a.pos, width: a.width, height: a.height}
}

func (a *Actor) Move(offset Vector) {
	a.pos.Translate(offset)
}

func (a *Actor) Shunt() {
	// moves actor to just inside borders if they are at or outside border
	// TODO also check for overlap with collidable objects

	collidesTop, overlapTop := a.BoundingBox().CollidesWithTopBorder()
	collidesBottom, overlapBottom := a.BoundingBox().CollidesWithBottomBorder()
	collidesLeft, overlapLeft := a.BoundingBox().CollidesWithLeftBorder()
	collidesRight, overlapRight := a.BoundingBox().CollidesWithRightBorder()

	if collidesTop {
		a.Move(Vector{len: overlapTop + 1, dir: math.Pi / 2})
	} else if collidesBottom {
		a.Move(Vector{len: overlapBottom + 1, dir: 3 * math.Pi / 2})
	}

	if collidesLeft {
		a.Move(Vector{len: overlapLeft + 1, dir: 0})
	} else if collidesRight {
		a.Move(Vector{len: overlapRight + 1, dir: math.Pi})
	}
}

func (g *Game) MoveActor(a Actor, v Vector, speedMultiplier float64) {
	v = g.CheckMovementActor(a, v)
	a.Move(v)
}

func (g *Game) CheckMovementActor(a Actor, v Vector) Vector {
	// check if move would result in collision
	newPosX := &ScreenCoordinate{a.pos.x + v.X(), a.pos.y}
	newPosY := &ScreenCoordinate{a.pos.x, a.pos.y + v.Y()}
	newPosXY := &ScreenCoordinate{a.pos.x + v.X(), a.pos.y + v.Y()}

	validX := !g.CheckCollision(BoundingBox{pos: *newPosX, width: a.width, height: a.height})
	validY := !g.CheckCollision(BoundingBox{pos: *newPosY, width: a.width, height: a.height})
	validXY := !g.CheckCollision(BoundingBox{pos: *newPosXY, width: a.width, height: a.height})

	if !validXY {
		if validX && !validY {
			v.RemoveY()
		} else if validY && !validX {
			v.RemoveX()
		} else {
			v = Vector{0, 0}
		}
	}

	return v
}
