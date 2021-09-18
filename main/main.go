package main

import (
	"os"
	"raid2"
	"fmt"
)

func main() {
	arguments := os.Args
	if raid2.CheckValid(os.Args[1]]) == nil {
		fmt.Printf("Board not valid try diffrent argument")
	} else {
		// here is 2d board passed as  (board) == ([9][9]int)
		// here should be solving algorithm
		// and printing of grid
	}

	raid2.CheckPossible()

	// tempGrid := [
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// 	[1 2 3 4 5 6 7 8 9]
	// ]int

	raid2.DrawGrid()
}
