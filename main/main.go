package main

import (
	"raidd2"
	"os"
	"fmt"
)

func main() {
	board := parseInput(os.Args[1])

	printBoard(board)

	if backtrack(&board) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(board)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
}
}
