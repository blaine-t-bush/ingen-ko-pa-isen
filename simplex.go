package main

import (
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

func GenerateNoise(w, h int, minValue, maxValue float64) Noise {
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

	noiseStruct := Noise{image: image, w: w, h: h, x: rand.Intn(w), y: rand.Intn(h)}
	noiseStruct.Normalize()
	noiseStruct.ShiftMean((maxValue + minValue) / 2)
	noiseStruct.Rescale(minValue, maxValue)

	return noiseStruct
}

func (n *Noise) ToPixels() []byte {
	pixels := []byte{}
	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			byteVal := byte(n.image[y][x] * 255)
			pixels = append(pixels, byteVal)
			pixels = append(pixels, byteVal)
			pixels = append(pixels, byteVal)
			pixels = append(pixels, 1)
		}
	}

	return pixels
}

func (n *Noise) Min() float64 {
	min := 1.0

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			if n.image[y][x] < min {
				min = n.image[y][x]
			}
		}
	}

	return min
}

func (n *Noise) Max() float64 {
	max := 1.0

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			if n.image[y][x] > max {
				max = n.image[y][x]
			}
		}
	}

	return max
}

func (n *Noise) Normalize() {
	copy := *n
	min := n.Min()
	max := n.Max()
	delta := max - min

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			n.image[y][x] = (copy.image[y][x] - min) / delta
		}
	}
}

func (n *Noise) ShiftMean(newMean float64) {
	sum := 0.0

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			sum += n.image[y][x]
		}
	}

	mean := sum / float64(n.h*n.w)
	meanDiff := mean - newMean

	copy := *n
	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			n.image[y][x] = (copy.image[y][x] - meanDiff)
		}
	}
}

func (n *Noise) Rescale(newMin, newMax float64) {
	copy := *n
	min := n.Min()
	max := n.Max()
	delta := max - min
	newDelta := newMax - newMin

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			n.image[y][x] = newDelta * copy.image[y][x] / delta
		}
	}
}

func (n *Noise) Smooth() {
	copy := *n

	for y := 1; y < n.h-1; y++ {
		for x := 1; x < n.w-1; x++ {
			n.image[y][x] = (copy.image[y-1][x] + copy.image[y+1][x] + copy.image[y][x-1] + copy.image[y][x+1]) / 4
		}
	}
}

func (n *Noise) GetValue() float64 {
	return n.image[n.y][n.x]
}

func (n *Noise) GetValueScaled(scale float64) float64 {
	return scale * n.image[n.y][n.x]
}

func (n *Noise) UpdateAndGetValue() float64 {
	n.SelectRandomNearbyCoordinate()
	return n.GetValue()
}

func (n *Noise) UpdateAndGetValueScaled(scale float64) float64 {
	n.SelectRandomNearbyCoordinate()
	return n.GetValueScaled(scale)
}

func (n *Noise) SelectCoordinate(x int, y int) {
	n.x = x
	n.y = y
}

func (n *Noise) SelectRandomNearbyCoordinate() {
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

	n.SelectCoordinate(newX, newY)
}

func (n *Noise) SelectCoordinateToMatch(value float64) {
	diff := 1.0
	closestX := n.x
	closestY := n.y

	for y := 0; y < n.h; y++ {
		for x := 0; x < n.w; x++ {
			if math.Abs(n.image[y][x]-value) < diff {
				diff = math.Abs(n.image[y][x] - value)
				closestX = x
				closestY = y
			}
		}
	}

	n.SelectCoordinate(closestX, closestY)
}
