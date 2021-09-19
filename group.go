package raid2

func SolveSudoku(grid *[9][9]int, solutions *int) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == 0 {
				for value := 1; value <= 9; value++ {
					if CheckPossible(grid, value, y, x) {
						grid[y][x] = value
						SolveSudoku(grid,solutions)
						grid[y][x] = 0
					}
				}
				return
			}
		}
	}
	if *solutions<2{
		DrawGrid(grid)
	}
	*solutions ++
	return
}
