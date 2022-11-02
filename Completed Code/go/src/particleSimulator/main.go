package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Particle simulator.")

	fmt.Println("Generating random particles and initializing board.")

	numParticles := 20000
	boardWidth := 1000.0
	boardHeight := 1000.0
	particleRadius := 5.0
	diffusionRate := 1.0

	//assumption: all particles are white

	random := false // make true if we want to scatter across board

	initialBoard := InitializeBoard(boardWidth, boardHeight, numParticles, particleRadius, diffusionRate, random)

	numSteps := 2000

	start := time.Now()
	UpdateBoards(initialBoard, numSteps, true)
	elapsed := time.Since(start)
	log.Printf("Simulating diffusion in parallel took %s", elapsed)

	start2 := time.Now()
	UpdateBoards(initialBoard, numSteps, false)
	elapsed2 := time.Since(start2)
	log.Printf("Simulating diffusion serially took %s", elapsed2)

	/*
		fmt.Println("Running simulation.")

		isParallel := true

		boards := UpdateBoards(initialBoard, numSteps, isParallel)

		fmt.Println("Simulation run. Animating system.")
		canvasWidth := 300
		frequency := 10
		images := AnimateSystem(boards, canvasWidth, frequency)

		fmt.Println("Images drawn. Generating GIF.")

		outFileName := "diffusion"
		gifhelper.ImagesToGIF(images, outFileName)
	*/
}
