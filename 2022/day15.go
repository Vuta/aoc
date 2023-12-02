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

type Area struct {
	top, left, right, bottom, center Point
}

func main() {
	input, _ := os.ReadFile("input/day15.txt")
	input = input[:len(input)-1]

	part_1(string(input))
	part_2(string(input))
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

// loop through the input to build a list of areas, where each area is formed by: for each sensor, find its four edges (top, bottom, right, left )
// O(n) where n is the number of sensor (38)
// for each point in (0,0)..(4_000_000,4_000_000), check if this point is not in any of area above
// the problem reduces to how to detect if a point (x,y) is in an area, formed by top (tx, ty), left (lx, ly), bottom (bx, by) and right (rx, ry)
// this is the formula H = (Bx - Ax) * (Cy - Ay) - (By - Ay) * (Cx - Ax)
// the rhombus formed by 4 edges: top-left, left-bottom, bottom-right and right-top
// a point is in this rhombus if and only if
// 1) it's on the right of top-left => H >= 0
// 2) it's on the right of left-bottom => H >= 0
// 3) it's on the left of bottom-right => H >= 0
// 4) it's on the left of right-top => H >= 0
// if current point is not in any of these areas, return that point
// if not, jump to the next col by using the current area's right's col
// because all points within that range are covered by the current's area already

func part_2(input string) {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	areas := []Area{}
	for _, line := range strings.Split(input, "\n") {
		sx, sy, bx, by := parse(line, re)
		b := Point{by, bx}
		s := Point{sy, sx}

		d := abs(s.row - b.row) + abs(s.col - b.col)
		top := Point{s.row+d,s.col}
		bottom := Point{s.row-d,s.col}
		left := Point{s.row,s.col-d}
		right := Point{s.row,s.col+d}

		area := Area{top, left, right, bottom, s}
		areas = append(areas, area)
	}

	magic := 4000000
	target := Point{}
	for i := 0; i <= magic; i++ {
		for j := 0; j <= magic; j++ {
			p := Point{i, j}
			flag := false
			for _, area := range areas {
				if inArea(p, area) {
					top := area.top
					right := area.right
					j = (top.col - right.col) * (p.row - right.row) / (top.row - right.row) + right.col

					flag = true
				}
			}

			if flag == false {
				target = p
				fmt.Println("HAHA", target, target.col * 4000000 + target.row)
				break
			}
		}
	}
}

// (Bx - Ax) * (Cy - Ay) - (By - Ay) * (Cx - Ax)
func inArea(p Point, area Area) bool {
	top, left, right, bottom := area.top, area.left, area.right, area.bottom

	// A = top, B = left
	topLeft := (left.col - top.col) * (p.row - top.row) - (left.row - top.row) * (p.col - top.col) >= 0

	// A = left, B = bottom
	leftBottom := (bottom.col - left.col) * (p.row - left.row) - (bottom.row - left.row) * (p.col - left.col) >= 0

	// A = bottom, B = right
	bottomRight := (right.col - bottom.col) * (p.row - bottom.row) - (right.row - bottom.row) * (p.col - bottom.col) >= 0

	// A = right, B = top
	rightTop := (top.col - right.col) * (p.row - right.row) - (top.row - right.row) * (p.col - right.col) >= 0

	return topLeft && leftBottom && bottomRight && rightTop
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
