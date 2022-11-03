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
	farmerPos := Coordinate{g.farmer.sprite.x, g.farmer.sprite.y}
	x, y := ebiten.CursorPosition()
	mousePos := Coordinate{float64(x), float64(y)}
	velocity := VectorFromPoints(farmerPos, mousePos)
	g.farmer.Move(velocity)
	return nil
}

func (g *Game) HandleKeyPresses(keys []ebiten.Key) error {
	pressedLeft := KeysIncludes(keys, ebiten.KeyArrowLeft) || KeysIncludes(keys, ebiten.KeyA)
	pressedRight := KeysIncludes(keys, ebiten.KeyArrowRight) || KeysIncludes(keys, ebiten.KeyD)
	pressedUp := KeysIncludes(keys, ebiten.KeyArrowUp) || KeysIncludes(keys, ebiten.KeyW)
	pressedDown := KeysIncludes(keys, ebiten.KeyArrowDown) || KeysIncludes(keys, ebiten.KeyS)

	if pressedLeft || pressedRight || pressedUp || pressedDown {
		var v Vector
		if pressedLeft && pressedUp {
			v = VectorFromXY(Coordinate{-1, -1})
		} else if pressedLeft && pressedDown {
			v = VectorFromXY(Coordinate{-1, 1})
		} else if pressedLeft {
			v = VectorFromXY(Coordinate{-1, 0})
		} else if pressedRight && pressedUp {
			v = VectorFromXY(Coordinate{1, -1})
		} else if pressedRight && pressedDown {
			v = VectorFromXY(Coordinate{1, 1})
		} else if pressedRight {
			v = VectorFromXY(Coordinate{1, 0})
		} else if pressedUp {
			v = VectorFromXY(Coordinate{0, -1})
		} else if pressedDown {
			v = VectorFromXY(Coordinate{0, 1})
		}

		g.farmer.Move(v)
	}

	return nil
}
