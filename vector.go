package main

import "math"

type Coordinate struct {
	x float64
	y float64
}

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
	if p.x < 0 {
		v.dir = math.Pi - math.Atan(-p.y/p.x)
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

func (p *Coordinate) DistanceFrom(from Coordinate) float64 {
	dX := from.x - p.x
	dY := from.y - p.y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Coordinate) WithinRadius(center Coordinate, r float64) bool {
	return p.DistanceFrom(center) <= r
}
