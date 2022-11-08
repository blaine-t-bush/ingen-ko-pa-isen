package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	inited  bool
	op      ebiten.DrawImageOptions
	farmer  *Actor
	cows    []*Actor
	objects []*Object
	tiles   map[TileCoordinate]*Tile
}

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	titleImage      *ebiten.Image
	farmerImage     *ebiten.Image
	cowImage        *ebiten.Image
	treeTrunkImage  *ebiten.Image
	treeCanopyImage *ebiten.Image
)

func init() {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	titleImage = PrepareImage("./assets/menu/title.png", op)
	farmerImage = PrepareImage("./assets/sprites/farmer.png", op)
	cowImage = PrepareImage("./assets/sprites/cow.png", op)
	treeTrunkImage = PrepareImage("./assets/sprites/tree_trunk.png", op)
	treeCanopyImage = PrepareImage("./assets/sprites/tree_canopy.png", op)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand.Seed(time.Now().UnixNano())

	g.tiles = GenerateTiles()

	for i := 0; i < 3; i++ {
		g.objects = append(g.objects, g.CreateRandomTree(treeTrunkImage, treeCanopyImage)...)
	}

	for i := 0; i < 5; i++ {
		g.cows = append(g.cows, g.CreateRandomCow(*cowImage))
	}

	g.farmer = g.CreateFarmer(*farmerImage)
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

	// Update cow states.
	g.UpdateCows()

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

	// Objects
	for index := range g.objects {
		s := g.objects[index]
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(s.image, &g.op)
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
