package main

import (
	"math"
	"testing"
)

const (
	eps = 0.0001
)

func withinEps(a float64, b float64) bool {
	return math.Abs(a-b) <= eps
}

func TestX(t *testing.T) {
	var v Vector

	v = Vector{len: 1, dir: 0}
	if !withinEps(v.X(), 1) {
		t.Errorf("Expected vector to have x=%d, got x=%.2f", 1, v.X())
	}

	v = Vector{len: 1, dir: math.Pi}
	if !withinEps(v.X(), -1) {
		t.Errorf("Expected vector to have x=%d, got x=%.2f", -1, v.X())
	}

	v = Vector{len: 1, dir: math.Pi / 2}
	if !withinEps(v.X(), 0) {
		t.Errorf("Expected vector to have x=%d, got x=%.2f", 0, v.X())
	}
}

func TestY(t *testing.T) {
	var v Vector

	v = Vector{len: 1, dir: 0}
	if !withinEps(v.Y(), 0) {
		t.Errorf("Expected vector to have y=%d, got y=%.2f", 0, v.Y())
	}

	v = Vector{len: 1, dir: math.Pi / 2}
	if !withinEps(v.Y(), 1) {
		t.Errorf("Expected vector to have y=%d, got y=%.2f", 1, v.Y())
	}

	v = Vector{len: 1, dir: math.Pi}
	if !withinEps(v.Y(), 0) {
		t.Errorf("Expected vector to have y=%d, got y=%.2f", 0, v.Y())
	}
}
