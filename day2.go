package main

import (
	"fmt"
	"strings"
	"os"
)

type Round struct {
	p1, p2 byte
}

func main() {
	input, _ := os.ReadFile("input/day2.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(string(input)))
	fmt.Println(part_2(string(input)))
}

func part_1(input string) int {
	outcomeMap := map[Round]int{
		Round{p1: 'A', p2: 'X'}: 4,
		Round{p1: 'B', p2: 'X'}: 1,
		Round{p1: 'C', p2: 'X'}: 7,
		Round{p1: 'A', p2: 'Y'}: 8,
		Round{p1: 'B', p2: 'Y'}: 5,
		Round{p1: 'C', p2: 'Y'}: 2,
		Round{p1: 'A', p2: 'Z'}: 3,
		Round{p1: 'B', p2: 'Z'}: 9,
		Round{p1: 'C', p2: 'Z'}: 6,
	}
	score := 0

	for _, pair := range strings.Split(input, "\n") {
		round := Round{p1: pair[0], p2: pair[2]}
		score += outcomeMap[round]
	}

	return score
}

func part_2(input string) int {
	outcomeMap := map[Round]int{
		Round{p1: 'A', p2: 'X'}: 3,
		Round{p1: 'B', p2: 'X'}: 1,
		Round{p1: 'C', p2: 'X'}: 2,
		Round{p1: 'A', p2: 'Y'}: 4,
		Round{p1: 'B', p2: 'Y'}: 5,
		Round{p1: 'C', p2: 'Y'}: 6,
		Round{p1: 'A', p2: 'Z'}: 8,
		Round{p1: 'B', p2: 'Z'}: 9,
		Round{p1: 'C', p2: 'Z'}: 7,
	}
	score := 0

	for _, pair := range strings.Split(input, "\n") {
		round := Round{p1: pair[0], p2: pair[2]}
		score += outcomeMap[round]
	}

	return score
}
