package main

import (
	"math/rand"
	"time"
)

//this is where we will put functions that correspond only to the parallel simulation.

// DiffuseParallel is a Board method that diffuses each Particle in the Board over a single
// time step, concurrently over numProcs processes.
func (b *Board) DiffuseParallel(numProcs int) {
	numParticles := len(b.particles)

	finished := make(chan bool, numProcs)

	//split the work over numProcs processes.
	for i := 0; i < numProcs; i++ {
		//each processor gets approx. numParticles/numProcs particles
		startIndex := i * numParticles / numProcs
		var endIndex int
		if i < numProcs-1 {
			endIndex = (i + 1) * numParticles / numProcs
		} else {
			endIndex = numParticles
		}
		//don't want a race condition where all processes share
		//a single PRNG object.
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source) // creates new PRNG object
		go DiffuseOneProc(b.particles[startIndex:endIndex], generator, finished)
	}

	// we need to receive a message from all our channels that they're finished
	for i := 0; i < numProcs; i++ {
		<-finished
	}

	// the function has to be finished here
}

func DiffuseOneProc(particles []*Particle, generator *(rand.Rand), finished chan bool) {

	for _, p := range particles {
		p.RandStep(generator)
	}

	//function is done. Indicate this by placing true (or false) into channel
	finished <- true
}
