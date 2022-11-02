package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	inited bool
	keys   []ebiten.Key
	farmer *Farmer
	op     ebiten.DrawImageOptions
}

const (
	screenWidth  = 640
	screenHeight = 480
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

	w, h := ebitenImage.Size()
	x := screenWidth / 2
	y := screenHeight / 2
	g.farmer = &Farmer{
		sprite: &Sprite{
			imageWidth:  w,
			imageHeight: h,
			x:           x,
			y:           y,
			vx:          0,
			vy:          0,
		},
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	// Listen for keyboard inputs.
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	for _, validKey := range ValidInputKeys {
		if KeysIncludes(g.keys, validKey) {
			g.HandleKeyPress(validKey)
		}
	}

	// Update player state.
	g.farmer.sprite.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	debugMsg := fmt.Sprintf("x: %d, y: %d, vx: %d, vy: %d", g.farmer.sprite.x, g.farmer.sprite.y, g.farmer.sprite.vx, g.farmer.sprite.vy)
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(float64(g.farmer.sprite.x), float64(g.farmer.sprite.y))
	screen.DrawImage(ebitenImage, &g.op)
	ebitenutil.DebugPrint(screen, debugMsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
