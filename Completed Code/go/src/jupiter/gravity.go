package main

import (
	"math"
)

//let's place our gravity simulation functions here.

//SimulateGravity simulates gravity over a series of snap shots separated by equal unit time.
//Input: an initial Universe, a number of generations, and a time parameter (in seconds).
//Output: a slice of Universe objects corresponding to simulating the
//force of gravity over the number of generations time points.
func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	timePoints := make([]Universe, numGens+1)
	timePoints[0] = initialUniverse

	//now range over the number of generations and update the universe each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time)
	}

	return timePoints
}

/*
UpdateUniverse(currentUniverse, time)
	newUniverse  CopyUniverse(currentUniverse)
	for every body b in newUniverse
		b.acceleration  UpdateAccel(newUniverse, b)
		b.velocity  UpdateVelocity(b, time)
		b.position  UpdatePosition(b, time)
	return newUniverse
*/

//UpdateUniverse updates a given universe over a specified time interval (in seconds).
//Input: A Universe object and a float time.
//Output: A Universe object corresponding to simulating gravity over time seconds, assuming that acceleration is constant over this time.
func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	newUniverse := CopyUniverse(currentUniverse)

	//range over all bodies in the universe and update their acceleration,
	//velocity, and position
	for _, b := range newUniverse.bodies {
		b.acceleration = UpdateAcceleration(newUniverse, b)
		b.velocity = UpdateVelocity(b, time)
		b.position = UpdatePosition(b, time)
	}

	return newUniverse
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
