package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) HandleKeyPresses() {
	// Determine direction
	var dx, dy float64 = 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = 1
	}

	// Normalize speed
	var dxNorm, dyNorm float64 = 0, 0
	speed := math.Sqrt(dx*dx + dy*dy)
	if speed > 0 {
		dxNorm = dx / speed
		dyNorm = dy / speed
	}

	// Scale to match player speed setting
	g.MovePlayer(dxNorm*PlayerSpeed, dyNorm*PlayerSpeed)
}
