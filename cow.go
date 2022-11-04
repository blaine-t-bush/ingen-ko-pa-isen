package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CowSpeedMultiplier    = 2
	FarmerDetectionRadius = 120
	WallDetectionRadius   = 5
)

func (g *Game) UpdateCows() {
	for _, cow := range g.cows {
		// possibly choose a new wander direction
		cow.velocity.dir = ChooseCowDirection(cow, g.farmer)

		// update position based on velocity
		g.MoveActor(*cow, *cow.velocity, CowSpeedMultiplier)

		// shunt cows if they've managed to move past borders
		cow.Shunt()
	}
}

func ChooseCowDirection(cow *Actor, farmer *Actor) float64 {
	dir := cow.velocity.dir

	// Small chance to choose new direction.
	if rand.Float64() >= 0.95 {
		dir = cow.velocity.dir + (math.Pi/2)*(0.5-rand.Float64())
	}

	// Moves away from walls.
	if cow.sprite.pos.x <= WallDetectionRadius {
		dir = 0
	} else if cow.sprite.pos.x+cow.sprite.imageWidth >= screenWidth-WallDetectionRadius {
		dir = math.Pi
	} else if cow.sprite.pos.y <= WallDetectionRadius {
		dir = math.Pi / 2
	} else if cow.sprite.pos.y+cow.sprite.imageHeight >= screenHeight-WallDetectionRadius {
		dir = -math.Pi / 2
	}

	// Flees directly away from farmer.
	if cow.sprite.pos.WithinRadius(farmer.sprite.Center(), FarmerDetectionRadius) {
		dir = VectorFromPoints(farmer.sprite.Center(), cow.sprite.Center()).dir
	}

	return dir
}

func (g *Game) CreateRandomCow(img ebiten.Image) *Actor {
	w, h := img.Size()
	potentialSprite := &Sprite{
		image:       &img,
		imageWidth:  float64(w),
		imageHeight: float64(h),
		pos:         &Coordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
	}

	for {
		if g.CheckCollision(potentialSprite) {
			potentialSprite = &Sprite{
				image:       &img,
				imageWidth:  float64(w),
				imageHeight: float64(h),
				pos:         &Coordinate{float64(rand.Intn(screenWidth - w)), float64(rand.Intn(screenHeight - h))},
			}
		} else {
			break
		}
	}

	return &Actor{
		sprite: potentialSprite,
		velocity: &Vector{
			dir: 2 * math.Pi * rand.Float64(),
			len: 1,
		},
	}
}
