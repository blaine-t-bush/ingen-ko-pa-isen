package main

import (
	"math"
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
	image      *ebiten.Image
	collidable bool
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
	tileIceHoleTLImage := PrepareImage("./assets/tiles/ice_hole_TL.png", op)
	tileIceHoleTRImage := PrepareImage("./assets/tiles/ice_hole_TR.png", op)
	tileIceHoleBLImage := PrepareImage("./assets/tiles/ice_hole_BL.png", op)
	tileIceHoleBRImage := PrepareImage("./assets/tiles/ice_hole_BR.png", op)

	// Use slice of tile coordinates to create map of tile coordinates to tiles.
	// Start with all ice, then add other things to spice it up.
	tiles := map[TileCoordinate]*Tile{}
	for _, coord := range tileCoordinates {
		if rand.Float64() <= 0.04 {
			tiles[coord] = &Tile{image: tileIceStreaksImage, collidable: false}
		} else {
			tiles[coord] = &Tile{image: tileIceImage, collidable: false}
		}
	}

	tiles[TileCoordinate{tileCountX / 2, tileCountY / 2}] = &Tile{image: tileIceHoleTLImage, collidable: true}
	tiles[TileCoordinate{tileCountX/2 + 1, tileCountY / 2}] = &Tile{image: tileIceHoleTRImage, collidable: true}
	tiles[TileCoordinate{tileCountX / 2, tileCountY/2 + 1}] = &Tile{image: tileIceHoleBLImage, collidable: true}
	tiles[TileCoordinate{tileCountX/2 + 1, tileCountY/2 + 1}] = &Tile{image: tileIceHoleBRImage, collidable: true}

	return tiles
}

func (t Tile) ToBoundingBox(c TileCoordinate) BoundingBox {
	return BoundingBox{pos: c.ToScreenCoordinate(), width: TileSize, height: TileSize}
}

func (c *ScreenCoordinate) TileCoordinateX() int {
	return int(math.Mod(c.x, TileSize))
}

func (c *ScreenCoordinate) TileCoordinateY() int {
	return int(math.Mod(c.y, TileSize))
}

func (c *ScreenCoordinate) ToTileCoordinate() TileCoordinate {
	return TileCoordinate{x: c.TileCoordinateX(), y: c.TileCoordinateY()}
}

func (t *TileCoordinate) ScreenCoordinateX() float64 {
	return float64(t.x * TileSize)
}

func (t *TileCoordinate) ScreenCoordinateY() float64 {
	return float64(t.y * TileSize)
}

func (t *TileCoordinate) ToScreenCoordinate() ScreenCoordinate {
	return ScreenCoordinate{x: t.ScreenCoordinateX(), y: t.ScreenCoordinateY()}
}
