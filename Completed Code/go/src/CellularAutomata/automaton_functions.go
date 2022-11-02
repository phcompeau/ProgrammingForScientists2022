package CellularAutomata

import "strconv"

type GameBoard [][]int

func PlayAutomaton(initialBoard GameBoard, numGens int, neighborhood string, ruleStrings []string) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1], neighborhood, ruleStrings)
	}

	return boards
}

func UpdateBoard(currBoard GameBoard, neighborhood string, ruleStrings []string) GameBoard {
	// first, create new board corresponding to the next generation.
	// let's have all cells have state 0 to begin.
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	//now, update values of newBoard
	//range through all cells of currBoard and update each one into newBoard.
	for r := 1; r < numRows-1; r++ {
		// r will range over rows of board
		// current row is currBoard[r]
		// range over values in currBoard[r]
		for c := 1; c < numCols-1; c++ {
			//curr value is currBoard[r][c]
			newBoard[r][c] = UpdateCell(currBoard, r, c, neighborhood, ruleStrings)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell takes a gameboard along with row and col indices and consults a collection of rule strings for a given neighborhood type to update cell state at given row and col indices.
func UpdateCell(board GameBoard, r, c int, neighborhood string, ruleStrings []string) int {
	// range through rule strings and look for a match ... very different for the two neighborhood types
	for _, rule := range ruleStrings {
		if RuleMatch(board, r, c, neighborhood, rule) == true {
			// match found! Return last element of current rule string
			finalSymbol := string(rule[len(rule)-1])
			// convert finalSymbol to int
			newState, err := strconv.Atoi(finalSymbol)
			if err != nil {
				panic("Error: problem in converting final rule string symbol.")
			}
			return newState
		}
	}
	// we need a default value
	panic("Error: we couldn't find a matching rule string to update cell.")
	return -1 // we know we didn't find a match
}

//RuleMatch takes a GameBoard board, row/col indices r, c, a neighborhood parameter (vonNeumann or Moore), and a rule string.
//It returns true if neighborhood of board[r][c] matches rule, and false otherwise.
func RuleMatch(board GameBoard, r, c int, neighborhood, rule string) bool {
	// branch based on neighborhood type.
	if neighborhood == "vonNeumann" {
		return RuleMatchVN(board, r, c, rule)
	} else if neighborhood == "Moore" {
		return RuleMatchMoore(board, r, c, rule)
	}
	panic("Error: invalid neighborhood type given to RuleMatch().")
	return false
}

func RuleMatchVN(board GameBoard, r, c int, rule string) bool {
	center := strconv.Itoa(board[r][c])
	north := strconv.Itoa(board[r-1][c])
	east := strconv.Itoa(board[r][c+1])
	south := strconv.Itoa(board[r+1][c])
	west := strconv.Itoa(board[r][c-1])
	if center != string(rule[0]) {
		return false
	}
	if north != string(rule[1]) {
		return false
	}
	if east != string(rule[2]) {
		return false
	}
	if south != string(rule[3]) {
		return false
	}
	if west != string(rule[4]) {
		return false
	}
	return true // we made it! rule must match
}

func RuleMatchMoore(board GameBoard, r, c int, rule string) bool {
	center := strconv.Itoa(board[r][c])
	northwest := strconv.Itoa(board[r-1][c-1])
	north := strconv.Itoa(board[r-1][c])
	northeast := strconv.Itoa(board[r-1][c+1])
	east := strconv.Itoa(board[r][c+1])
	southeast := strconv.Itoa(board[r+1][c+1])
	south := strconv.Itoa(board[r+1][c])
	southwest := strconv.Itoa(board[r+1][c-1])
	west := strconv.Itoa(board[r][c-1])

	// see if we find a mismatch with the current rule and any of the cell and its eight neighbors
	if center != string(rule[0]) || northwest != string(rule[1]) || north != string(rule[2]) || northeast != string(rule[3]) || east != string(rule[4]) || southeast != string(rule[5]) || south != string(rule[6]) || southwest != string(rule[7]) || west != string(rule[8]) {
		return false
	}
	// if we survived, the rule matches!
	return true
}

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

//InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and colums, where all values = 0.
func InitializeBoard(numRows, numCols int) GameBoard {
	// make a 2-D slice (default values = false)
	var board GameBoard
	board = make(GameBoard, numRows)
	// now we need to make the rows too
	for r := range board {
		board[r] = make([]int, numCols)
	}

	return board
}
