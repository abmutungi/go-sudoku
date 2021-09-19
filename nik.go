package raid2

import (
	"fmt"
)

func NiksCheckValid(input []string) (bool, [9][9]int) {
	explain := false

	if explain {
		fmt.Println(input)
	}

	// init 2d grid
	var grid [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid[i][j] = 0
		}
	}

	// fmt.Println(grid)

	var tempRow []int
	var tempRows []int

	// first we need to extract the digits
	for _, row := range input {

		// check if row length is 9 digits
		if len(row) != 9 {
			if explain {
				fmt.Println("More or fewer than 9 digits in one of the rows")
			}
			return false, grid
		}

		// iterating row
		for _, char := range row {
			// fmt.Println(string(char), int(char-48))
			// fmt.Printf("%T, %T", string(char), int(char-48))

			// non digit or '.'
			if char < 48 && char > 57 && char != '.' {
				return false, grid
			} else {
				if char == '.' {
					tempRow = append(tempRow, 0)
				} else {
					tempRow = append(tempRow, int(char-48)) // ascii 1
				}
			}
			// end of characters in row
		}
		if explain {
			fmt.Println("made temp row", tempRow)
		}
		// check if digits inside the input row are not repeating

		if ArnoldsCheckNumbersRepeated(tempRow) {
			if explain {
				fmt.Println("numbers repeated in row.. exiting")
			}
			return false, grid
		}

		// add the current row into a temporary slice of int called tempRows
		tempRows = append(tempRows, tempRow...)

		// reset the slice of temporary numbers
		tempRow = nil
		// end of row
	}

	// add the values of tempRows into the grid

	countrow := 0
	countcol := 0

	for i := 0; i < len(tempRows); i++ {

		if countrow == 9 {
			countrow = 0
			countcol++
			if explain {
				fmt.Println()
			}
		}
		// this was so tricky!!...
		grid[countcol][countrow] = tempRows[i]
		if explain {
			fmt.Print(tempRows[i], " ")
		}
		countrow++
	}

	if explain {
		fmt.Println()
		fmt.Println("input valid")
	}
	return true, grid
}

func CheckPossible(grid *[9][9]int, num, y, x int) bool {
	// check column
	for i := 0; i < 9; i++ {
		if grid[y][i] == num {
			return false
		}
	}
	// check row
	for i := 0; i < 9; i++ {
		if grid[i][x] == num {
			return false
		}
	}

	// check mini square
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[y0+i][x0+j] == num {
				return false
			}
		}
	}

	return true
}
