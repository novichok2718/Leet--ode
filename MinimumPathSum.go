package main

func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	
	for j := 1; j < column; j++ {
		grid[0][j] += grid[0][j - 1]
	}
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i - 1][0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
		} 
	}
	return grid[row - 1][column - 1]
}