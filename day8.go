package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input/day8.txt")
	input = input[:len(input)-1]

	grid := [][]int{}

	for i := 0; i < len(input)-1; i++ {
		row := []int{}
		for i < len(input) && input[i] != '\n' {
			row = append(row, int(input[i]-48))
			i++
		}
		grid = append(grid, row)
	}

	fmt.Println("Part 1", part_1(grid))
	fmt.Println("Part 2", part_2(grid))
}

func part_1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			current := grid[i][j]
			if look(current, i, j, grid) {
				count++
			}
		}
	}

	return count
}

func look(current, i, j int, grid [][]int) bool {
	return dfs(current, i, j, grid, "up") ||
		dfs(current, i, j, grid, "down") ||
		dfs(current, i, j, grid, "left") ||
		dfs(current, i, j, grid, "right")
}

func dfs(current, i, j int, grid [][]int, direction string) bool {
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
		return true
	}

	switch direction {
	case "up":
		if current > grid[i-1][j] {
			return dfs(current, i-1, j, grid, direction)
		}
	case "down":
		if current > grid[i+1][j] {
			return dfs(current, i+1, j, grid, direction)
		}
	case "left":
		if current > grid[i][j-1] {
			return dfs(current, i, j-1, grid, direction)
		}
	case "right":
		if current > grid[i][j+1] {
			return dfs(current, i, j+1, grid, direction)
		}
	}

	return false
}

func part_2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	max := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			current := grid[i][j]
			score := dfs2(current, i, j, grid, "up", 0) * 
				dfs2(current, i, j, grid, "down", 0) *
				dfs2(current, i, j, grid, "left", 0) *
				dfs2(current, i, j, grid, "right", 0)

			if max < score {
				max = score
			}
		}
	}

	return max
}

func dfs2(current, i, j int, grid [][]int, direction string, result int) int {
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
		return result
	}

	switch direction {
	case "up":
		if current > grid[i-1][j] {
			return dfs2(current, i-1, j, grid, direction, result+1)
		}
	case "down":
		if current > grid[i+1][j] {
			return dfs2(current, i+1, j, grid, direction, result+1)
		}
	case "left":
		if current > grid[i][j-1] {
			return dfs2(current, i, j-1, grid, direction, result+1)
		}
	case "right":
		if current > grid[i][j+1] {
			return dfs2(current, i, j+1, grid, direction, result+1)
		}
	}

	return result+1
}
