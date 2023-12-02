package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	input, _ := os.ReadFile("input/day3.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(string(input)))
	fmt.Println(part_2(string(input)))
}

func part_2(input string) int {
	sum := 0
	data := strings.Split(input, "\n")

	for i := 0; i < len(data); i += 3 {
		badge := find2(data[i], data[i+1], data[i+2])
		sum += value(badge)
	}

	return sum
}

func part_1(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		leng := len(line)/2
		c1 := line[:leng]
		c2 := line[leng:]

		commonItem := find(c1, c2)
		sum += value(commonItem)
	}

	return sum
}

func find(c1, c2 string) byte {
	var result byte

	mapping := make(map[rune]bool)	
	for _, c := range c1 {
		mapping[c] = true
	}

	for _, c := range c2 {
		if mapping[c] {
			result = byte(c)
			break
		}
	}

	return result
}

func find2(b1, b2, b3 string) byte {
	var result byte

	mapping := make(map[rune]int)	
	for _, c := range b1 {
		mapping[c] = 1
	}

	for _, c := range b2 {
		_, ok := mapping[c]
		if ok {
			mapping[c] = 2
		}
	}

	for _, c := range b3 {
		if mapping[c] == 2 {
			result = byte(c)
			break
		}
	}

	return result
}

func value(c byte) int {
	var result int

	if c <= 'z' && c >= 'a' {
		result = int(c - 96)
	} else {
		result = int(c - 38)
	}

	return result
}
