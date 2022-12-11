package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Command struct {
	ins string
	val int
	cycle int
}

func main() {
	input, _ := os.ReadFile("input/day10.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(string(input)))
}

func part_1(input string) int {
	list := strings.Split(string(input), "\n")
	cycleCount := 0
	data := []int{20,60,100,140,180,220}
	output := []int{1,1,1,1,1,1}
	sum := 0

	for i := 0; i < len(list); i++ {
		cmd := parse(list[i])
		cycleCount += cmd.cycle

		for i, e := range data {
			if cycleCount < e {
				output[i] += cmd.val
			}
		}
	}

	for i, e := range output {
		sum += data[i]*e 
	}

	return sum
}

func parse(instruction string) Command {
	data := strings.Split(instruction, " ")
	val := 0
	cycle := 0

	switch data[0] {
	case "addx":
		cycle = 2
		val, _ = strconv.Atoi(data[1])
	case "noop":
		cycle = 1
	}

	return Command{ins: data[0], val: val, cycle: cycle}
}
