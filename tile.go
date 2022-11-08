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
	tileTreeTrunkTLImage := PrepareImage("./assets/tiles/tree_trunk_TL.png", op)
	tileTreeTrunkTRImage := PrepareImage("./assets/tiles/tree_trunk_TR.png", op)
	tileTreeTrunkBLImage := PrepareImage("./assets/tiles/tree_trunk_BL.png", op)
	tileTreeTrunkBRImage := PrepareImage("./assets/tiles/tree_trunk_BR.png", op)

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

	tileGroupIceHole := map[int]Tile{
		0: {
			image:      tileIceHoleTLImage,
			collidable: true,
		},
		1: {
			image:      tileIceHoleTRImage,
			collidable: true,
		},
		2: {
			image:      tileIceHoleBLImage,
			collidable: true,
		},
		3: {
			image:      tileIceHoleBRImage,
			collidable: true,
		},
	}

	tileGroupTreeTrunk := map[int]Tile{
		0: {
			image:      tileTreeTrunkTLImage,
			collidable: true,
		},
		1: {
			image:      tileTreeTrunkTRImage,
			collidable: true,
		},
		2: {
			image:      tileTreeTrunkBLImage,
			collidable: true,
		},
		3: {
			image:      tileTreeTrunkBRImage,
			collidable: true,
		},
	}

	for i := 0; i < 5; i++ {
		tiles = AddTileGroup(tiles, RandomUnoccupiedTileCoordinate(tiles, 2, 4), 2, tileGroupIceHole)
	}

	tiles = AddTileGroup(tiles, RandomUnoccupiedTileCoordinate(tiles, 2, 4), 2, tileGroupTreeTrunk)

	return tiles
}

func AddTileGroup(tiles map[TileCoordinate]*Tile, startPos TileCoordinate, width int, tileGroup map[int]Tile) map[TileCoordinate]*Tile {
	// startPos: top left tile coordinate of tile group
	// width: width of tile group in tiles
	// images: slice of images for populating tile group, ordered from left to right, top to bottom
	// collidable: true if all tiles in group are collidable, false if all tiles are not collidable
	if len(tileGroup)%width != 0 {
		panic("Tile group not rectangular")
	}

	for index, tile := range tileGroup {
		dx := index % width
		dy := (index - dx) / width
		tiles[TileCoordinate{x: startPos.x + dx, y: startPos.y + dy}] = &Tile{image: tile.image, collidable: tile.collidable}
	}

	return tiles
}

func RandomTileCoordinate() TileCoordinate {
	tileCountX := screenWidth / TileSize
	tileCountY := screenHeight / TileSize

	return TileCoordinate{x: rand.Intn(tileCountX), y: rand.Intn(tileCountY)}
}

func RandomUnoccupiedTileCoordinate(tiles map[TileCoordinate]*Tile, width int, count int) TileCoordinate {
	tileCoord := RandomTileCoordinate()

	for {
		occupied := false
		for i := 0; i < count; i++ {
			dx := i % width
			dy := (i - dx) / width
			currentCoord := TileCoordinate{x: tileCoord.x + dx, y: tileCoord.y + dy}
			if tiles[currentCoord].collidable {
				occupied = true
			}
		}

		if !occupied {
			break
		}

		tileCoord = RandomTileCoordinate()
	}

	return tileCoord
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
