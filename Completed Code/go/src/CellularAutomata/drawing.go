package CellularAutomata

import (
	"image"
)

func DrawGameBoards(boards []GameBoard, cellWidth int) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth)
	}
	return imageList
}

func DrawGameBoard(board GameBoard, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	darkGray := MakeColor(50, 50, 50)
	// black := MakeColor(0, 0, 0)
	blue := MakeColor(0, 0, 255)
	red := MakeColor(255, 0, 0)
	green := MakeColor(0, 255, 0)
	yellow := MakeColor(255, 255, 0)
	magenta := MakeColor(255, 0, 255)
	white := MakeColor(255, 255, 255)
	cyan := MakeColor(0, 255, 255)

	/*
		//set the entire board as black
		c.SetFillColor(gray)
		c.ClearRect(0, 0, height, width)
		c.Clear()
	*/

	/*
		// draw the grid lines in white
		c.SetStrokeColor(white)
		DrawGridLines(c, cellWidth)
	*/

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				c.SetFillColor(darkGray)
			} else if board[i][j] == 1 {
				c.SetFillColor(blue)
			} else if board[i][j] == 2 {
				c.SetFillColor(red)
			} else if board[i][j] == 3 {
				c.SetFillColor(green)
			} else if board[i][j] == 4 {
				c.SetFillColor(yellow)
			} else if board[i][j] == 5 {
				c.SetFillColor(magenta)
			} else if board[i][j] == 6 {
				c.SetFillColor(white)
			} else if board[i][j] == 7 {
				c.SetFillColor(cyan)
			} else {
				panic("Error: Out of range value " + string(board[i][j]) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return GetImage(c)
}

func DrawGridLines(pic Canvas, cellWidth int) {
	w, h := pic.Width(), pic.Height()
	// first, draw vertical lines
	for i := 1; i < w/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < h/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}
