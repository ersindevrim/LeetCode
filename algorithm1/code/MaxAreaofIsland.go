package algorithm1

import "github.com/ersindevrim/leetcode/global"

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for row, line := range grid {
		for col, value := range line {
			if value == 0 {
				continue
			}

			area := islandCalculator(grid, row, col)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func islandCalculator(grid [][]int, row, col int) int {
	if global.IsInGrid(grid, row, col) {
		return 0
	}

	grid[row][col] = 0
	return 1 + islandCalculator(grid, row-1, col) + islandCalculator(grid, row+1, col) + islandCalculator(grid, row, col-1) + islandCalculator(grid, row, col+1)
}
