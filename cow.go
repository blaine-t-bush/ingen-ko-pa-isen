package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CowSpeedMultiplier            = 2
	CowSpeedMin                   = 0
	CowSpeedMax                   = 2
	FarmerDetectionRadius         = 100
	WallDetectionRadius           = 5
	CowNoiseSize                  = 100
	CowDirectionChangeProbability = 0.05
	CowSpeedChangeProbability     = 0.05
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
		dir = cow.noiseDir.UpdateAndGetValueScaled(2 * math.Pi)
	}

	// Moves away from walls.
	if cow.sprite.pos.x <= WallDetectionRadius {
		dir = 0.0
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.sprite.pos.x+cow.sprite.imageWidth >= screenWidth-WallDetectionRadius {
		dir = math.Pi
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.sprite.pos.y <= WallDetectionRadius {
		dir = math.Pi / 2
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.sprite.pos.y+cow.sprite.imageHeight >= screenHeight-WallDetectionRadius {
		dir = 3 * math.Pi / 2
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	}

	// Flees directly away from farmer.
	if cow.sprite.pos.WithinRadius(farmer.sprite.Center(), FarmerDetectionRadius) {
		dir = VectorFromPoints(farmer.sprite.Center(), cow.sprite.Center()).dir
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	}

	return dir
}

func ChooseCowSpeed(cow *Actor) float64 {
	speed := cow.velocity.len

	if rand.Float64() >= 1-CowSpeedChangeProbability {
		speed = cow.noiseSpeed.UpdateAndGetValueScaled(CowSpeedMultiplier)
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
			dir: noiseDir.GetValueScaled(2 * math.Pi),
			len: noiseSpeed.GetValueScaled(CowSpeedMultiplier),
		},
		noiseSpeed: &noiseSpeed,
		noiseDir:   &noiseDir,
	}
}
