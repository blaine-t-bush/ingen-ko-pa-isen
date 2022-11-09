package main

import (
	"math"
	"math/rand"
)

type ScreenCoordinate struct {
	x float64
	y float64
}

func (c *ScreenCoordinate) Translate(offset Vector) {
	c.x += offset.X()
	c.y += offset.Y()
}

func (c ScreenCoordinate) DistanceFrom(from ScreenCoordinate) float64 {
	dX := from.x - c.x
	dY := from.y - c.y
	return math.Sqrt(dX*dX + dY*dY)
}

func (c ScreenCoordinate) WithinRadius(center ScreenCoordinate, r float64) bool {
	return c.DistanceFrom(center) <= r
}

func RandomCoordinate(w, h int) ScreenCoordinate {
	return ScreenCoordinate{x: float64(rand.Intn(screenWidth - w)), y: float64(rand.Intn(screenHeight - h))}
}

func RandomSnowCoordinate(w, h int) ScreenCoordinate {
	var minX, minY, maxX, maxY int

	if rand.Intn(2) == 0 {
		minX = 0
		maxX = SnowBorderSize * TileSize
	} else {
		minX = screenWidth - (SnowBorderSize * TileSize) - w
		maxX = screenWidth - w
	}

	if rand.Intn(2) == 0 {
		minY = 0
		maxY = SnowBorderSize * TileSize
	} else {
		minY = screenHeight - (SnowBorderSize * TileSize) - h
		maxY = screenHeight - h
	}

	return ScreenCoordinate{x: float64(minX + rand.Intn(maxX-minX)), y: float64(minY + rand.Intn(maxY-minY))}
}

func RandomIceCoordinate(w, h int) ScreenCoordinate {
	minX := SnowBorderSize * TileSize
	minY := SnowBorderSize * TileSize
	maxX := screenWidth - (SnowBorderSize * TileSize) - w
	maxY := screenHeight - (SnowBorderSize * TileSize) - h

	return ScreenCoordinate{x: float64(minX + rand.Intn(maxX-minX)), y: float64(minY + rand.Intn(maxY-minY))}
}
