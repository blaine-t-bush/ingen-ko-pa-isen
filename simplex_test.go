package main

import (
	"math/rand"
	"testing"
)

func TestGenerateNoise(t *testing.T) {
	w := 50
	h := 50
	n := GenerateNoise(50, 50)

	if len(n.image) != w {
		t.Errorf("Expected noise to have width %d, got width %d", 1, len(n.image))
	}

	if len(n.image[0]) != h {
		t.Errorf("Expected noise to have height %d, got height %d", 1, len(n.image[0]))
	}
}

func TestGetValue(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := GenerateNoise(50, 50)
		for x := 0; x < 50; x++ {
			for y := 0; y < 50; y++ {
				n.SelectCoordinate(x, y)
				value := n.GetValue()
				if value < 0.0 {
					t.Errorf("Expected noise to have value >= 0, got value %.2f", value)
				} else if value >= 1.0 {
					t.Errorf("Expected noise to have value < 1, got value %.2f", value)
				}
			}
		}
	}
}

func TestGetValueScaled(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := GenerateNoise(50, 50)
		scale := 50 * rand.Float64()
		for x := 0; x < 50; x++ {
			for y := 0; y < 50; y++ {
				n.SelectCoordinate(x, y)
				value := n.GetValueScaled(scale)
				if value < 0.0 {
					t.Errorf("Expected noise to have value >= 0, got value %.2f", value)
				} else if value >= scale {
					t.Errorf("Expected noise to have value < 1, got value %.2f", value)
				}
			}
		}
	}
}

func TestSelectCoordinate(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := GenerateNoise(50, 50)
		for x := 0; x < 50; x++ {
			for y := 0; y < 50; y++ {
				n.SelectCoordinate(x, y)
				if n.x != x {
					t.Errorf("Expected coordinate to change to x=%d, got %d", x, n.x)
				}
				if n.y != y {
					t.Errorf("Expected coordinate to change to y=%d, got %d", y, n.y)
				}
			}
		}
	}
}

func TestSelectRandomNearbyCoordinate(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := GenerateNoise(50, 50)
		for j := 0; j < 100; j++ {
			prevX := n.x
			prevY := n.y
			n.SelectRandomNearbyCoordinate()
			if n.x-prevX > 1 || n.x-prevX < -1 {
				t.Errorf("Expected to select x within 1 step of x=%d but got x=%d", prevX, n.x)
			}
			if n.y-prevY > 1 || n.y-prevY < -1 {
				t.Errorf("Expected to select y within 1 step of y=%d but got y=%d", prevY, n.y)
			}
		}
	}
}
