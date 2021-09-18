package main

import (
	"os"
	"raid2"
)

func main() {
	arguments := os.Args

	raid2.CheckValid(arguments[1:])

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


//"raidd2"
//"os"
//"fmt"
//)
//
//func main() {
//board := parseInput(os.Args[1])
//
//printBoard(board)
//
//if backtrack(&board) {
//	fmt.Println("The Sudoku was solved successfully:")
//	printBoard(board)
//} else {
//	fmt.Printf("The Sudoku can't be solved.")
//}
//}