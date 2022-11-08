package main

import "math"

type ScreenCoordinate struct {
	x float64
	y float64
}

func (c *ScreenCoordinate) Translate(offset Vector) {
	c.x += offset.X()
	c.y += offset.Y()
}

func (c ScreenCoordinate) DistanceFrom(from ScreenCoordinate) float64 {
	dX := from.x - c.x
	dY := from.y - c.y
	return math.Sqrt(dX*dX + dY*dY)
}

func (c ScreenCoordinate) WithinRadius(center ScreenCoordinate, r float64) bool {
	return c.DistanceFrom(center) <= r
}
