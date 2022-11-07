package main

import (
	"math"
	"testing"
)

func TestX(t *testing.T) {
	helpX(t, 1, 0, 1)
	helpX(t, 1, math.Pi/2, 0)
	helpX(t, 1, math.Pi, -1)
	helpX(t, 1, 3*math.Pi/2, 0)
	helpX(t, 1, 2*math.Pi, 1)
}

func helpX(t *testing.T, len float64, dir float64, expectedX float64) {
	v := Vector{len: len, dir: dir}
	if !WithinEps(v.X(), expectedX) {
		t.Errorf("Expected vector to have x=%.2f, got x=%.2f", expectedX, v.X())
	}
}

func TestY(t *testing.T) {
	helpY(t, 1, 0, 0)
	helpY(t, 1, math.Pi/2, 1)
	helpY(t, 1, math.Pi, 0)
	helpY(t, 1, 3*math.Pi/2, -1)
	helpY(t, 1, 2*math.Pi, 0)
}

func helpY(t *testing.T, len float64, dir float64, expectedY float64) {
	v := Vector{len: len, dir: dir}
	if !WithinEps(v.Y(), expectedY) {
		t.Errorf("Expected vector to have x=%.2f, got x=%.2f", expectedY, v.Y())
	}
}

func TestFromXY(t *testing.T) {
	helpFromXY(t, 0, 0, 0)
	helpFromXY(t, 1, 0, 0)
	helpFromXY(t, -1, 0, math.Pi)
	helpFromXY(t, 0, 1, math.Pi/2)
	helpFromXY(t, 0, -1, 3*math.Pi/2)
	helpFromXY(t, 1, 1, math.Pi/4)
	helpFromXY(t, -1, 1, 3*math.Pi/4)
	helpFromXY(t, -1, -1, 5*math.Pi/4)
	helpFromXY(t, 1, -1, 7*math.Pi/4)
}

func helpFromXY(t *testing.T, x float64, y float64, expectedDir float64) {
	v := VectorFromXY(Coordinate{x, y})

	if !WithinEps(v.dir, expectedDir) {
		t.Errorf("Expected vector to have dir=%.2f, got dir=%.2f", expectedDir, v.dir)
	}
}

func TestRemoveX(t *testing.T) {
	helpRemoveX(t, 1, 0, 0)
	helpRemoveX(t, 1, math.Pi, 0)
	helpRemoveX(t, 1, math.Pi/2, 1)
	helpRemoveX(t, 1, 3*math.Pi/2, 1)
	helpRemoveX(t, 1, 2*math.Pi, 0)
}

func TestRemoveY(t *testing.T) {
	helpRemoveY(t, 1, 0, 1)
	helpRemoveY(t, 1, math.Pi, 1)
	helpRemoveY(t, 1, math.Pi/2, 0)
	helpRemoveY(t, 1, 3*math.Pi/2, 0)
	helpRemoveY(t, 1, 2*math.Pi, 1)
}

func helpRemoveX(t *testing.T, len float64, dir float64, expectedY float64) {
	v := Vector{len: len, dir: dir}
	v.RemoveX()

	if !WithinEps(v.X(), 0) {
		t.Errorf("Expected vector to have x=%d, got x=%.2f", 0, v.X())
	}

	if !WithinEps(v.Y(), expectedY) {
		t.Errorf("Expected vector to have y=%.2f, got y=%.2f", expectedY, v.Y())
	}
}

func helpRemoveY(t *testing.T, len float64, dir float64, expectedX float64) {
	v := Vector{len: len, dir: dir}
	v.RemoveY()

	if !WithinEps(v.Y(), 0) {
		t.Errorf("Expected vector to have y=%d, got y=%.2f", 0, v.Y())
	}

	if !WithinEps(v.X(), expectedX) {
		t.Errorf("Expected vector to have x=%.2f, got x=%.2f", expectedX, v.X())
	}
}

func TestBoundDirection(t *testing.T) {
	helpBoundDirection(t, 1, 0, 0)
	helpBoundDirection(t, 1, math.Pi, math.Pi)
	helpBoundDirection(t, 1, 2*math.Pi, 0)
	helpBoundDirection(t, 1, 2*math.Pi+0.1, 0.1)
	helpBoundDirection(t, 1, 4*math.Pi+0.1, 0.1)
	helpBoundDirection(t, 1, 4*math.Pi+math.Pi/2, math.Pi/2)
}

func helpBoundDirection(t *testing.T, len float64, dir float64, expectedDir float64) {
	v := Vector{len: len, dir: dir}
	v.BoundDirection()

	if !WithinEps(v.dir, expectedDir) {
		t.Errorf("Expected vector to have dir=%.2f, got dir=%.2f", expectedDir, v.dir)
	}
}
