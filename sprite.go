package main

import "github.com/hajimehoshi/ebiten/v2"

type Direction int

const (
	MoveUnitDiagonal           = 1
	MoveUnitStraight           = 1.41421
	DirectionLeft    Direction = iota
	DirectionRight
	DirectionUp
	DirectionDown
	DirectionLeftUp
	DirectionLeftDown
	DirectionRightUp
	DirectionRightDown
)

type Sprite struct {
	image       *ebiten.Image
	imageWidth  float64
	imageHeight float64
	pos         *Coordinate
}

func (s *Sprite) Center() Coordinate {
	return Coordinate{x: s.pos.x + s.imageWidth/2, y: s.pos.y + s.imageHeight/2}
}

func (s *Sprite) Move(velocity Vector) {
	s.pos.x += velocity.X()
	s.pos.y += velocity.Y()
}

func (s *Sprite) CollidesWithTopOf(t Sprite) (bool, float64) {
	x1 := s.pos.x
	x2 := s.pos.x + s.imageWidth
	y2 := s.pos.y + s.imageHeight
	tx1 := t.pos.x
	tx2 := t.pos.x + t.imageWidth
	ty1 := t.pos.y
	ty2 := t.pos.y + t.imageHeight

	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if y2 >= ty1 && y2 <= ty2 && ((x2 >= tx1 && x2 <= tx2) || (x1 >= tx1 && x1 <= tx2) || (x1 <= tx1 && x2 >= tx2)) {
		collides = true
		overlap = y2 - ty1
	}

	return collides, overlap
}

func (s *Sprite) CollidesWithBottomOf(t Sprite) (bool, float64) {
	x1 := s.pos.x
	x2 := s.pos.x + s.imageWidth
	y1 := s.pos.y
	tx1 := t.pos.x
	tx2 := t.pos.x + t.imageWidth
	ty1 := t.pos.y
	ty2 := t.pos.y + t.imageHeight

	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if y1 >= ty1 && y1 <= ty2 && ((x2 >= tx1 && x2 <= tx2) || (x1 >= tx1 && x1 <= tx2) || (x1 <= tx1 && x2 >= tx2)) {
		collides = true
		overlap = ty2 - y1
	}

	return collides, overlap
}

func (s *Sprite) CollidesWithLeftOf(t Sprite) (bool, float64) {
	x2 := s.pos.x + s.imageWidth
	y1 := s.pos.y
	y2 := s.pos.y + s.imageHeight
	tx1 := t.pos.x
	tx2 := t.pos.x + t.imageWidth
	ty1 := t.pos.y
	ty2 := t.pos.y + t.imageHeight

	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if x2 >= tx1 && x2 <= tx2 && ((y2 >= ty1 && y2 <= ty2) || (y1 >= ty1 && y1 <= ty2) || (y1 <= ty1 && y2 >= ty2)) {
		collides = true
		overlap = x2 - tx1
	}

	return collides, overlap
}

func (s *Sprite) CollidesWithRightOf(t Sprite) (bool, float64) {
	x1 := s.pos.x
	y1 := s.pos.y
	y2 := s.pos.y + s.imageHeight
	tx1 := t.pos.x
	tx2 := t.pos.x + t.imageWidth
	ty1 := t.pos.y
	ty2 := t.pos.y + t.imageHeight

	collides := false
	overlap := 0.0

	// bottom edge of shape passes through object
	if x1 >= tx1 && x1 <= tx2 && ((y2 >= ty1 && y2 <= ty2) || (y1 >= ty1 && y1 <= ty2) || (y1 <= ty1 && y2 >= ty2)) {
		collides = true
		overlap = tx2 - x1
	}

	return collides, overlap
}

func (s *Sprite) ContainedWithin(t Sprite) bool {
	x1 := s.pos.x
	x2 := s.pos.x + s.imageWidth
	y1 := s.pos.y
	y2 := s.pos.y + s.imageHeight
	tx1 := t.pos.x
	tx2 := t.pos.x + t.imageWidth
	ty1 := t.pos.y
	ty2 := t.pos.y + t.imageHeight

	collides := false

	if x1 >= tx1 && x2 <= tx2 && y1 >= ty1 && y2 <= ty2 {
		collides = true
	}

	return collides
}

func (s *Sprite) CollidesWithSprite(t Sprite) bool {
	collides := false

	collidesTop, _ := s.CollidesWithTopOf(t)
	collidesBottom, _ := s.CollidesWithBottomOf(t)
	collidesLeft, _ := s.CollidesWithLeftOf(t)
	collidesRight, _ := s.CollidesWithRightOf(t)
	collidesWithin := s.ContainedWithin(t)

	if collidesTop || collidesBottom || collidesLeft || collidesRight || collidesWithin {
		collides = true
	}

	return collides
}

func (s *Sprite) CollidesWithTopBorder() (bool, float64) {
	if s.pos.y <= 0 {
		return true, s.pos.y
	}

	return false, 0
}

func (s *Sprite) CollidesWithBottomBorder() (bool, float64) {
	if s.pos.y+s.imageHeight >= screenHeight {
		return true, screenHeight - s.pos.y - s.imageHeight
	}

	return false, 0
}

func (s *Sprite) CollidesWithLeftBorder() (bool, float64) {
	if s.pos.x <= 0 {
		return true, s.pos.x
	}

	return false, 0
}

func (s *Sprite) CollidesWithRightBorder() (bool, float64) {
	if s.pos.x+s.imageWidth >= screenWidth {
		return true, screenWidth - s.pos.x - s.imageWidth
	}

	return false, 0
}

func (s *Sprite) CollidesWithBorders() bool {
	collidesTop, _ := s.CollidesWithTopBorder()
	collidesBottom, _ := s.CollidesWithBottomBorder()
	collidesLeft, _ := s.CollidesWithLeftBorder()
	collidesRight, _ := s.CollidesWithRightBorder()
	return collidesTop || collidesBottom || collidesLeft || collidesRight
}

func (g *Game) CheckCollision(s *Sprite) bool {
	collides := false

	// checks screen boundaries
	if s.CollidesWithBorders() {
		collides = true
	}

	// loop over objects
	for _, object := range g.objects {
		// if object is collidable and rectangle is within boundaries
		if object.collidable && s.CollidesWithSprite(*object.sprite) {
			collides = true
		}
	}

	return collides
}

// type BoundingBox struct {
// 	pos    Coordinate
// 	width  float64
// 	height float64
// }

// func (b *BoundingBox) Top() float64 {
// 	return b.pos.y
// }

// func (b *BoundingBox) Bottom() float64 {
// 	return b.pos.y + b.height
// }

// func (b *BoundingBox) Left() float64 {
// 	return b.pos.x
// }

// func (b *BoundingBox) Right() float64 {
// 	return b.pos.x + b.width
// }

// func (b *BoundingBox) TopLeft() Coordinate {
// 	return b.pos
// }

// func (b *BoundingBox) TopRight() Coordinate {
// 	return Coordinate{b.Right(), b.Top()}
// }

// func (b *BoundingBox) BottomLeft() Coordinate {
// 	return Coordinate{b.Left(), b.Bottom()}
// }

// func (b *BoundingBox) BottomRight() Coordinate {
// 	return Coordinate{b.Right(), b.Bottom()}
// }

// func (g *Game) OccupiedAreas() []Sprite {
// 	var occupiedAreas []Sprite

// 	// add objects to occupiedAreas
// 	for _, object := range g.objects {
// 		occupiedAreas = append(occupiedAreas, *object.sprite)
// 	}

// 	return occupiedAreas
// }
