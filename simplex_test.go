package main

import (
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
