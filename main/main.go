package main

import (
	"os"
	"raid2"
)

func main() {
	arguments := os.Args

	raid2.CheckValid(os.Args[1]])

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
