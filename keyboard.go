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

func (g *Game) HandleKeyPresses(keys []ebiten.Key) error {
	pressedLeft := KeysIncludes(keys, ebiten.KeyArrowLeft)
	pressedRight := KeysIncludes(keys, ebiten.KeyArrowRight)
	pressedUp := KeysIncludes(keys, ebiten.KeyArrowUp)
	pressedDown := KeysIncludes(keys, ebiten.KeyArrowDown)

	if pressedLeft && pressedUp {
		g.farmer.Move(DirectionLeftUp)
	} else if pressedLeft && pressedDown {
		g.farmer.Move(DirectionLeftDown)
	} else if pressedLeft {
		g.farmer.Move(DirectionLeft)
	} else if pressedRight && pressedUp {
		g.farmer.Move(DirectionRightUp)
	} else if pressedRight && pressedDown {
		g.farmer.Move(DirectionRightDown)
	} else if pressedRight {
		g.farmer.Move(DirectionRight)
	} else if pressedUp {
		g.farmer.Move(DirectionUp)
	} else if pressedDown {
		g.farmer.Move(DirectionDown)
	}

	return nil
}
