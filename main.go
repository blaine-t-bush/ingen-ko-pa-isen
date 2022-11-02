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

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          int
	vy          int
}

type Game struct {
	inited bool
	keys   []ebiten.Key
	player *Sprite
	op     ebiten.DrawImageOptions
}

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	ebitenImage *ebiten.Image
)

var validInputKeys []ebiten.Key = []ebiten.Key{
	ebiten.KeyArrowLeft,
	ebiten.KeyArrowRight,
	ebiten.KeyArrowUp,
	ebiten.KeyArrowDown,
}

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

func (s *Sprite) Update() {
	// Update position based on current speed.
	s.x += s.vx
	s.y += s.vy

	// Boundaries at left and right borders.
	if s.x < 0 {
		s.x = -s.x
		s.vx = 0
	} else if mx := screenWidth - s.imageWidth; mx <= s.x {
		s.x = 2*mx - s.x
		s.vx = 0
	}

	// Boundaries at top and bottom borders.
	if s.y < 0 {
		s.y = -s.y
		s.vy = 0
	} else if my := screenHeight - s.imageHeight; my <= s.y {
		s.y = 2*my - s.y
		s.vy = 0
	}
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	w, h := ebitenImage.Size()
	x := screenWidth / 2
	y := screenHeight / 2
	g.player = &Sprite{
		imageWidth:  w,
		imageHeight: h,
		x:           x,
		y:           y,
		vx:          0,
		vy:          0,
	}
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}

	// Listen for keyboard inputs.
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	for _, validKey := range validInputKeys {
		if keysIncludes(g.keys, validKey) {
			g.handleKeyPress(validKey)
		}
	}

	// Update player state.
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	debugMsg := fmt.Sprintf("x: %d, y: %d, vx: %d, vy: %d", g.player.x, g.player.y, g.player.vx, g.player.vy)
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(float64(g.player.x), float64(g.player.y))
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

func keysIncludes(keys []ebiten.Key, includes ebiten.Key) bool {
	for _, key := range keys {
		if key == includes {
			return true
		}
	}

	return false
}

func (g *Game) handleKeyPress(key ebiten.Key) error {
	switch key {
	case ebiten.KeyArrowLeft:
		g.player.vx -= 1
	case ebiten.KeyArrowRight:
		g.player.vx += 1
	case ebiten.KeyArrowDown:
		g.player.vy += 1
	case ebiten.KeyArrowUp:
		g.player.vy -= 1
	}

	return nil
}
