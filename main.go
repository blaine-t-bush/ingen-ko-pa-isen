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
	g.farmer = &Farmer{sprite: &Sprite{
		imageWidth:  float64(w),
		imageHeight: float64(h),
		x:           float64(x),
		y:           float64(y),
	},
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	// Listen for keyboard inputs.
	keys := []ebiten.Key{}
	keys = inpututil.AppendPressedKeys(keys[:0])
	g.HandleKeyPresses(keys)

	// Update player state.
	g.farmer.sprite.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	debugMsg := fmt.Sprintf("x: %.2f, y: %.2f", g.farmer.sprite.x, g.farmer.sprite.y)
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.farmer.sprite.x, g.farmer.sprite.y)
	screen.DrawImage(ebitenImage, &g.op)
	ebitenutil.DebugPrint(screen, debugMsg)
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
