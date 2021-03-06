package main

import (
	"fmt"
	"os"
	"raid2"
)

func main() {
	errorMessage := "Error"

	noOfSolutions := 1

	arguments := os.Args

	sudokuInput := arguments[1:]

	validInput, grid := raid2.NiksCheckValid(sudokuInput)

	if !validInput || len(arguments)==1{
		fmt.Printf("%s\n", errorMessage)
	} else {
		raid2.SolveSudoku(&grid, &noOfSolutions)
	}
}
