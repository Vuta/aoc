package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	input, _ := os.ReadFile("input/day12.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}

func part_2(input []byte) int {
	grid, starts, end := buildGrid2(input)	
	min := -1 

	for _, start := range starts {
		count := bfs(grid, start, end)

		if (count < min || min == -1) && count != -1 {
			min = count
		}
	}

	return min
}

func part_1(input []byte) int {
	grid, start, end := buildGrid(input)	
	return bfs(grid, start, end)
}

func bfs(grid [][]byte, start, end Point) int {
	visited := make(map[Point]bool)
	visited[start] = true

	queue := []Point{start}
	count := 0

	for len(queue) > 0 {
		leng := len(queue)
		for i := 0; i < leng; i++ {
			p := queue[0]
			queue = queue[1:]

			if p == end {
				return count
			}

			for _, next := range children(p, grid) {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		count++
	}

	// when there is no path
	return -1
}

func buildGrid(input []byte) ([][]byte, Point, Point) {
	grid := [][]byte{}
	start, end := Point{}, Point{}
	for i := 0; i < len(input)-1; i++ {
		row := []byte{}

		for i < len(input) && input[i] != '\n' {
			row = append(row, input[i])

			if input[i] == 'S' {
				start = Point{x: len(grid), y: (len(row) - 1) % len(input)}
			}

			if input[i] == 'E' {
				end = Point{x: len(grid), y: (len(row) - 1) % len(input)}
			}

			i++
		}

		grid = append(grid, row)
	}
	grid[end.x][end.y] = 'z'
	grid[start.x][start.y] = 'a'

	return grid, start, end
}

func buildGrid2(input []byte) ([][]byte, []Point, Point) {
	grid := [][]byte{}
	start := []Point{}
	end := Point{}
	for i := 0; i < len(input)-1; i++ {
		row := []byte{}

		for i < len(input) && input[i] != '\n' {
			row = append(row, input[i])

			if input[i] == 'S' || input[i] == 'a' {
				node := Point{x: len(grid), y: (len(row) - 1) % len(input)}
				start = append(start, node)
			}

			if input[i] == 'E' {
				end = Point{x: len(grid), y: (len(row) - 1) % len(input)}
			}

			i++
		}

		grid = append(grid, row)
	}

	for _, node := range start {
		grid[node.x][node.y] = 'a'
	}
	grid[end.x][end.y] = 'z'

	return grid, start, end
}

func children(p Point, grid [][]byte) []Point {
	result := []Point{}

	if p.x < len(grid)-1 && canStep(grid[p.x+1][p.y], grid[p.x][p.y]){
		result = append(result, Point{x: p.x+1, y: p.y})
	}

	if p.x > 0 && canStep(grid[p.x-1][p.y], grid[p.x][p.y]) {
		result = append(result, Point{x: p.x-1, y: p.y})
	}

	if p.y < len(grid[0])-1 && canStep(grid[p.x][p.y+1], grid[p.x][p.y]) {
		result = append(result, Point{x: p.x, y: p.y+1})
	}

	if p.y > 0 && canStep(grid[p.x][p.y-1], grid[p.x][p.y]) {
		result = append(result, Point{x: p.x, y: p.y-1})
	}

	return result
}

func canStep(next, current byte) bool {
	return rune(next) - rune(current) <= 1
}
