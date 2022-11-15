package main

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSize       = 20
	SnowBorderSize = 5
	TerrainTypeIce = iota
	TerrainTypeSnow
	TileIce = iota
	TileIceWithStreaks
	TileSnow
	TileSnowWithSpeckles
	TileSnowIceTop
	TileSnowIceRight
	TileSnowIceBottom
	TileSnowIceLeft
	TileSnowIceStripHorizontal
	TileSnowIceStripVertical
	TileSnowIceJutTop
	TileSnowIceJutRight
	TileSnowIceJutBottom
	TileSnowIceJutLeft
	TileSnowIceInnerCornerTopRight
	TileSnowIceInnerCornerBottomRight
	TileSnowIceInnerCornerBottomLeft
	TileSnowIceInnerCornerTopLeft
	TileSnowIceOuterCornerTopRight
	TileSnowIceOuterCornerBottomRight
	TileSnowIceOuterCornerBottomLeft
	TileSnowIceOuterCornerTopLeft
	TileIceHoleTop
	TileIceHoleTopRight
	TileIceHoleRight
	TileIceHoleBottomRight
	TileIceHoleBottom
	TileIceHoleBottomLeft
	TileIceHoleLeft
	TileIceHoleTopLeft
)

var (
	TileSymbols = map[string]int{
		"I": TileIce,
		"S": TileSnow,
	}

	TileImage = map[int]*ebiten.Image{
		TileIce:                           PrepareImage("./assets/tiles/ice.png", &ebiten.DrawImageOptions{}),
		TileIceWithStreaks:                PrepareImage("./assets/tiles/ice_streaks.png", &ebiten.DrawImageOptions{}),
		TileSnow:                          PrepareImage("./assets/tiles/snow.png", &ebiten.DrawImageOptions{}),
		TileSnowWithSpeckles:              PrepareImage("./assets/tiles/snow_speckled.png", &ebiten.DrawImageOptions{}),
		TileSnowIceTop:                    PrepareImage("./assets/tiles/snow_ice_top.png", &ebiten.DrawImageOptions{}),
		TileSnowIceRight:                  PrepareImage("./assets/tiles/snow_ice_right.png", &ebiten.DrawImageOptions{}),
		TileSnowIceBottom:                 PrepareImage("./assets/tiles/snow_ice_bottom.png", &ebiten.DrawImageOptions{}),
		TileSnowIceLeft:                   PrepareImage("./assets/tiles/snow_ice_left.png", &ebiten.DrawImageOptions{}),
		TileSnowIceStripHorizontal:        PrepareImage("./assets/tiles/snow_ice_strip_horizontal.png", &ebiten.DrawImageOptions{}),
		TileSnowIceStripVertical:          PrepareImage("./assets/tiles/snow_ice_strip_vertical.png", &ebiten.DrawImageOptions{}),
		TileSnowIceJutTop:                 PrepareImage("./assets/tiles/snow_ice_jut_top.png", &ebiten.DrawImageOptions{}),
		TileSnowIceJutRight:               PrepareImage("./assets/tiles/snow_ice_jut_right.png", &ebiten.DrawImageOptions{}),
		TileSnowIceJutBottom:              PrepareImage("./assets/tiles/snow_ice_jut_bottom.png", &ebiten.DrawImageOptions{}),
		TileSnowIceJutLeft:                PrepareImage("./assets/tiles/snow_ice_jut_left.png", &ebiten.DrawImageOptions{}),
		TileSnowIceInnerCornerTopRight:    PrepareImage("./assets/tiles/snow_ice_inner_corner_topright.png", &ebiten.DrawImageOptions{}),
		TileSnowIceInnerCornerBottomRight: PrepareImage("./assets/tiles/snow_ice_inner_corner_bottomright.png", &ebiten.DrawImageOptions{}),
		TileSnowIceInnerCornerBottomLeft:  PrepareImage("./assets/tiles/snow_ice_inner_corner_bottomleft.png", &ebiten.DrawImageOptions{}),
		TileSnowIceInnerCornerTopLeft:     PrepareImage("./assets/tiles/snow_ice_inner_corner_topleft.png", &ebiten.DrawImageOptions{}),
		TileSnowIceOuterCornerTopRight:    PrepareImage("./assets/tiles/snow_ice_outer_corner_topright.png", &ebiten.DrawImageOptions{}),
		TileSnowIceOuterCornerBottomRight: PrepareImage("./assets/tiles/snow_ice_outer_corner_bottomright.png", &ebiten.DrawImageOptions{}),
		TileSnowIceOuterCornerBottomLeft:  PrepareImage("./assets/tiles/snow_ice_outer_corner_bottomleft.png", &ebiten.DrawImageOptions{}),
		TileSnowIceOuterCornerTopLeft:     PrepareImage("./assets/tiles/snow_ice_outer_corner_topleft.png", &ebiten.DrawImageOptions{}),
		TileIceHoleTop:                    PrepareImage("./assets/tiles/ice_hole_T.png", &ebiten.DrawImageOptions{}),
		TileIceHoleTopRight:               PrepareImage("./assets/tiles/ice_hole_TR.png", &ebiten.DrawImageOptions{}),
		TileIceHoleRight:                  PrepareImage("./assets/tiles/ice_hole_R.png", &ebiten.DrawImageOptions{}),
		TileIceHoleBottomRight:            PrepareImage("./assets/tiles/ice_hole_BR.png", &ebiten.DrawImageOptions{}),
		TileIceHoleBottom:                 PrepareImage("./assets/tiles/ice_hole_B.png", &ebiten.DrawImageOptions{}),
		TileIceHoleBottomLeft:             PrepareImage("./assets/tiles/ice_hole_BL.png", &ebiten.DrawImageOptions{}),
		TileIceHoleLeft:                   PrepareImage("./assets/tiles/ice_hole_L.png", &ebiten.DrawImageOptions{}),
		TileIceHoleTopLeft:                PrepareImage("./assets/tiles/ice_hole_TL.png", &ebiten.DrawImageOptions{}),
	}

	TileCollidable = map[int]bool{
		TileIce:                           false,
		TileIceWithStreaks:                false,
		TileSnow:                          false,
		TileSnowWithSpeckles:              false,
		TileSnowIceTop:                    false,
		TileSnowIceRight:                  false,
		TileSnowIceBottom:                 false,
		TileSnowIceLeft:                   false,
		TileSnowIceStripHorizontal:        false,
		TileSnowIceStripVertical:          false,
		TileSnowIceJutTop:                 false,
		TileSnowIceJutRight:               false,
		TileSnowIceJutBottom:              false,
		TileSnowIceJutLeft:                false,
		TileSnowIceInnerCornerTopRight:    false,
		TileSnowIceInnerCornerBottomRight: false,
		TileSnowIceInnerCornerBottomLeft:  false,
		TileSnowIceInnerCornerTopLeft:     false,
		TileSnowIceOuterCornerTopRight:    false,
		TileSnowIceOuterCornerBottomRight: false,
		TileSnowIceOuterCornerBottomLeft:  false,
		TileSnowIceOuterCornerTopLeft:     false,
		TileIceHoleTop:                    true,
		TileIceHoleTopRight:               true,
		TileIceHoleRight:                  true,
		TileIceHoleBottomRight:            true,
		TileIceHoleBottom:                 true,
		TileIceHoleBottomLeft:             true,
		TileIceHoleLeft:                   true,
		TileIceHoleTopLeft:                true,
	}

	TileTerrainType = map[int]int{
		TileIce:                           TerrainTypeIce,
		TileIceWithStreaks:                TerrainTypeIce,
		TileSnow:                          TerrainTypeSnow,
		TileSnowWithSpeckles:              TerrainTypeSnow,
		TileSnowIceTop:                    TerrainTypeSnow,
		TileSnowIceRight:                  TerrainTypeSnow,
		TileSnowIceBottom:                 TerrainTypeSnow,
		TileSnowIceLeft:                   TerrainTypeSnow,
		TileSnowIceStripHorizontal:        TerrainTypeSnow,
		TileSnowIceStripVertical:          TerrainTypeSnow,
		TileSnowIceJutTop:                 TerrainTypeSnow,
		TileSnowIceJutRight:               TerrainTypeSnow,
		TileSnowIceJutBottom:              TerrainTypeSnow,
		TileSnowIceJutLeft:                TerrainTypeSnow,
		TileSnowIceInnerCornerTopRight:    TerrainTypeSnow,
		TileSnowIceInnerCornerBottomRight: TerrainTypeSnow,
		TileSnowIceInnerCornerBottomLeft:  TerrainTypeSnow,
		TileSnowIceInnerCornerTopLeft:     TerrainTypeSnow,
		TileSnowIceOuterCornerTopRight:    TerrainTypeSnow,
		TileSnowIceOuterCornerBottomRight: TerrainTypeSnow,
		TileSnowIceOuterCornerBottomLeft:  TerrainTypeSnow,
		TileSnowIceOuterCornerTopLeft:     TerrainTypeSnow,
		TileIceHoleTop:                    TerrainTypeIce,
		TileIceHoleTopRight:               TerrainTypeIce,
		TileIceHoleRight:                  TerrainTypeIce,
		TileIceHoleBottomRight:            TerrainTypeIce,
		TileIceHoleBottom:                 TerrainTypeIce,
		TileIceHoleBottomLeft:             TerrainTypeIce,
		TileIceHoleLeft:                   TerrainTypeIce,
		TileIceHoleTopLeft:                TerrainTypeIce,
	}
)

type TileCoordinate struct {
	x int
	y int
}

type Tile struct {
	image       *ebiten.Image
	collidable  bool
	terrainType int
}

func ReadMap() map[TileCoordinate]*Tile {
	file, err := os.Open("./maps/default.map")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tiles := map[TileCoordinate]*Tile{}
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		for colIndex, symbol := range scanner.Text() {
			if string(symbol) == "\n" {
				continue
			}
			tileNumber := TileSymbols[string(symbol)]
			tileCoordinate := TileCoordinate{colIndex, rowIndex}
			tiles[tileCoordinate] = &Tile{
				image:       TileImage[tileNumber],
				collidable:  TileCollidable[tileNumber],
				terrainType: TileTerrainType[tileNumber],
			}
		}
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Update tiles based on their neighbors.
	for coord, tile := range tiles {
		if coord.x > 0 && coord.x < screenWidth/TileSize-1 && coord.y > 0 && coord.y < screenHeight/TileSize-1 {
			if tile.terrainType == TerrainTypeSnow {
				surroundingCoords := []TileCoordinate{
					{coord.x, coord.y - 1},
					{coord.x + 1, coord.y - 1},
					{coord.x + 1, coord.y},
					{coord.x + 1, coord.y + 1},
					{coord.x, coord.y + 1},
					{coord.x - 1, coord.y + 1},
					{coord.x - 1, coord.y},
					{coord.x - 1, coord.y - 1},
				}

				surroundingIceCount := 0
				for _, surroundingCoord := range surroundingCoords {
					if tiles[surroundingCoord].terrainType == TerrainTypeIce {
						surroundingIceCount++
					}
				}

				iceAbove := tiles[TileCoordinate{coord.x, coord.y - 1}].terrainType == TerrainTypeIce
				iceAboveRight := tiles[TileCoordinate{coord.x + 1, coord.y - 1}].terrainType == TerrainTypeIce
				iceRight := tiles[TileCoordinate{coord.x + 1, coord.y}].terrainType == TerrainTypeIce
				iceBelowRight := tiles[TileCoordinate{coord.x + 1, coord.y + 1}].terrainType == TerrainTypeIce
				iceBelow := tiles[TileCoordinate{coord.x, coord.y + 1}].terrainType == TerrainTypeIce
				iceBelowLeft := tiles[TileCoordinate{coord.x - 1, coord.y + 1}].terrainType == TerrainTypeIce
				iceLeft := tiles[TileCoordinate{coord.x - 1, coord.y}].terrainType == TerrainTypeIce
				iceAboveLeft := tiles[TileCoordinate{coord.x - 1, coord.y - 1}].terrainType == TerrainTypeIce

				if surroundingIceCount != 0 {
					var newTileNumber int

					if iceAbove && iceAboveLeft && iceLeft && iceBelowLeft && iceBelow {
						newTileNumber = TileSnowIceJutRight
					} else if iceLeft && iceBelowLeft && iceBelow && iceBelowRight && iceRight {
						newTileNumber = TileSnowIceJutTop
					} else if iceBelow && iceBelowRight && iceRight && iceAboveRight && iceAbove {
						newTileNumber = TileSnowIceJutLeft
					} else if iceRight && iceAboveRight && iceAbove && iceAboveLeft && iceLeft {
						newTileNumber = TileSnowIceJutBottom
					} else if iceRight && iceBelow && !iceAbove && !iceLeft {
						newTileNumber = TileSnowIceOuterCornerBottomRight
					} else if !iceRight && iceBelow && !iceAbove && iceLeft {
						newTileNumber = TileSnowIceOuterCornerBottomLeft
					} else if !iceRight && !iceBelow && iceAbove && iceLeft {
						newTileNumber = TileSnowIceOuterCornerTopLeft
					} else if iceRight && !iceBelow && iceAbove && !iceLeft {
						newTileNumber = TileSnowIceOuterCornerTopRight
					} else if !iceAbove && !iceAboveRight && !iceRight && !iceBelowRight && !iceBelow && iceBelowLeft && !iceLeft && !iceAboveLeft {
						newTileNumber = TileSnowIceInnerCornerBottomLeft
					} else if !iceAbove && !iceAboveRight && !iceRight && !iceBelowRight && !iceBelow && !iceBelowLeft && !iceLeft && iceAboveLeft {
						newTileNumber = TileSnowIceInnerCornerTopLeft
					} else if !iceAbove && iceAboveRight && !iceRight && !iceBelowRight && !iceBelow && !iceBelowLeft && !iceLeft && !iceAboveLeft {
						newTileNumber = TileSnowIceInnerCornerTopRight
					} else if !iceAbove && !iceAboveRight && !iceRight && iceBelowRight && !iceBelow && !iceBelowLeft && !iceLeft && !iceAboveLeft {
						newTileNumber = TileSnowIceInnerCornerBottomRight
					} else if iceAbove && iceBelow {
						newTileNumber = TileSnowIceStripHorizontal
					} else if iceLeft && iceRight {
						newTileNumber = TileSnowIceStripVertical
					} else if iceAbove {
						newTileNumber = TileSnowIceTop
					} else if iceRight {
						newTileNumber = TileSnowIceRight
					} else if iceBelow {
						newTileNumber = TileSnowIceBottom
					} else if iceLeft {
						newTileNumber = TileSnowIceLeft
					} else {
						newTileNumber = TileSnow
					}

					tiles[coord] = &Tile{
						image:       TileImage[newTileNumber],
						collidable:  TileCollidable[newTileNumber],
						terrainType: TileTerrainType[newTileNumber],
					}
				} else {
					if rand.Float64() <= 0.2 {
						newTileNumber := TileSnowWithSpeckles
						tiles[coord] = &Tile{
							image:       TileImage[newTileNumber],
							collidable:  TileCollidable[newTileNumber],
							terrainType: TileTerrainType[newTileNumber],
						}
					}
				}
			} else if tile.terrainType == TerrainTypeIce {
				if rand.Float64() <= 0.02 {
					newTileNumber := TileIceWithStreaks
					tiles[coord] = &Tile{
						image:       TileImage[newTileNumber],
						collidable:  TileCollidable[newTileNumber],
						terrainType: TileTerrainType[newTileNumber],
					}
				}
			}
		}
	}

	return tiles
}

func GenerateTiles() map[TileCoordinate]*Tile {
	// // Create slice of tile coordinates based on screen dimensions and tile size.
	// tileCoordinates := []TileCoordinate{}
	// tileCountX := screenWidth / TileSize
	// tileCountY := screenHeight / TileSize

	// for x := 0; x < tileCountX; x++ {
	// 	for y := 0; y < tileCountY; y++ {
	// 		tileCoordinates = append(tileCoordinates, TileCoordinate{x, y})
	// 	}
	// }

	// // Prepare images that will be used in tiles.
	// op := &ebiten.DrawImageOptions{}
	// tileIceImage := PrepareImage("./assets/tiles/ice.png", op)
	// tileIceStreaksImage := PrepareImage("./assets/tiles/ice_streaks.png", op)
	// tileIceHoleTLImage := PrepareImage("./assets/tiles/ice_hole_TL.png", op)
	// tileIceHoleTRImage := PrepareImage("./assets/tiles/ice_hole_TR.png", op)
	// tileIceHoleBLImage := PrepareImage("./assets/tiles/ice_hole_BL.png", op)
	// tileIceHoleBRImage := PrepareImage("./assets/tiles/ice_hole_BR.png", op)
	// tileIceHoleTImage := PrepareImage("./assets/tiles/ice_hole_T.png", op)
	// tileIceHoleBImage := PrepareImage("./assets/tiles/ice_hole_B.png", op)

	// tileSnowToIceTImage := PrepareImage("./assets/tiles/snow_to_ice_T.png", op)
	// tileSnowToIceBImage := PrepareImage("./assets/tiles/snow_to_ice_B.png", op)
	// tileSnowToIceRImage := PrepareImage("./assets/tiles/snow_to_ice_R.png", op)
	// tileSnowToIceLImage := PrepareImage("./assets/tiles/snow_to_ice_L.png", op)

	// tileSnowToIceTRImage := PrepareImage("./assets/tiles/snow_to_ice_TR.png", op)
	// tileSnowToIceBRImage := PrepareImage("./assets/tiles/snow_to_ice_BR.png", op)
	// tileSnowToIceTLImage := PrepareImage("./assets/tiles/snow_to_ice_TL.png", op)
	// tileSnowToIceBLImage := PrepareImage("./assets/tiles/snow_to_ice_BL.png", op)

	// tileSnowImage := PrepareImage("./assets/tiles/snow.png", op)
	// tileSnowSpeckledImage := PrepareImage("./assets/tiles/snow_speckled.png", op)

	// // Use slice of tile coordinates to create map of tile coordinates to tiles.
	// // Start with all ice, then add other things to spice it up.
	// tiles := map[TileCoordinate]*Tile{}
	// snowBorderLeft := float64(SnowBorderSize * TileSize)
	// snowBorderRight := float64(screenWidth - TileSize*(SnowBorderSize+1))
	// snowBorderTop := float64(SnowBorderSize * TileSize)
	// snowBorderBottom := float64(screenHeight - TileSize*(SnowBorderSize+1))
	// for _, coord := range tileCoordinates {
	// 	screenCoordX := coord.ToScreenCoordinate().x
	// 	screenCoordY := coord.ToScreenCoordinate().y
	// 	if screenCoordX < snowBorderLeft || screenCoordX > snowBorderRight || screenCoordY < snowBorderTop || screenCoordY > snowBorderBottom {
	// 		if rand.Float64() <= 0.2 {
	// 			tiles[coord] = &Tile{image: tileSnowSpeckledImage, collidable: false, terrainType: TerrainTypeSnow}
	// 		} else {
	// 			tiles[coord] = &Tile{image: tileSnowImage, collidable: false, terrainType: TerrainTypeSnow}
	// 		}
	// 	} else if screenCoordX == snowBorderLeft && screenCoordY == snowBorderTop {
	// 		tiles[coord] = &Tile{image: tileSnowToIceTLImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordX == snowBorderRight && screenCoordY == snowBorderTop {
	// 		tiles[coord] = &Tile{image: tileSnowToIceTRImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordX == snowBorderLeft && screenCoordY == snowBorderBottom {
	// 		tiles[coord] = &Tile{image: tileSnowToIceBLImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordX == snowBorderRight && screenCoordY == snowBorderBottom {
	// 		tiles[coord] = &Tile{image: tileSnowToIceBRImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordX == snowBorderLeft {
	// 		tiles[coord] = &Tile{image: tileSnowToIceLImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordX == snowBorderRight {
	// 		tiles[coord] = &Tile{image: tileSnowToIceRImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordY == snowBorderTop {
	// 		tiles[coord] = &Tile{image: tileSnowToIceTImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if screenCoordY == snowBorderBottom {
	// 		tiles[coord] = &Tile{image: tileSnowToIceBImage, collidable: false, terrainType: TerrainTypeSnow}
	// 	} else if rand.Float64() <= 0.04 {
	// 		tiles[coord] = &Tile{image: tileIceStreaksImage, collidable: false, terrainType: TerrainTypeIce}
	// 	} else {
	// 		tiles[coord] = &Tile{image: tileIceImage, collidable: false, terrainType: TerrainTypeIce}
	// 	}
	// }

	// tileGroupIceHole := map[int]Tile{
	// 	0: {
	// 		image:      tileIceHoleTLImage,
	// 		collidable: true,
	// 	},
	// 	1: {
	// 		image:      tileIceHoleTRImage,
	// 		collidable: true,
	// 	},
	// 	2: {
	// 		image:      tileIceHoleBLImage,
	// 		collidable: true,
	// 	},
	// 	3: {
	// 		image:      tileIceHoleBRImage,
	// 		collidable: true,
	// 	},
	// }

	// tileGroupIceHoleWide := map[int]Tile{
	// 	0: {
	// 		image:      tileIceHoleTLImage,
	// 		collidable: true,
	// 	},
	// 	1: {
	// 		image:      tileIceHoleTImage,
	// 		collidable: true,
	// 	},
	// 	2: {
	// 		image:      tileIceHoleTRImage,
	// 		collidable: true,
	// 	},
	// 	3: {
	// 		image:      tileIceHoleBLImage,
	// 		collidable: true,
	// 	},
	// 	4: {
	// 		image:      tileIceHoleBImage,
	// 		collidable: true,
	// 	},
	// 	5: {
	// 		image:      tileIceHoleBRImage,
	// 		collidable: true,
	// 	},
	// }

	// for i := 0; i < 5; i++ {
	// 	tiles = AddTileGroup(tiles, RandomUnoccupiedIceTileCoordinate(tiles, 2, 4), 2, tileGroupIceHole)
	// }

	// tiles = AddTileGroup(tiles, RandomUnoccupiedIceTileCoordinate(tiles, 3, 6), 3, tileGroupIceHoleWide)

	return ReadMap()
}

func AddTileGroup(tiles map[TileCoordinate]*Tile, startPos TileCoordinate, width int, tileGroup map[int]Tile) map[TileCoordinate]*Tile {
	// startPos: top left tile coordinate of tile group
	// width: width of tile group in tiles
	// images: slice of images for populating tile group, ordered from left to right, top to bottom
	// collidable: true if all tiles in group are collidable, false if all tiles are not collidable
	maxIndex := 0
	for index := range tileGroup {
		if index > maxIndex {
			maxIndex = index
		}
	}

	for index := 0; index <= maxIndex; index++ {
		dx := index % width
		dy := (index - dx) / width
		if tile, ok := tileGroup[index]; ok {
			tiles[TileCoordinate{x: startPos.x + dx, y: startPos.y + dy}] = &Tile{image: tile.image, collidable: tile.collidable}
		}
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
			if tile, ok := tiles[currentCoord]; ok && tile.collidable {
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

func RandomUnoccupiedSnowTileCoordinate(tiles map[TileCoordinate]*Tile, width int, count int) TileCoordinate {
	tileCoord := RandomSnowTileCoordinate()

	for {
		occupied := false
		for i := 0; i < count; i++ {
			dx := i % width
			dy := (i - dx) / width
			currentCoord := TileCoordinate{x: tileCoord.x + dx, y: tileCoord.y + dy}
			if tile, ok := tiles[currentCoord]; ok && tile.collidable {
				occupied = true
			}
		}

		if !occupied {
			break
		}

		tileCoord = RandomSnowTileCoordinate()
	}

	return tileCoord
}

func RandomUnoccupiedIceTileCoordinate(tiles map[TileCoordinate]*Tile, width int, count int) TileCoordinate {
	tileCoord := RandomIceTileCoordinate()

	for {
		occupied := false
		for i := 0; i < count; i++ {
			dx := i % width
			dy := (i - dx) / width
			currentCoord := TileCoordinate{x: tileCoord.x + dx, y: tileCoord.y + dy}
			if tile, ok := tiles[currentCoord]; ok && tile.collidable {
				occupied = true
			}
		}

		if !occupied {
			break
		}

		tileCoord = RandomIceTileCoordinate()
	}

	return tileCoord
}

func RandomSnowTileCoordinate() TileCoordinate {
	var x, y int
	tileCountX := screenWidth / TileSize
	tileCountY := screenHeight / TileSize

	if rand.Intn(2) == 0 {
		x = rand.Intn(SnowBorderSize)
	} else {
		x = tileCountX - SnowBorderSize + rand.Intn(SnowBorderSize)
	}

	if rand.Intn(2) == 0 {
		y = rand.Intn(SnowBorderSize)
	} else {
		y = tileCountY - SnowBorderSize + rand.Intn(SnowBorderSize)
	}

	return TileCoordinate{x, y}

}

func RandomIceTileCoordinate() TileCoordinate {
	tileCountX := screenWidth / TileSize
	tileCountY := screenHeight / TileSize

	return TileCoordinate{x: SnowBorderSize + rand.Intn(tileCountX-2*SnowBorderSize), y: SnowBorderSize + rand.Intn(tileCountY-2*SnowBorderSize)}
}

func (t Tile) ToBoundingBox(c TileCoordinate) BoundingBox {
	return BoundingBox{pos: c.ToScreenCoordinate(), width: TileSize, height: TileSize}
}

func (c *ScreenCoordinate) TileCoordinateX() int {
	return int(math.Floor(c.x / TileSize))
}

func (c *ScreenCoordinate) TileCoordinateY() int {
	return int(math.Floor(c.y / TileSize))
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
