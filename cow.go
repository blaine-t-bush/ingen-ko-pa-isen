package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CowSpeedMultiplier             = 2
	CowSpeedMin                    = 0
	CowSpeedMax                    = 2
	FarmerDetectionRadius          = 120
	WallDetectionRadius            = 5
	CowNoiseSize                   = 50
	CowNoiseDirectionModifierScale = 0.8
	CowNoiseSpeedModifierScale     = 0.1
	CowDirectionChangeProbability  = 0.05
	CowSpeedChangeProbability      = 0.05
)

func (g *Game) UpdateCows() {
	for _, cow := range g.cows {
		// possibly choose a new wander direction
		cow.velocity.dir = ChooseCowDirection(cow, g.farmer)
		cow.velocity.len = ChooseCowSpeed(cow)

		// update position based on velocity
		g.MoveActor(*cow, *cow.velocity, CowSpeedMultiplier)

		// shunt cows if they've managed to move past borders
		cow.Shunt()
	}
}

func ChooseCowDirection(cow *Actor, farmer *Actor) float64 {
	dir := cow.velocity.dir

	// Small chance to choose new direction.
	if rand.Float64() >= 1-CowDirectionChangeProbability {
		dir = 2 * math.Pi * cow.noiseDir.UpdateAndGetValue()
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

func ChooseCowSpeed(cow *Actor) float64 {
	speed := cow.velocity.len

	if rand.Float64() >= 1-CowSpeedChangeProbability {
		speed = CowSpeedMultiplier * cow.noiseSpeed.UpdateAndGetValue()
	}

	if speed < CowSpeedMin {
		speed = CowSpeedMin
	} else if speed > CowSpeedMax {
		speed = CowSpeedMax
	}

	return speed
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

	noiseSpeed := GenerateNoise(CowNoiseSize, CowNoiseSize)
	noiseDir := GenerateNoise(CowNoiseSize, CowNoiseSize)

	return &Actor{
		sprite: potentialSprite,
		velocity: &Vector{
			dir: 2 * math.Pi * rand.Float64(),
			len: float64(CowSpeedMax) / 2,
		},
		noiseSpeed: &noiseSpeed,
		noiseDir:   &noiseDir,
	}
}
