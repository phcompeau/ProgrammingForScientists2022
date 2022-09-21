package main

import (
	"math"
)

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

//UpdateUniverse updates a given universe over a specified time interval (in seconds).
//Input: A Universe object and a float time.
//Output: A Universe object corresponding to simulating gravity over time seconds, assuming that acceleration is constant over this time.
func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	newUniverse := CopyUniverse(currentUniverse)

	//range over all bodies in the universe and update their acceleration,
	//velocity, and position
	for i := range newUniverse.bodies {
		newUniverse.bodies[i].acceleration = UpdateAcceleration(currentUniverse, newUniverse.bodies[i])
		newUniverse.bodies[i].velocity = UpdateVelocity(newUniverse.bodies[i], time)
		newUniverse.bodies[i].position = UpdatePosition(newUniverse.bodies[i], time)
	}

	return newUniverse
}

//UpdateAcceleration
//Input: Universe object and a body B
//Output: The net acceleration on B due to gravity calculated
//by every body in the Universe
func UpdateAcceleration(currentUniverse Universe, b Body) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(currentUniverse.bodies, b)

	//now, calculate acceleration (F = ma)
	accel.x = force.x / b.mass
	accel.y = force.y / b.mass

	return accel
}

//ComputeNetForce
//Input: A slice of body objects and an individual body
//Output: The net force vector (OrderedPair) acting on the given body
//due to gravity from all other bodies in the Universe
func ComputeNetForce(bodies []Body, b Body) OrderedPair {
	var netForce OrderedPair

	for i := range bodies {
		//only do a force computation if current body is not the input Body
		if bodies[i] != b {
			force := ComputeForce(b, bodies[i])

			//now add its components into net force components
			netForce.x += force.x
			netForce.y += force.y
		}
	}

	return netForce
}

//ComputeForce
//Input: Two body objects b1 and b2.
//Output: The force due to gravity (as a vector) acting on b1 subject to b2.
func ComputeForce(b1, b2 Body) OrderedPair {
	var force OrderedPair

	dist := Distance(b1.position, b2.position)

	F := G * b1.mass * b2.mass / (dist * dist) // magnitude of gravity
	deltaX := b2.position.x - b1.position.x
	deltaY := b2.position.y - b1.position.y

	force.x = F * deltaX / dist //deltaX/dist = cos theta
	force.y = F * deltaY / dist //deltaY/dist = sin theta

	return force
}

//CopyUniverse
//Input: a Universe object
//Output: a new Universe object, all of whose fields are copied over
//into the new Universe's fields. (Deep copy.)
func CopyUniverse(currentUniverse Universe) Universe {
	var newUniverse Universe

	newUniverse.width = currentUniverse.width

	//let's make the new universe's slice of Body objects
	numBodies := len(currentUniverse.bodies)
	newUniverse.bodies = make([]Body, numBodies)

	//now, copy all of the bodies' fields into our new bodies
	for i := range currentUniverse.bodies {
		newUniverse.bodies[i].name = currentUniverse.bodies[i].name
		newUniverse.bodies[i].mass = currentUniverse.bodies[i].mass
		newUniverse.bodies[i].radius = currentUniverse.bodies[i].radius
		newUniverse.bodies[i].position.x = currentUniverse.bodies[i].position.x
		newUniverse.bodies[i].position.y = currentUniverse.bodies[i].position.y
		newUniverse.bodies[i].velocity.x = currentUniverse.bodies[i].velocity.x
		newUniverse.bodies[i].velocity.y = currentUniverse.bodies[i].velocity.y
		newUniverse.bodies[i].acceleration.x = currentUniverse.bodies[i].acceleration.x
		newUniverse.bodies[i].acceleration.y = currentUniverse.bodies[i].acceleration.y
		newUniverse.bodies[i].red = currentUniverse.bodies[i].red
		newUniverse.bodies[i].green = currentUniverse.bodies[i].green
		newUniverse.bodies[i].blue = currentUniverse.bodies[i].blue
	}

	return newUniverse
}

//UpdateVelocity
//Input: a Body object and a time step (float64).
//Output: The orderedPair corresponding to the velocity of this object
//after a single time step, using the body's current acceleration.
func UpdateVelocity(b Body, time float64) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

	return vel
}

//UpdatePosition
//Input: a Body b and a time step (float64).
//Output: The OrderedPair corresponding to the updated position of the body after a single time step, using the body's current acceleration and velocity.
func UpdatePosition(b Body, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	pos.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

	return pos
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
