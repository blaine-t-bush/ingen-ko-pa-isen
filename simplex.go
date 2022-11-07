package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ojrac/opensimplex-go"
)

type Noise struct {
	image [][]float64
	w     int
	h     int
	x     int
	y     int
}

func GenerateNoise(w, h int) Noise {
	noise := opensimplex.NewNormalized(rand.Int63())
	image := make([][]float64, h)
	for i := range image {
		image[i] = make([]float64, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)
			image[y][x] = noise.Eval2(xFloat, yFloat)
		}
	}

	return Noise{image: image, w: w, h: h, x: rand.Intn(w), y: rand.Intn(h)}
}

func (n *Noise) GetValue() float64 {
	return n.image[n.y][n.x]
}

func (n *Noise) GetValueScaled(scale float64) float64 {
	return scale * n.image[n.y][n.x]
}

func (n *Noise) UpdateAndGetValue() float64 {
	n.UpdateCoordinate()
	return n.GetValue()
}

func (n *Noise) UpdateAndGetValueScaled(scale float64) float64 {
	n.UpdateCoordinate()
	return n.GetValueScaled(scale)
}

func (n *Noise) UpdateCoordinate() {
	newX := n.x
	newY := n.y

	switch rand.Intn(4) {
	case 0:
		newX++
	case 1:
		newX--
	case 2:
		newY++
	case 3:
		newY--
	}

	if newX < 0 {
		newX = 0
	} else if newX >= n.w {
		newX = n.w - 1
	}

	if newY < 0 {
		newY = 0
	} else if newY >= n.h {
		newY = n.h - 1
	}

	n.x = newX
	n.y = newY
}

func (n *Noise) ChangeCoordinateToMatch(value float64) {
	diff := value
	closestX := n.x
	closestY := n.y

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			if math.Abs(n.image[y][x]-value) < diff {
				diff = math.Abs(n.image[y][x] - value)
				closestX = x
				closestY = y
				fmt.Println("ping")
			}
		}
	}

	n.x = closestX
	n.y = closestY
}
