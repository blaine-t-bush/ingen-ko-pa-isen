package main

import "github.com/hajimehoshi/ebiten/v2"

var ValidInputKeys []ebiten.Key = []ebiten.Key{
	ebiten.KeyArrowLeft,
	ebiten.KeyArrowRight,
	ebiten.KeyArrowUp,
	ebiten.KeyArrowDown,
}

func KeysIncludes(keys []ebiten.Key, includes ebiten.Key) bool {
	for _, key := range keys {
		if key == includes {
			return true
		}
	}

	return false
}

func (g *Game) HandleMouseLeftClick() error {
	farmerPos := g.farmer.BoundingBox().Center()
	x, y := ebiten.CursorPosition()
	mousePos := ScreenCoordinate{float64(x), float64(y)}
	if !mousePos.WithinRadius(farmerPos, 2) {
		velocity := VectorFromPoints(farmerPos, mousePos)
		velocity.Normalize()
		velocity.Scale(g.farmer.speedMultiplier)
		g.farmer.velocityDesired = &velocity
	}

	return nil
}

func (g *Game) HandleKeyPresses(keys []ebiten.Key) error {
	pressedLeft := KeysIncludes(keys, ebiten.KeyArrowLeft) || KeysIncludes(keys, ebiten.KeyA)
	pressedRight := KeysIncludes(keys, ebiten.KeyArrowRight) || KeysIncludes(keys, ebiten.KeyD)
	pressedUp := KeysIncludes(keys, ebiten.KeyArrowUp) || KeysIncludes(keys, ebiten.KeyW)
	pressedDown := KeysIncludes(keys, ebiten.KeyArrowDown) || KeysIncludes(keys, ebiten.KeyS)

	if pressedLeft || pressedRight || pressedUp || pressedDown {
		var velocity Vector
		if pressedLeft && pressedUp {
			velocity = VectorFromXY(ScreenCoordinate{-1, -1})
		} else if pressedLeft && pressedDown {
			velocity = VectorFromXY(ScreenCoordinate{-1, 1})
		} else if pressedLeft {
			velocity = VectorFromXY(ScreenCoordinate{-1, 0})
		} else if pressedRight && pressedUp {
			velocity = VectorFromXY(ScreenCoordinate{1, -1})
		} else if pressedRight && pressedDown {
			velocity = VectorFromXY(ScreenCoordinate{1, 1})
		} else if pressedRight {
			velocity = VectorFromXY(ScreenCoordinate{1, 0})
		} else if pressedUp {
			velocity = VectorFromXY(ScreenCoordinate{0, -1})
		} else if pressedDown {
			velocity = VectorFromXY(ScreenCoordinate{0, 1})
		}

		velocity.Normalize()
		velocity.Scale(g.farmer.speedMultiplier)
		g.farmer.velocityDesired = &velocity
	}

	return nil
}
