package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	inited bool
	op     ebiten.DrawImageOptions
	farmer *Farmer
	cows   []*Cow
}

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultCowCount = 10
)

var (
	titleImage  *ebiten.Image
	farmerImage *ebiten.Image
	cowImage    *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	origTitleImage := NewImageFromFilePath("./assets/menu/title.png")
	origFarmerImage := NewImageFromFilePath("./assets/sprites/farmer.png")
	origCowImage := NewImageFromFilePath("./assets/sprites/cow.png")

	var w, h int

	w, h = origTitleImage.Size()
	titleImage = ebiten.NewImage(w, h)

	w, h = origFarmerImage.Size()
	farmerImage = ebiten.NewImage(w, h)

	w, h = origCowImage.Size()
	cowImage = ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)
	titleImage.DrawImage(origTitleImage, op)
	farmerImage.DrawImage(origFarmerImage, op)
	cowImage.DrawImage(origCowImage, op)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand.Seed(time.Now().UnixNano())

	w, h := farmerImage.Size()
	g.farmer = &Farmer{
		sprite: &Sprite{
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(screenWidth / 2), float64(screenHeight / 2)},
		},
	}

	for i := 0; i < defaultCowCount; i++ {
		w, h := cowImage.Size()
		g.cows = append(g.cows, &Cow{
			sprite: &Sprite{
				imageWidth:  float64(w),
				imageHeight: float64(h),
				pos:         &Coordinate{float64(rand.Intn(screenWidth)), float64(rand.Intn(screenHeight))},
			},
			velocity: &Vector{
				dir: 2 * math.Pi * rand.Float64(),
				len: 1,
			},
		})
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
