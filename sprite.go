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

func (s *Sprite) Update() {
	// Boundaries at left and right borders.
	if s.pos.x < 0 {
		s.pos.x = -s.pos.x
	} else if mx := screenWidth - s.imageWidth; mx <= s.pos.x {
		s.pos.x = 2*mx - s.pos.x
	}

	// Boundaries at top and bottom borders.
	if s.pos.y < 0 {
		s.pos.y = -s.pos.y
	} else if my := screenHeight - s.imageHeight; my <= s.pos.y {
		s.pos.y = 2*my - s.pos.y
	}
}

func (s *Sprite) Move(velocity Vector) {
	s.pos.x += velocity.X()
	s.pos.y += velocity.Y()
}
