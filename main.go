package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 768
	CellSize     = 1
	PlayerSize   = 16
	ActorSize    = 12
	EntitySize   = 10
)

var (
	playerImage = ebiten.NewImage(PlayerSize, PlayerSize)
	actorImage  = ebiten.NewImage(ActorSize, ActorSize)
	entityImage = ebiten.NewImage(EntitySize, EntitySize)
)

type Game struct {
	inited   bool
	op       *ebiten.DrawImageOptions
	space    *resolv.Space
	player   *Actor
	actors   []*Actor
	entities []*Entity
	tiles    map[TileCoord]*Tile
}

func (g *Game) Init() {
	defer func() {
		g.inited = true
	}()

	// Prepare images
	g.op = &ebiten.DrawImageOptions{}
	g.op.ColorM.Scale(1, 1, 1, 1)
	playerImage.Fill(color.Black)
	actorImage.Fill(color.White)
	entityImage.Fill(color.White)

	// Define space
	g.space = resolv.NewSpace(ScreenWidth, ScreenHeight, CellSize, CellSize)

	// Create player struct
	g.CreatePlayer(playerImage)

	// Create additional actor
	g.CreateActor(actorImage, 400, 400, float64(actorImage.Bounds().Dx()), float64(actorImage.Bounds().Dy()))

	// Create entity structs
	g.CreateTree()
	g.CreateEntity(entityImage, 20, 20, float64(entityImage.Bounds().Dx()), float64(entityImage.Bounds().Dy()), true)
	g.CreateEntity(entityImage, 80, 20, float64(entityImage.Bounds().Dx()), float64(entityImage.Bounds().Dy()), true)
	g.CreateEntity(entityImage, 70, 40, float64(entityImage.Bounds().Dx()), float64(entityImage.Bounds().Dy()), true)
	g.CreateEntity(entityImage, 120, 70, float64(entityImage.Bounds().Dx()), float64(entityImage.Bounds().Dy()), true)

	// Generate tiles
	g.GenerateTiles()
}

func (g *Game) Update() error {
	if !g.inited {
		g.Init()
	}

	g.HandleKeyPresses()

	g.MoveActors()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x60, 0x60, 0x60, 0xff})

	// Render tiles
	for tileCoord, tile := range g.tiles {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(tileCoord.X)*TileSize, float64(tileCoord.Y)*TileSize)
		screen.DrawImage(tile.Image, g.op)
	}

	// Render player
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.player.Object.X, g.player.Object.Y)
	screen.DrawImage(g.player.Image, g.op)

	// Render other actors
	for _, actor := range g.actors {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(actor.Object.X, actor.Object.Y)
		screen.DrawImage(actor.Image, g.op)
	}

	// Render entities
	for _, entity := range g.entities {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(entity.Object.X, entity.Object.Y)
		screen.DrawImage(entity.Image, g.op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("My Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
