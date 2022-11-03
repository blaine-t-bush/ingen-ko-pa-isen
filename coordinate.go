package main

import "math"

type Coordinate struct {
	x float64
	y float64
}

func (p *Coordinate) DistanceFrom(from Coordinate) float64 {
	dX := from.x - p.x
	dY := from.y - p.y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Coordinate) WithinRadius(center Coordinate, r float64) bool {
	return p.DistanceFrom(center) <= r
}
