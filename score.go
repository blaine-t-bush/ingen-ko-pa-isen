package main

func (g *Game) UpdateScore() {
	g.numberOfCowsOnIce = g.NumberOfCowsOnIce()
}

func (g *Game) NumberOfCows() int {
	count := 0
	for range g.cows {
		count++
	}

	return count
}

func (g *Game) NumberOfCowsOnIce() int {
	count := 0
	for _, cow := range g.cows {
		if g.CoordinateIsOnTerrainType(cow.BoundingBox().BottomCenter(), TerrainTypeIce) {
			count++
		}
	}

	return count
}
