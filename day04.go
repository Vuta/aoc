package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
	input, _ := os.ReadFile("input/day4.txt")
	input = input[:len(input)-1]

	fmt.Println(part(string(input)))
}

func part(input string) (int, int) {
	containsCount := 0
	overlapCount := 0
	data := strings.Split(input, "\n")

	for _, pair := range data {
		val := strings.Split(pair, ",")
		p1, p2 := val[0], val[1]

		a1, a2 := strings.Split(p1, "-")[0], strings.Split(p1, "-")[1]
		lower1, _ := strconv.Atoi(a1)
		upper1, _ := strconv.Atoi(a2)

		b1, b2 := strings.Split(p2, "-")[0], strings.Split(p2, "-")[1]
		lower2, _ := strconv.Atoi(b1)
		upper2, _ := strconv.Atoi(b2)

		if contains(lower1, upper1, lower2, upper2) {
			containsCount++
		}
		if overlap(lower1, upper1, lower2, upper2) {
			overlapCount++
		}
	}

	return containsCount, overlapCount
}

func contains(a1, b1, c1, d1 int) bool {
	if (a1 >= c1 && b1 <= d1) || (a1 <= c1 && d1 <= b1) {
		return true
	}

	return false
}

func overlap(a1, b1, c1, d1 int) bool {
	if (a1 <= c1 && c1 <= b1) || (c1 <= a1 && a1 <= d1) {
		return true
	}

	return false
}
