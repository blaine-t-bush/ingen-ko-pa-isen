package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FarmerSpeedMultiplier = 4
	FarmerSpeedMax        = 4
)

func (g *Game) UpdateFarmer() {
	g.MoveActor(*g.farmer, *g.farmer.velocityDesired)
	g.farmer.Shunt()
	g.farmer.velocityDesired = &Vector{0, 0}
}

func (g *Game) CreateFarmer(img ebiten.Image) *Actor {
	w, h := img.Size()
	boundingBox := &BoundingBox{
		pos:    ScreenCoordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
		width:  float64(w),
		height: float64(h),
	}

	for {
		if g.CheckCollision(*boundingBox) {
			boundingBox.pos = ScreenCoordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))}
		} else {
			break
		}
	}

	distanceSinceLastFootprint := 0.0

	return &Actor{
		image:                      &img,
		pos:                        &boundingBox.pos,
		width:                      boundingBox.width,
		height:                     boundingBox.height,
		distanceSinceLastFootprint: &distanceSinceLastFootprint,
		speedMax:                   FarmerSpeedMax,
		speedMultiplier:            FarmerSpeedMultiplier,
		velocityActual:             &Vector{0, 0},
		velocityDesired:            &Vector{0, 0},
	}
}
