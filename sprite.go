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
	x           float64
	y           float64
}

func (s *Sprite) Update() {
	// Boundaries at left and right borders.
	if s.x < 0 {
		s.x = -s.x
	} else if mx := screenWidth - s.imageWidth; mx <= s.x {
		s.x = 2*mx - s.x
	}

	// Boundaries at top and bottom borders.
	if s.y < 0 {
		s.y = -s.y
	} else if my := screenHeight - s.imageHeight; my <= s.y {
		s.y = 2*my - s.y
	}
}

func (s *Sprite) Move(velocity Vector) {
	s.x += velocity.X()
	s.y += velocity.Y()
}
