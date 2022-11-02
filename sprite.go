package main

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          int
	vy          int
}

func (s *Sprite) SpeedSquared() int {
	return s.vx*s.vx + s.vy*s.vy
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
