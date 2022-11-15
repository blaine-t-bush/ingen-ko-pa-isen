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

func (g *Game) RandomCoordinateOfTerrainType(terrainType int) ScreenCoordinate {
	tileCoordinates := []TileCoordinate{}

	for coord, tile := range g.tiles {
		if tile.terrainType == terrainType {
			tileCoordinates = append(tileCoordinates, coord)
		}
	}

	tileCoordinate := tileCoordinates[rand.Intn(len(tileCoordinates)-1)]
	x := tileCoordinate.ToScreenCoordinate().x + float64(rand.Intn(TileSize))
	y := tileCoordinate.ToScreenCoordinate().y + float64(rand.Intn(TileSize))

	return ScreenCoordinate{x: x, y: y}
}

func (g *Game) CoordinateIsOnTerrainType(c ScreenCoordinate, terrainType int) bool {
	tileCoordinate := c.ToTileCoordinate()
	if tile, exists := g.tiles[tileCoordinate]; exists {
		return tile.terrainType == terrainType
	} else {
		return false
	}
}
