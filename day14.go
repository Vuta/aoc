package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	// input, _ := os.ReadFile("temp.txt")
	input, _ := os.ReadFile("input/day14.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(string(input)))
}

func part_1(input string) int {
	grid, startPoint := buildGrid(input)
	found := false
	for !found {
		grid, found = move(startPoint, grid)
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "o" {
				count++
			}
		}
	}

	return count
}

func move(s Point, grid [][]string) ([][]string, bool) {
	current := s
	for grid[current.x][current.y] != "o" {
		if target(current, grid) {
			fmt.Println("FOUND IT:", current)
			return grid, true
		}

		if canMoveDown(current, grid) {
			current = Point{current.x+1, current.y}
			continue
		}

		if canMoveLeft(current, grid) {
			current = Point{current.x+1, current.y-1}
			continue
		}

		if canMoveRight(current, grid) {
			current = Point{current.x+1, current.y+1}
			continue
		}

		grid[current.x][current.y] = "o"
	}

	return grid, false
}

func target(p Point, grid [][]string) bool {
	return p.x+1 == len(grid)-1 && canMoveLeft(p, grid)
}

func canMoveDown(p Point, grid [][]string) bool {
	down := Point{p.x+1, p.y}

	return (down.x < len(grid) && grid[down.x][down.y] == ".")
}

func canMoveLeft(p Point, grid [][]string) bool {
	down := Point{p.x+1, p.y-1}

	return (down.x < len(grid) && down.y >= 0 && grid[down.x][down.y] == ".") || down.y < 0
}

func canMoveRight(p Point, grid [][]string) bool {
	down := Point{p.x+1, p.y+1}

	return (down.x < len(grid) && down.y < len(grid[0]) && grid[down.x][down.y] == ".") || down.y > len(grid[0])-1
}

func isBlocked(p Point, grid [][]string) bool {
	return !canMoveDown(p, grid) && !canMoveLeft(p, grid) && !canMoveRight(p, grid)
}

func buildGrid(input string) ([][]string, Point) {
	startCol, endCol, endRow := -1, -1, -1
	data := strings.Split(input, "\n")
	for _, line := range data {
		row := strings.Split(line, " -> ")
		for i := 0; i < len(row); i++ {
			temp := strings.Split(row[i], ",")
			c, _ := strconv.Atoi(temp[0])
		 	r, _ := strconv.Atoi(temp[1])

			if startCol > c || startCol == -1 {
				startCol = c
			}

			if endCol < c {
				endCol = c
			}

			if endRow < r {
				endRow = r
			}
		}
	}

	rocks := [][]Point{}
	for _, line := range data {
		row := strings.Split(line, " -> ")
		path := []Point{}
		for i := 0; i < len(row); i++ {
			temp := strings.Split(row[i], ",")
			c, _ := strconv.Atoi(temp[0])
		 	r, _ := strconv.Atoi(temp[1])
			point := Point{r, c-startCol}
			path = append(path, point)

			if i != len(row)-1 {
				next := strings.Split(row[i+1], ",")
				nc, _ := strconv.Atoi(next[0])
				nr, _ := strconv.Atoi(next[1])

				nc = nc-startCol

				if point.y != nc && point.x == nr {
					if point.y < nc {
						for j := 1; j < nc-point.y; j++ {
							path = append(path, Point{point.x, point.y+j})
						}
					} else {
						for j := 1; j < point.y-nc; j++ {
							path = append(path, Point{point.x, point.y-j})
						}
					}
				} else if point.y == nc && point.x != nr {
					if point.x < nr {
						for j := 1; j < nr-point.x; j++ {
							path = append(path, Point{point.x+j, point.y})
						}
					} else {
						for j := 1; j < point.x-nr; j++ {
							path = append(path, Point{point.x-j, point.y})
						}
					}
				} else {
					panic("invalid path")
				}
			}
		}

		rocks = append(rocks, path)
	}

	grid := [][]string{}
	for i := 0; i < endRow+1; i++ {
		row := []string{}
		for j := 0; j < endCol-startCol+1; j++ {
			row = append(row, ".")
		}

		grid = append(grid, row)
	}

	for _, path := range rocks {
		for _, p := range path {
			grid[p.x][p.y] = "#"
		}
	}

	startPoint := Point{0, 500-startCol}
	grid[startPoint.x][startPoint.y] = "+"

	return grid, startPoint
}
