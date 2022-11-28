package main

import (
	"bufio"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type TileCoord struct {
	X int
	Y int
}

type Tile struct {
	ID     int
	Object *resolv.Object
	Image  *ebiten.Image
}

type TileInfo struct {
	Collidable bool
	Image      *ebiten.Image
}

const (
	TileSize = 16
	TileIce  = iota
	TileSnow
	TileWater
	TileDirt
)

var (
	tileIceImage   = ebiten.NewImage(TileSize, TileSize)
	tileSnowImage  = ebiten.NewImage(TileSize, TileSize)
	tileWaterImage = ebiten.NewImage(TileSize, TileSize)
	tileDirtImage  = ebiten.NewImage(TileSize, TileSize)

	tilesSymbol = map[string]int{
		"I": TileIce,
		"S": TileSnow,
		"W": TileWater,
	}

	tilesInfo = map[int]TileInfo{
		TileIce:   {false, tileIceImage},
		TileSnow:  {false, tileSnowImage},
		TileWater: {true, tileWaterImage},
		TileDirt:  {false, tileDirtImage},
	}
)

func (g *Game) PrepareTileImages() {
	tileIceImage.Fill(color.NRGBA{0xb5, 0xdd, 0xff, 0xff})
	tileSnowImage.Fill(color.NRGBA{0xfa, 0xfa, 0xfa, 0xff})
	tileWaterImage.Fill(color.NRGBA{0x5e, 0x7c, 0xff, 0xff})
	tileDirtImage.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
}

func (g *Game) ReadMap() map[TileCoord]*Tile {
	file, err := os.Open("./maps/default.map")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tiles := map[TileCoord]*Tile{}
	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		for colIndex, symbol := range scanner.Text() {
			// Get tile info based on character in map.
			if string(symbol) == "\n" {
				continue
			}
			tileID := tilesSymbol[string(symbol)]
			tileInfo := tilesInfo[tileID]
			tileCoord := TileCoord{colIndex, rowIndex}

			// Define tile object.
			obj := resolv.NewObject(float64(tileCoord.X)*TileSize, float64(tileCoord.Y)*TileSize, TileSize, TileSize)
			if tileInfo.Collidable {
				obj.AddTags("collidable")
			}
			g.space.Add(obj)

			// Define tile image.
			img := tileInfo.Image

			tiles[tileCoord] = &Tile{
				ID:     tileID,
				Image:  img,
				Object: obj,
			}
		}
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tiles
}

func (g *Game) GenerateTiles() {
	g.PrepareTileImages()
	g.tiles = g.ReadMap()
}
