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

func (g *Game) HandleKeyPress(key ebiten.Key) error {
	switch key {
	case ebiten.KeyArrowLeft:
		g.farmer.UpdateSpeed(-1, 0)
	case ebiten.KeyArrowRight:
		g.farmer.UpdateSpeed(1, 0)
	case ebiten.KeyArrowDown:
		g.farmer.UpdateSpeed(0, 1)
	case ebiten.KeyArrowUp:
		g.farmer.UpdateSpeed(0, -1)
	}

	return nil
}
