package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input/day6.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(input))
	fmt.Println(part_2(input))
}

func part_1(input []byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if isMarker(input[i:i+4]) {
			count += 4
			break
		} else {
			count++
		}
	}

	return count
}

func part_2(input []byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if isMarker(input[i:i+14]) {
			count += 14
			break
		} else {
			count++
		}
	}

	return count
}

func isMarker(input []byte) bool {
	mapping := make(map[byte]bool)
	for _, c := range input {
		mapping[c] = true
	}

	return len(mapping) == len(input)
}
