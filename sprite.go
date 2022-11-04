package main

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
		if object.collidable {
			collidesTop, _ := s.CollidesWithTopOf(*object.sprite)
			collidesBottom, _ := s.CollidesWithBottomOf(*object.sprite)
			collidesLeft, _ := s.CollidesWithLeftOf(*object.sprite)
			collidesRight, _ := s.CollidesWithRightOf(*object.sprite)

			if collidesTop || collidesBottom || collidesLeft || collidesRight {
				collides = true
			}
		}
	}

	return collides
}
