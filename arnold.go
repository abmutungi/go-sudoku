package raid2

import "fmt"

func DrawGrid(grid [][]int) {
	/* grid := [9][9]int{ // a slice of a slice with 9 rows and 9 columns
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	*/
	for i := 0; i < 9; i++ { // declare i and j as integers
		for j := 0; j < 9; j++ {
			fmt.Print(grid[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
