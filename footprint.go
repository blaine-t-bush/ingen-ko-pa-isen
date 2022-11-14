package main

import (
	"time"

	"github.com/google/uuid"
)

func (a *Actor) UpdateDistanceSinceLastFootprint(distance float64, replace bool) {
	if replace {
		*a.distanceSinceLastFootprint = distance
	} else {
		*a.distanceSinceLastFootprint += distance
	}
}

func (g *Game) AddFootprint(a Actor, v Vector) {
	if *a.distanceSinceLastFootprint > FootprintSpacing {
		w, h := footprintIceImage.Size()
		coord := ScreenCoordinate{a.BoundingBox().CenterX(), a.BoundingBox().Bottom()}
		image := footprintIceImage
		if g.CoordinateIsOnTerrainType(coord, TerrainTypeIce) {
			image = footprintIceImage
		} else if g.CoordinateIsOnTerrainType(coord, TerrainTypeSnow) {
			image = footprintSnowImage
		}

		g.footprints = append(g.footprints, &Object{
			id:          uuid.NewString(),
			image:       image,
			pos:         &coord,
			width:       float64(w),
			height:      float64(h),
			collidable:  false,
			aboveActors: false,
			belongsTo:   a.id,
			createdAt:   time.Now().UnixMilli(),
		})

		a.UpdateDistanceSinceLastFootprint(0, true)
	} else {
		a.UpdateDistanceSinceLastFootprint(v.len, false)
	}
}

func (g *Game) UpdateFootprints() {
	footprintsCopy := []*Object{}
	for _, footprint := range g.footprints {
		if time.Now().UnixMilli()-footprint.createdAt < FootprintLifetimeMs*1000 {
			footprintsCopy = append(footprintsCopy, footprint)
		}
	}

	g.footprints = footprintsCopy
}
