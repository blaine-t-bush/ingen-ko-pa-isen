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
}

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	titleImage      *ebiten.Image
	farmerImage     *ebiten.Image
	cowImage        *ebiten.Image
	rockImage       *ebiten.Image
	treeImage       *ebiten.Image
	iceHoleImage    *ebiten.Image
	iceStreaksImage *ebiten.Image
)

func init() {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	titleImage = prepareImage("./assets/menu/title.png", op)
	farmerImage = prepareImage("./assets/sprites/farmer.png", op)
	cowImage = prepareImage("./assets/sprites/cow.png", op)
	rockImage = prepareImage("./assets/sprites/rock.png", op)
	treeImage = prepareImage("./assets/sprites/tree.png", op)
	iceHoleImage = prepareImage("./assets/sprites/ice_hole.png", op)
	iceStreaksImage = prepareImage("./assets/sprites/ice_streaks.png", op)
}

func prepareImage(filePath string, op *ebiten.DrawImageOptions) *ebiten.Image {
	origImage := NewImageFromFilePath(filePath)
	w, h := origImage.Size()
	image := ebiten.NewImage(w, h)
	image.DrawImage(origImage, op)

	return image
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		g.objects = append(g.objects, g.CreateRandomObject(*rockImage, true))
	}

	for i := 0; i < 5; i++ {
		g.objects = append(g.objects, g.CreateRandomObject(*iceHoleImage, true))
	}

	for i := 0; i < 3; i++ {
		g.objects = append(g.objects, g.CreateRandomObject(*treeImage, true))
	}

	for i := 0; i < 10; i++ {
		g.objects = append(g.objects, g.CreateRandomObject(*iceStreaksImage, false))
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
	// Background: 89BAFF
	screen.Fill(color.NRGBA{0x89, 0xba, 0xff, 0xff})

	// Objects
	for index := range g.objects {
		s := g.objects[index].sprite
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(s.image, &g.op)
	}

	// Cows
	for index := range g.cows {
		s := g.cows[index].sprite
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(s.image, &g.op)
	}

	// Farmer
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.farmer.sprite.pos.x, g.farmer.sprite.pos.y)
	screen.DrawImage(g.farmer.sprite.image, &g.op)

	// Title
	g.op.GeoM.Reset()
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
