package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
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
	defaultCowCount = 3
)

var (
	ebitenImage *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	origEbitenImage := ebiten.NewImageFromImage(img)

	w, h := origEbitenImage.Size()
	ebitenImage = ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	ebitenImage.DrawImage(origEbitenImage, op)
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	rand.Seed(time.Now().UnixNano())

	w, h := ebitenImage.Size()
	x := screenWidth / 2
	y := screenHeight / 2
	g.farmer = &Farmer{
		sprite: &Sprite{
			imageWidth:  float64(w),
			imageHeight: float64(h),
			pos:         &Coordinate{float64(x), float64(y)},
		},
	}

	for i := 0; i < defaultCowCount; i++ {
		w, h := ebitenImage.Size()
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenHeight)
		dir := 2 * math.Pi * rand.Float64()
		g.cows = append(g.cows, &Cow{
			sprite: &Sprite{
				imageWidth:  float64(w),
				imageHeight: float64(h),
				pos:         &Coordinate{float64(x), float64(y)},
			},
			velocity: &Vector{
				dir: dir,
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
		cow.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Farmer
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.farmer.sprite.pos.x, g.farmer.sprite.pos.y)
	screen.DrawImage(ebitenImage, &g.op)

	// Cows
	for index := range g.cows {
		s := g.cows[index].sprite
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(s.pos.x, s.pos.y)
		screen.DrawImage(ebitenImage, &g.op)
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
