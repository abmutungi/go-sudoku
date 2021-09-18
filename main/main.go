package main

import (
	"fmt"
	"os"
	"raid2"
)

func main() {
	errorMessage := "Error"

	arguments := os.Args

	sudokuInput := arguments[1:]

	validInput, grid := raid2.NiksCheckValid(sudokuInput)

	if !validInput {
		fmt.Println(errorMessage)
	}

	raid2.DrawGrid(grid)
}
