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
	farmer  *Farmer
	cows    []*Cow
	objects []*Object
}

const (
	screenWidth      = 640
	screenHeight     = 480
	defaultCowCount  = 10
	defaultRockCount = 1
)

var (
	titleImage  *ebiten.Image
	farmerImage *ebiten.Image
	cowImage    *ebiten.Image
	rockImage   *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	origTitleImage := NewImageFromFilePath("./assets/menu/title.png")
	origFarmerImage := NewImageFromFilePath("./assets/sprites/farmer.png")
	origCowImage := NewImageFromFilePath("./assets/sprites/cow.png")
	origRockImage := NewImageFromFilePath("./assets/sprites/rock.png")

	var w, h int

	w, h = origTitleImage.Size()
	titleImage = ebiten.NewImage(w, h)

	w, h = origFarmerImage.Size()
	farmerImage = ebiten.NewImage(w, h)

	w, h = origCowImage.Size()
	cowImage = ebiten.NewImage(w, h)

	w, h = origRockImage.Size()
	rockImage = ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)
	titleImage.DrawImage(origTitleImage, op)
	farmerImage.DrawImage(origFarmerImage, op)
	cowImage.DrawImage(origCowImage, op)
	rockImage.DrawImage(origRockImage, op)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand.Seed(time.Now().UnixNano())

	g.farmer = CreateFarmer(*farmerImage)

	for i := 0; i < defaultCowCount; i++ {
		g.cows = append(g.cows, CreateRandomCow(*cowImage))
	}

	for i := 0; i < defaultRockCount; i++ {
		g.objects = append(g.objects, CreateRandomRock(*rockImage))
	}
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

	// Update player state.
	g.CheckFarmerCollision()
	g.farmer.Update()

	// Update cow states.
	for _, cow := range g.cows {
		cow.Update(g.farmer)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	screen.Fill(color.NRGBA{0x0, 0xff, 0xff, 0xff})

	// Title
	g.op.GeoM.Reset()
	screen.DrawImage(titleImage, &g.op)

	// Farmer
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.farmer.sprite.pos.x, g.farmer.sprite.pos.y)
	screen.DrawImage(farmerImage, &g.op)

	// Cows
	for index := range g.cows {
		s := g.cows[index].sprite
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(cowImage, &g.op)
	}

	// Rocks
	for index := range g.objects {
		s := g.objects[index].sprite
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(rockImage, &g.op)
	}
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
