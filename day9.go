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
	input, _ := os.ReadFile("input/day9.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}

func part_1(input []byte) int {
	position := []*Point{&Point{0,0}, &Point{0,0}}
	mapping := make(map[Point]bool)

	for _, move := range strings.Split(string(input), "\n") {
		position = run(move, position, mapping)
	}

	return len(mapping)
}

func part_2(input []byte) int {
	position := []*Point{}
	for i := 0; i < 10; i++ {
		position = append(position, &Point{0,0})
	}
	mapping := make(map[Point]bool)

	for _, move := range strings.Split(string(input), "\n") {
		position = run(move, position, mapping)
	}

	return len(mapping)
}

func parse(move string) (byte, int) {
	direction := move[0]
	num, _ := strconv.Atoi(string(move[2:]))

	return direction, num
}

func run(move string, position []*Point, mapping map[Point]bool) []*Point {
	direction, num := parse(move)
	for i := 0; i < num; i++ {
		head := position[0]
		switch direction {
		case 'R': head.y++
		case 'L': head.y--
		case 'U': head.x++
		case 'D': head.x--
		}

		for j := 1; j < len(position); j++ {
			prev := position[j-1]
			current := position[j]

			currentTail := &Point{x: current.x, y: current.y}
			currentHead := &Point{x: prev.x, y: prev.y}

			position[j] = rePosition(currentHead, currentTail)
		}

		newTail := position[len(position)-1]
		mapping[*newTail] = true
	}

	return position
}

func rePosition(hd, tl *Point) *Point {
	r := tl.x - hd.x
	c := tl.y - hd.y

	if (r == 2 || r == -2) && (c == 2 || c == -2) {
		if tl.x < hd.x {
			tl.x++
		} else {
			tl.x--
		}

		if tl.y < hd.y {
			tl.y++
		} else {
			tl.y--
		}
	} else if (r == 2 || r == -2) {
		tl.y = hd.y
		if tl.x < hd.x {
			tl.x++
		} else {
			tl.x--
		}
	} else if (c == 2 || c == -2) {
		tl.x = hd.x
		if tl.y < hd.y {
			tl.y++
		} else {
			tl.y--
		}
	}

	return tl
}
