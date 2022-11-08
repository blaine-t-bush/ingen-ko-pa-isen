package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSize = 20
)

type TileCoordinate struct {
	x int
	y int
}

type Tile struct {
	image *ebiten.Image
}

func GenerateTiles() map[TileCoordinate]*Tile {
	// Create slice of tile coordinates based on screen dimensions and tile size.
	tileCoordinates := []TileCoordinate{}
	tileCountX := screenWidth / TileSize
	tileCountY := screenHeight / TileSize

	for x := 0; x < tileCountX; x++ {
		for y := 0; y < tileCountY; y++ {
			tileCoordinates = append(tileCoordinates, TileCoordinate{x, y})
		}
	}

	// Prepare images that will be used in tiles.
	op := &ebiten.DrawImageOptions{}
	tileIceImage := PrepareImage("./assets/tiles/ice.png", op)
	tileIceStreaksImage := PrepareImage("./assets/tiles/ice_streaks.png", op)

	// Use slice of tile coordinates to create map of tile coordinates to tiles.
	tiles := map[TileCoordinate]*Tile{}
	for _, coord := range tileCoordinates {
		if rand.Float64() <= 0.05 {
			tiles[coord] = &Tile{image: tileIceStreaksImage}
		} else {
			tiles[coord] = &Tile{image: tileIceImage}
		}
	}

	return tiles
}
