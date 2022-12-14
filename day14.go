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
	input, _ := os.ReadFile("temp.txt")
	// input, _ := os.ReadFile("input/day14.txt")
	input = input[:len(input)-1]

	startCol, endCol, endRow := -1, -1, -1
	data := strings.Split(string(input), "\n")
	rocks := [][]Point{}
	for _, line := range data {
		row := strings.Split(line, " -> ")
		path := []Point{}
		for i := 0; i < len(row); i++ {
			temp := strings.Split(row[i], ",")
			x, _ := strconv.Atoi(temp[0])
		 	y, _ := strconv.Atoi(temp[1])
			path = append(path, Point{x, y})
			if startCol > x || startCol == -1 {
				startCol = x
			}

			if endCol < x {
				endCol = x
			}

			if endRow < y {
				endRow = y
			}

			if i != len(row)-1 {
				next := strings.Split(row[i+1], ",")
				nx, _ := strconv.Atoi(next[0])
				ny, _ := strconv.Atoi(next[1])

				if x != nx && y == ny {
					if x < nx {
						for j := 1; j < nx-x; j++ {
							path = append(path, Point{x+j, y})
						}
					} else {
						for j := 1; j < x-nx; j++ {
							path = append(path, Point{x-j, y})
						}
					}
				} else if x == nx && y != ny {
					if y < ny {
						for j := 1; j < ny-y; j++ {
							path = append(path, Point{x, y+j})
						}
					} else {
						for j := 1; j < y-ny; j++ {
							path = append(path, Point{x, y-j})
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
			grid[p.y][p.x-startCol] = "#"
		}
	}

	grid[0][500-startCol] = "+"

	for _, a := range grid {
		fmt.Println(a)
	}
}
