package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	inited     bool
	op         ebiten.DrawImageOptions
	farmer     *Actor
	cows       []*Actor
	objects    []*Object
	footprints []*Object
	tiles      map[TileCoordinate]*Tile
}

const (
	screenWidth  = 1280
	screenHeight = 960
)

var (
	titleImage         *ebiten.Image
	farmerImage        *ebiten.Image
	cowImage           *ebiten.Image
	cowPieImage        *ebiten.Image
	treeTrunkImage     *ebiten.Image
	treeCanopyImage    *ebiten.Image
	footprintIceImage  *ebiten.Image
	footprintSnowImage *ebiten.Image
)

func init() {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	titleImage = PrepareImage("./assets/menu/title.png", op)
	farmerImage = PrepareImage("./assets/sprites/farmer.png", op)
	cowImage = PrepareImage("./assets/sprites/cow.png", op)
	cowPieImage = PrepareImage("./assets/sprites/cow_pie.png", op)
	treeTrunkImage = PrepareImage("./assets/sprites/tree_trunk.png", op)
	treeCanopyImage = PrepareImage("./assets/sprites/tree_canopy.png", op)
	footprintIceImage = PrepareImage("./assets/sprites/footprint_ice.png", op)
	footprintSnowImage = PrepareImage("./assets/sprites/footprint_snow.png", op)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	fmt.Println("Initializing game...")
	fmt.Println(" - Seeding random number generator...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println(" - Generating tilemap...")
	g.tiles = GenerateTiles()

	fmt.Println(" - Creating trees...")
	for i := 0; i < 3; i++ {
		g.objects = append(g.objects, g.CreateRandomTree()...)
	}

	fmt.Println(" - Creating cow pies...")
	for i := 0; i < 3; i++ {
		g.objects = append(g.objects, g.CreateRandomCowPie())
	}

	fmt.Println(" - Creating cows...")
	for i := 0; i < 5; i++ {
		g.cows = append(g.cows, g.CreateRandomCow())
	}

	fmt.Println(" - Creating farmer...")
	g.farmer = g.CreateFarmer(*farmerImage)

	fmt.Println("Done initializing. Running game...")
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	// Listen for mouse inputs.
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.HandleMouseLeftClick()
	}

	// Listen for keyboard inputs.
	keys := []ebiten.Key{}
	keys = inpututil.AppendPressedKeys(keys[:0])
	g.HandleKeyPresses(keys)

	// Update actor positions based on their velocities.
	g.UpdateCows()
	g.UpdateFarmer()

	// Update footprints.
	g.UpdateFootprints()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xb5, 0xdd, 0xff, 0xff})

	// Tiles
	for coord, tile := range g.tiles {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(coord.ScreenCoordinateX(), coord.ScreenCoordinateY())
		screen.DrawImage(tile.image, &g.op)
	}

	// Footprints
	for index := range g.footprints {
		s := g.footprints[index]
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(s.image, &g.op)
	}

	// Objects: below actors
	for index := range g.objects {
		s := g.objects[index]
		if !s.aboveActors {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(s.pos.x, s.pos.y)
			screen.DrawImage(s.image, &g.op)
		}
	}

	// Cows
	for index := range g.cows {
		s := g.cows[index]
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(s.image, &g.op)
	}

	// Farmer
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.farmer.pos.x, g.farmer.pos.y)
	screen.DrawImage(g.farmer.image, &g.op)

	// Objects: above actors
	for index := range g.objects {
		s := g.objects[index]
		if s.aboveActors {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(s.pos.x, s.pos.y)
			screen.DrawImage(s.image, &g.op)
		}
	}

	// Title
	w, _ := titleImage.Size()
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(float64(screenWidth)/2-float64(w)/2, 0)
	screen.DrawImage(titleImage, &g.op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ingen Ko På Isen!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
