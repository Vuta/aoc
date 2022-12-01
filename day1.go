package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

type Set struct {
	list []int
	maxLen int
}

// sorted in desc order
// only maintain the top maxLen elements
func (s *Set) insert(element int) {
	index := len(s.list)

	for i := 0; i < len(s.list); i++ {
		if s.list[i] <= element {
			index = i
			break
		}
	}

	if index == len(s.list) {
		s.list = append(s.list, element)
	} else {
		s.list = append(s.list[:index+1], s.list[index:]...)
		s.list[index] = element
	}

	if len(s.list) > s.maxLen {
		s.list = s.list[:s.maxLen]
	}
}

func (s *Set) total() int {
	result := 0
	for _, val := range s.list {
		result += val
	}

	return result
}

func main() {
	input, _ := os.ReadFile("input/day1.txt")

	fmt.Println(part_1(string(input)))
	fmt.Println(part_2(string(input)))
}

func part_1(input string) int {
	max := 0
	for _, item := range strings.Split(input, "\n\n") {
		list := strings.Split(item, "\n")
		count := 0

		for _, data := range list {
			val, _ := strconv.Atoi(data)
			count += val
		}

		if max < count {
			max = count
		}
	}

	return max
}

func part_2(input string) int {
	result := &Set{maxLen: 3}
	for _, item := range strings.Split(input, "\n\n") {
		list := strings.Split(item, "\n")
		count := 0

		for _, data := range list {
			val, _ := strconv.Atoi(data)
			count += val
		}

		result.insert(count)
	}

	return result.total()
}
