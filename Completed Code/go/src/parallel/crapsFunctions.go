package main

import (
	"math/rand"
	"time"
)

func CrapsHouseEdgeMultiProc(numTrials, numProcs int) float64 {
	winnings := 0 // amount won or lost

	c := make(chan int)

	//play the game concurrently over numProcs processes
	for i := 0; i < numProcs; i++ {
		//create a new PRNG object that only this goroutine has access to
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source) // PRNG object
		go TotalWinOneProc(numTrials/numProcs, generator, c)
	}

	//grab all values from channel
	for i := 0; i < numProcs; i++ {
		winnings += <-c
	}

	return float64(winnings) / float64(numTrials)
}

func CrapsHouseEdgeSerial(numTrials int) float64 {
	count := 0
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source) // PRNG object
	for i := 0; i < numTrials; i++ {
		count += PlayCraps(generator) // true if win, false if loss
	}
	return float64(count) / float64(numTrials)
}

//TotalWinOneProc takes a number of trials and an integer channel as input.
//It simulates craps numTrials times. It then places
//the total winnings into the channel.
func TotalWinOneProc(numTrials int, generator *(rand.Rand), c chan int) {
	winnings := 0

	for i := 0; i < numTrials; i++ {
		winnings += PlayCraps(generator) // returns +1 or -1
	}

	c <- winnings
}

//PlayCrapsBetter takes a PRNG as input and returns true (win) or false (loss) corresponding to playing the game of craps once.
func PlayCraps(generator *(rand.Rand)) int {
	firstRoll := SumTwoDice(generator)
	if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		//loss!
		return -1
	} else if firstRoll == 7 || firstRoll == 11 {
		//win!
		return 1
	} else {
		// roll again until we hit 7 (loss) or firstRoll (win)
		for true {
			currentRoll := SumTwoDice(generator)
			if currentRoll == firstRoll { // win!
				return 1
			} else if currentRoll == 7 { // loss :(
				return -1
			}
		}
	}
	// this will never be hit
	panic("We should not be here.")
}

func SumTwoDice(generator *(rand.Rand)) int {
	return DieRoll(generator) + DieRoll(generator)
}

func DieRoll(generator *(rand.Rand)) int {
	return generator.Intn(6) + 1
}
