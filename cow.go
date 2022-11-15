package main

import (
	"math"
	"math/rand"

	"github.com/google/uuid"
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
		cow.velocityDesired.dir = ChooseCowDirection(cow, g.farmer)
		cow.velocityDesired.len = ChooseCowSpeed(cow)

		// update position based on velocity
		g.MoveActor(*cow, *cow.velocityDesired)

		// shunt cows if they've managed to move past borders
		cow.Shunt()
	}
}

func ChooseCowDirection(cow *Actor, farmer *Actor) float64 {
	dir := cow.velocityDesired.dir

	// Small chance to choose new direction.
	if rand.Float64() >= 1-CowDirectionChangeProbability {
		dir = cow.noiseDir.UpdateAndGetValueScaled(2 * math.Pi)
	}

	// Moves away from walls.
	if cow.BoundingBox().Left() <= WallDetectionRadius {
		dir = 0.0
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.BoundingBox().Right() >= screenWidth-WallDetectionRadius {
		dir = math.Pi
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.BoundingBox().Top() <= WallDetectionRadius {
		dir = math.Pi / 2
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	} else if cow.BoundingBox().Bottom() >= screenHeight-WallDetectionRadius {
		dir = 3 * math.Pi / 2
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	}

	// Flees directly away from farmer.
	if cow.BoundingBox().Center().WithinRadius(farmer.BoundingBox().Center(), FarmerDetectionRadius) {
		dir = VectorFromPoints(farmer.BoundingBox().Center(), cow.BoundingBox().Center()).dir
		cow.noiseDir.SelectCoordinateToMatch(dir / (2 * math.Pi))
	}

	return dir
}

func ChooseCowSpeed(cow *Actor) float64 {
	speed := cow.velocityDesired.len

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

func (g *Game) CreateRandomCow() *Actor {
	w, h := cowImage.Size()
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

	noiseSpeed := GenerateNoise(CowNoiseSize, CowNoiseSize)
	noiseDir := GenerateNoise(CowNoiseSize, CowNoiseSize)
	distanceSinceLastFootprint := 0.0
	currentSpeed := noiseDir.GetValueScaled(2 * math.Pi)
	currentDir := noiseSpeed.GetValueScaled(CowSpeedMultiplier)

	return &Actor{
		id:     uuid.NewString(),
		image:  cowImage,
		pos:    &boundingBox.pos,
		width:  boundingBox.width,
		height: boundingBox.height,
		velocityActual: &Vector{
			dir: currentSpeed,
			len: currentDir,
		},
		velocityDesired: &Vector{
			dir: currentSpeed,
			len: currentDir,
		},
		distanceSinceLastFootprint: &distanceSinceLastFootprint,
		speedMultiplier:            CowSpeedMultiplier,
		noiseSpeed:                 &noiseSpeed,
		noiseDir:                   &noiseDir,
	}
}
