package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FootprintSpacing      = 10
	FootprintLifetimeMs   = 3 // in seconds
	MomentumScalingFactor = 10
	SpeedMax              = 3
)

type Actor struct {
	id                         string
	image                      *ebiten.Image
	pos                        *ScreenCoordinate
	width                      float64
	height                     float64
	distanceSinceLastFootprint *float64
	speedMax                   float64
	speedMultiplier            float64
	velocityActual             *Vector
	velocityDesired            *Vector
	noiseSpeed                 *Noise
	noiseDir                   *Noise
}

func (a Actor) BoundingBox() BoundingBox {
	return BoundingBox{pos: *a.pos, width: a.width, height: a.height}
}

func (g *Game) MoveActor(a Actor, velocityDesired Vector) {
	velocityActual := g.CheckMovementActor(a, velocityDesired)
	a.velocityActual.SetXY(ScreenCoordinate{x: velocityActual.X(), y: velocityActual.Y()})
	if !WithinEps(velocityActual.len, eps) {
		a.Move(velocityActual)
		g.AddFootprint(a, velocityActual)
	}
}

func (g *Game) CheckMovementActor(a Actor, velocityDesired Vector) Vector {
	// check if actor is on ice. this may result in sliding.
	// if on ice, translation is a combination of the input velocity and existing velocity.
	// i.e. momentum is conserved on ice which results in sliding.
	var newVelocity Vector
	if g.CoordinateIsOnTerrainType(ScreenCoordinate{a.BoundingBox().CenterX(), a.BoundingBox().Bottom()}, TerrainTypeIce) {
		newVelocity.SetXY(ScreenCoordinate{x: velocityDesired.X()/MomentumScalingFactor + a.velocityActual.X(), y: velocityDesired.Y()/MomentumScalingFactor + a.velocityActual.Y()})
	} else {
		newVelocity.SetXY(ScreenCoordinate{x: velocityDesired.X(), y: velocityDesired.Y()})
	}
	newVelocity.SetXY(ScreenCoordinate{x: velocityDesired.X(), y: velocityDesired.Y()})
	newVelocity.BoundLength(a.speedMax)

	// check collision
	collisionX := playerObj.Check(newVelocity.X(), 0)
	if collisionX != nil {
		newVelocity.SetXY(ScreenCoordinate{x: collisionX.ContactWithObject(collisionX.Objects[0]).X(), y: newVelocity.Y()})
	}
	playerObj.X += newVelocity.X()
	playerObj.Update()

	collisionY := playerObj.Check(0, newVelocity.Y())
	if collisionY != nil {
		newVelocity.SetXY(ScreenCoordinate{x: newVelocity.X(), y: collisionY.ContactWithObject(collisionY.Objects[0]).Y()})
	}
	playerObj.Y += newVelocity.Y()
	playerObj.Update()

	// // check if move would result in collision
	// newPosX := &ScreenCoordinate{a.pos.x + newVelocity.X(), a.pos.y}
	// newPosY := &ScreenCoordinate{a.pos.x, a.pos.y + newVelocity.Y()}
	// newPosXY := &ScreenCoordinate{a.pos.x + newVelocity.X(), a.pos.y + newVelocity.Y()}

	// validX := !g.CheckCollision(BoundingBox{pos: *newPosX, width: a.width, height: a.height})
	// validY := !g.CheckCollision(BoundingBox{pos: *newPosY, width: a.width, height: a.height})
	// validXY := !g.CheckCollision(BoundingBox{pos: *newPosXY, width: a.width, height: a.height})

	// if !validXY {
	// 	if validX && !validY {
	// 		newVelocity.RemoveY()
	// 	} else if validY && !validX {
	// 		newVelocity.RemoveX()
	// 	} else {
	// 		newVelocity = Vector{0, 0}
	// 	}
	// }

	return newVelocity
}

func (a *Actor) Move(velocity Vector) {
	a.pos.Translate(velocity)
	a.velocityActual = &velocity
}

func (a *Actor) Shunt() {
	// moves actor to just inside borders if they are at or outside border
	// TODO also check for overlap with collidable objects and tiles

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
