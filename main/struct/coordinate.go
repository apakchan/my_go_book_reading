package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type myLocation struct {
	lat, long float64
}

type world struct {
	radius float64
}

type report struct {
	sol      int
	temp     temperature
	location myLocation
}

type temperature struct {
	high, low celsius
}

type celsius float64

type gps struct {
	cur, tar myLocation
	curWorld world
}

func (g gps) calculateDistance() float64 {
	return g.curWorld.distance(g.cur, g.tar)
}

func (g gps) message() {
	fmt.Printf("距离目的地还有 %f km\n", g.calculateDistance())
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newMyLocation(lat, long coordinate) myLocation {
	return myLocation{lat.decimal(), long.decimal()}
}

func (w world) distance(p1, p2 myLocation) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func newWorld(distance float64) world {
	return world{distance}
}

func (w world) calculateLandPoint(landPoints [][]myLocation) ([]myLocation, []myLocation) {
	if len(landPoints) == 0 {
		return nil, nil
	}
	minPair := landPoints[0]
	maxPair := landPoints[0]
	for _, locationPair := range landPoints {
		if w.distance(minPair[0], minPair[1]) > w.distance(locationPair[0], locationPair[1]) {
			minPair = locationPair
		}
		if w.distance(maxPair[0], maxPair[1]) < w.distance(locationPair[0], locationPair[1]) {
			maxPair = locationPair
		}
	}
	return minPair, maxPair
}

func main() {
	g := gps{
		cur:      myLocation{lat: -4.5895, long: 137.4417},
		tar:      myLocation{lat: 4.5, long: 135.9},
		curWorld: newWorld(3389.5),
	}
	g.message()
}