package CellularAutomata

import (
	"fmt"
	"gifhelper"
	"os"
	"strconv"
)

// command-line parameters are stored in an array of strings called os.Args
//its length is equal to #parameters +1
//os.Args[0] is name of program
//os.Args[1] is first parameter given
//os.Args[2] is second parameter given ...
//...
//os.Args[len(os.Args)-1] is final parameter given

func main() {
	fmt.Println("Cellular automata!")

	neighborhood := os.Args[1]     //"vonNeumann" or "Moore"
	ruleFile := os.Args[2]         //where are rule strings found?
	initialBoardFile := os.Args[3] // my starting GameBoard file name
	outputFile := os.Args[4]       // where to draw the final animated GIF of boards

	// how many pixels wide should each cell be?
	cellWidth, err := strconv.Atoi(os.Args[5])
	if err != nil {
		panic("Error: Problem converting cell width parameter to an integer.")
	}

	// how many generations to play the automaton?
	numGens, err2 := strconv.Atoi(os.Args[6])
	if err2 != nil {
		panic("Error: Problem converting number of generations to an integer.")
	}

	fmt.Println("Parameters read in successfully!")

	ruleStrings := ReadRulesFromFile(ruleFile)
	fmt.Println("Rules are read in successfully!")

	initialBoard := ReadBoardFromFile(initialBoardFile)

	fmt.Println("Playing the automaton.")

	boards := PlayAutomaton(initialBoard, numGens, neighborhood, ruleStrings)

	fmt.Println("Automaton played. Now, drawing images.")

	// we need a slice of image objects
	imglist := DrawGameBoards(boards, cellWidth)
	fmt.Println("Boards drawn to images! Now, convert to animated GIF.")

	// convert images to a GIF
	gifhelper.ImagesToGIF(imglist, outputFile)

	fmt.Println("Success! GIF produced.")
}
