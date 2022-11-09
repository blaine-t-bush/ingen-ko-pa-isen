package main

import (
	"time"

	"github.com/google/uuid"
)

func (g *Game) AddFootprint(a Actor) {
	if int(*a.distanceMoved)%FootprintSpacing == 0 {
		w, h := footprintIceImage.Size()
		g.footprints = append(g.footprints, &Object{
			id:          uuid.NewString(),
			image:       footprintIceImage,
			pos:         &ScreenCoordinate{a.BoundingBox().CenterX(), a.BoundingBox().Bottom()},
			width:       float64(w),
			height:      float64(h),
			collidable:  false,
			aboveActors: false,
			belongsTo:   a.id,
			createdAt:   time.Now().UnixMilli(),
		})
	}
}

func (g *Game) UpdateFootprints() {
	footprintsCopy := []*Object{}
	for _, footprint := range g.footprints {
		if time.Now().UnixMilli()-footprint.createdAt < FootprintLifetime*1000 {
			footprintsCopy = append(footprintsCopy, footprint)
		}
	}

	g.footprints = footprintsCopy
}
