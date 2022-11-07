package main

import "math"

type Vector struct {
	len float64
	dir float64 // angle in radians measured counter-clockwise from +X axis. recall that the +Y direction is down.
}

func (v *Vector) X() float64 {
	return v.len * math.Cos(v.dir)
}

func (v *Vector) Y() float64 {
	return v.len * math.Sin(v.dir)
}

func (v *Vector) SetXY(p Coordinate) {
	v.len = math.Sqrt(p.x*p.x + p.y*p.y)
	if p.x == 0 && p.y == 0 {
		v.dir = 0
	} else if p.x == 0 && p.y > 0 {
		v.dir = math.Pi / 2
	} else if p.x == 0 && p.y < 0 {
		v.dir = 3 * math.Pi / 2
	} else if p.x > 0 && p.y == 0 {
		v.dir = 0
	} else if p.x < 0 && p.y == 0 {
		v.dir = math.Pi
	} else if p.x < 0 && p.y > 0 {
		v.dir = math.Pi - math.Atan(-p.y/p.x)
	} else if p.x < 0 && p.y < 0 {
		v.dir = 3*math.Pi/2 - math.Atan(p.y/p.x)
	} else if p.x > 0 && p.y > 0 {
		v.dir = 2*math.Pi - math.Atan(-p.y/p.x)
	} else {
		v.dir = math.Atan(p.y / p.x)
	}
}

func (v *Vector) Normalize() {
	v.len = 1
}

func (v *Vector) Scale(s float64) {
	v.len = s * v.len
}

func (v *Vector) Rotate(t float64) {
	v.dir = t + v.dir
}

func (v *Vector) RemoveX() {
	v.len = v.Y()
	if math.Sin(v.dir) >= 0 {
		v.dir = math.Pi / 2
	} else {
		v.dir = 3 * math.Pi / 2
	}
}

func (v *Vector) RemoveY() {
	v.len = v.X()
	if math.Cos(v.dir) >= 0 {
		v.dir = 0
	} else {
		v.dir = math.Pi
	}
}

func VectorFromXY(p Coordinate) Vector {
	v := Vector{}
	v.SetXY(p)
	return v
}

func VectorFromPoints(from Coordinate, to Coordinate) Vector {
	x := to.x - from.x
	y := to.y - from.y
	return VectorFromXY(Coordinate{x, y})
}
