package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	/*
		  //this is all nonsense
			n := 40
		  factorial1 := go Perm(1, 21)
		  factorial2 := go Perm(21,41)
		  fmt.Println(factorial1 * factorial2)
	*/

	TimingCraps()

}

func TimingCraps() {

	numTrials := 20000000
	numProcs := runtime.NumCPU()

	start := time.Now()
	CrapsHouseEdgeMultiProc(numTrials, numProcs)
	elapsed := time.Since(start)
	log.Printf("Simulating craps in parallel took %s", elapsed)

	start2 := time.Now()
	CrapsHouseEdgeSerial(numTrials)
	elapsed2 := time.Since(start2)
	log.Printf("Simulating craps serially took %s", elapsed2)
}

func Push(k int, c chan int) {
	time.Sleep(time.Second)
	c <- k
}

func BufferedChannels() {
	n := 10
	c := make(chan int, n) //capacity of our "buffered" channel
	//we can get capacity of channel with cap() function
	fmt.Println("Capacity:", cap(c))

	// buffered channels are not synchronous, meaning you can put
	// stuff into channel without having to take out at exact same time
	for k := 0; k < n; k++ {
		go Push(k, c)
	}
	// this is not happening synchronously! (because it's buffered)

	for i := 0; i < n; i++ {
		fmt.Println(<-c) // will wait until there's something in channel if a process is running
	}
}

func SummingParallel() {
	//declare a slice of ints
	a := make([]int, 10000000)
	for i := range a {
		a[i] = i + 1
	}

	numProcs := runtime.NumCPU()

	start := time.Now()
	SumMultiProc(a, numProcs)
	elapsed := time.Since(start)
	log.Printf("Summing in parallel took %s", elapsed)

	start2 := time.Now()
	SumSerial(a)
	elapsed2 := time.Since(start2)
	log.Printf("Summing serially took %s", elapsed2)

}

func SumSerial(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}

func SumMultiProc(a []int, numProcs int) int {
	n := len(a)
	s := 0
	c := make(chan int)

	//split the array into numProcs approximately equal pieces
	for i := 0; i < numProcs; i++ {
		startIndex := i * (n / numProcs)
		endIndex := (i + 1) * (n / numProcs)
		if i < numProcs-1 {
			//normal case
			go Sum(a[startIndex:endIndex], c)
		} else { // i == numProcs - 1
			//end of the slice -- make sure you go to the very end
			go Sum(a[startIndex:], c)
		}
	}

	//get values from channel numProcs times and add them to the growing sum s
	for i := 0; i < numProcs; i++ {
		s += <-c
	}

	return s
}

func Sum(a []int, c chan int) {
	s := 0
	for _, v := range a {
		s += v
	}
	c <- s
}

func ParallelFactorial() {
	c := make(chan int)
	go PermChannel(1, 11, c)
	go PermChannel(11, 21, c)
	fmt.Println(<-c * <-c)
}

func PermChannel(k, n int, c chan int) {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	c <- p
}

func BasicChannels() {

	//we can force Go to use a different max number of procs
	runtime.GOMAXPROCS(1)

	//channels store a value of a given type and allow functions to communicate with each other
	c := make(chan string) // this channel is "synchronous"
	//c <- "Hello" // the channel blocks, meaning that it will not continue in this serial process until someone is on the other end of the channel ready to receive the message
	go SayHi(c)
	fmt.Println(<-c)

	fmt.Println("Exiting normally.")
}

func SayHi(c chan string) {
	fmt.Println("Yo")
	c <- "Hello!"
	//only block what comes after this in the subroutine
	//(which is nothing)
}

func Perm(k, n int) int {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	return p
}

func PrintFactorials(n int) {
	p := 1
	for i := 1; i <= n; i++ {
		fmt.Println(p)
		p *= i
	}
}

func NumProcessor() {
	fmt.Println("Parallel and concurrent programming.")

	//Go will tell us how many processors we have.
	fmt.Println("Num processors:", runtime.NumCPU())

	n := 100000000

	start := time.Now()
	Factorial(n)
	elapsed := time.Since(start)
	log.Printf("Multiprocessors took %s", elapsed)

	//we can force Go to use a different max number of procs
	runtime.GOMAXPROCS(1)

	start2 := time.Now()
	Factorial(n)
	elapsed2 := time.Since(start2)
	log.Printf("Single processor took %s", elapsed2)
}

func Factorial(n int) int {
	prod := 1
	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod
}
