package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
)

type Point struct {
	row, col int
}

func main() {
	input, _ := os.ReadFile("input/day15.txt")
	input = input[:len(input)-1]

	part_1(string(input))
}

func part_1(input string) {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	target := 2000000
	toLoop := make(map[Point]Point)
	names := make(map[Point]string)
	for _, line := range strings.Split(input, "\n") {
		// y is row, x is col duh
		sx, sy, bx, by := parse(line, re)
		b := Point{by, bx}
		s := Point{sy, sx}
		names[b] = "B"
		names[s] = "S"

		d := abs(s.row - b.row) + abs(s.col - b.col)
		if s.row - d <= target && target <= s.row + d {
			toLoop[s] = b
		}
	}

	mapping := make(map[Point]bool)
	for sensor, beacon := range toLoop {
	 	d := abs(sensor.row - beacon.row) + abs(sensor.col - beacon.col)
		height := abs(target - sensor.row)
		width := d - height
		for j := sensor.col-width; j <= sensor.col+width; j++ {
			p := Point{target, j}
			if !mapping[p] && names[p] != "S" && names[p] != "B" {
				mapping[p] = true
			}
		}
	}


	fmt.Println(len(mapping))
}

func parse(line string, re *regexp.Regexp) (int, int, int, int) {
	data := re.FindAllStringSubmatch(line, -1)[0]

	a, _ := strconv.Atoi(data[1])
	b, _ := strconv.Atoi(data[2])
	c, _ := strconv.Atoi(data[3])
	d, _ := strconv.Atoi(data[4])

	return a, b, c, d
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
