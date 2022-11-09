package main

type BoundingBox struct {
	pos    ScreenCoordinate
	width  float64
	height float64
}

func (b BoundingBox) Top() float64 {
	return b.pos.y
}

func (b BoundingBox) Bottom() float64 {
	return b.pos.y + b.height
}

func (b BoundingBox) Left() float64 {
	return b.pos.x
}

func (b BoundingBox) Right() float64 {
	return b.pos.x + b.width
}

func (b BoundingBox) CenterX() float64 {
	return b.pos.x + b.width/2
}

func (b BoundingBox) CenterY() float64 {
	return b.pos.y + b.height/2
}

func (b BoundingBox) TopLeft() ScreenCoordinate {
	return b.pos
}

func (b BoundingBox) TopRight() ScreenCoordinate {
	return ScreenCoordinate{b.Right(), b.Top()}
}

func (b BoundingBox) BottomLeft() ScreenCoordinate {
	return ScreenCoordinate{b.Left(), b.Bottom()}
}

func (b BoundingBox) BottomRight() ScreenCoordinate {
	return ScreenCoordinate{b.Right(), b.Bottom()}
}

func (b BoundingBox) Center() ScreenCoordinate {
	return ScreenCoordinate{x: b.CenterX(), y: b.CenterY()}
}

func (b BoundingBox) CollidesWithTopBorder() (bool, float64) {
	if b.Top() <= 0 {
		return true, b.Top()
	}

	return false, 0
}

func (b BoundingBox) CollidesWithBottomBorder() (bool, float64) {
	if b.Bottom() >= screenHeight {
		return true, screenHeight - b.Bottom()
	}

	return false, 0
}

func (b BoundingBox) CollidesWithLeftBorder() (bool, float64) {
	if b.Left() <= 0 {
		return true, b.Left()
	}

	return false, 0
}

func (b BoundingBox) CollidesWithRightBorder() (bool, float64) {
	if b.Right() >= screenWidth {
		return true, screenWidth - b.Right()
	}

	return false, 0
}

func (b BoundingBox) CollidesWithBorders() bool {
	collidesTop, _ := b.CollidesWithTopBorder()
	collidesBottom, _ := b.CollidesWithBottomBorder()
	collidesLeft, _ := b.CollidesWithLeftBorder()
	collidesRight, _ := b.CollidesWithRightBorder()
	return collidesTop || collidesBottom || collidesLeft || collidesRight
}

func (g *Game) CheckCollision(b BoundingBox) bool {
	collides := false

	// checks screen boundaries
	if b.CollidesWithBorders() {
		collides = true
	}

	// loop over objects
	for _, object := range g.objects {
		// if object is collidable and rectangle is within boundaries
		if object.collidable && b.CollidesWithBox(object.BoundingBox()) {
			collides = true
		}
	}

	// loop over tiles
	for coord, tile := range g.tiles {
		// if object is collidable and rectangle is within boundaries
		if tile.collidable && b.CollidesWithBox(tile.ToBoundingBox(coord)) {
			collides = true
		}
	}

	// TODO loop over actors?

	return collides
}

func (b1 BoundingBox) CollidesWithTopOf(b2 BoundingBox) (bool, float64) {
	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if b1.Bottom() >= b2.Top() && b1.Bottom() <= b2.Bottom() && ((b1.Right() >= b2.Left() && b1.Right() <= b2.Right()) || (b1.Left() >= b2.Left() && b1.Left() <= b2.Right()) || (b1.Left() <= b2.Left() && b1.Right() >= b2.Right())) {
		collides = true
		overlap = b1.Bottom() - b2.Top()
	}

	return collides, overlap
}

func (b1 BoundingBox) CollidesWithBottomOf(b2 BoundingBox) (bool, float64) {
	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if b1.Top() >= b2.Top() && b1.Top() <= b2.Bottom() && ((b1.Right() >= b2.Left() && b1.Right() <= b2.Right()) || (b1.Left() >= b2.Left() && b1.Left() <= b2.Right()) || (b1.Left() <= b2.Left() && b1.Right() >= b2.Right())) {
		collides = true
		overlap = b2.Bottom() - b1.Top()
	}

	return collides, overlap
}

func (b1 BoundingBox) CollidesWithLeftOf(b2 BoundingBox) (bool, float64) {
	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if b1.Right() >= b2.Left() && b1.Right() <= b2.Right() && ((b1.Bottom() >= b2.Top() && b1.Bottom() <= b2.Bottom()) || (b1.Top() >= b2.Top() && b1.Top() <= b2.Bottom()) || (b1.Top() <= b2.Top() && b1.Bottom() >= b2.Bottom())) {
		collides = true
		overlap = b1.Right() - b2.Left()
	}

	return collides, overlap
}

func (b1 BoundingBox) CollidesWithRightOf(b2 BoundingBox) (bool, float64) {
	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if b1.Left() >= b2.Left() && b1.Left() <= b2.Right() && ((b1.Bottom() >= b2.Top() && b1.Bottom() <= b2.Bottom()) || (b1.Top() >= b2.Top() && b1.Top() <= b2.Bottom()) || (b1.Top() <= b2.Top() && b1.Bottom() >= b2.Bottom())) {
		collides = true
		overlap = b2.Right() - b1.Left()
	}

	return collides, overlap
}

func (b1 BoundingBox) ContainedWithin(b2 BoundingBox) bool {
	collides := false

	if b1.Left() >= b2.Left() && b1.Right() <= b2.Right() && b1.Top() >= b2.Top() && b1.Bottom() <= b2.Bottom() {
		collides = true
	}

	return collides
}

func (b1 BoundingBox) CollidesWithBox(b2 BoundingBox) bool {
	collides := false

	collidesTop, _ := b1.CollidesWithTopOf(b2)
	collidesBottom, _ := b1.CollidesWithBottomOf(b2)
	collidesLeft, _ := b1.CollidesWithLeftOf(b2)
	collidesRight, _ := b1.CollidesWithRightOf(b2)
	collidesWithin := b1.ContainedWithin(b2)

	if collidesTop || collidesBottom || collidesLeft || collidesRight || collidesWithin {
		collides = true
	}

	return collides
}
