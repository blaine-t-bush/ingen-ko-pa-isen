package main

import "math"

const (
	eps = 0.0001
)

func WithinEps(a float64, b float64) bool {
	return math.Abs(a-b) <= eps
}
