package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input/day5.txt")
	input = input[:len(input)-1]

	data := strings.Split(string(input), "\n\n")
	config_input := strings.Split(data[0], "\n")
	config := [][]string{}
	for range config_input {
		config = append(config, []string{})
	}

	for i := len(config_input)-2; i >= 0; i-- {
		for j, c := range config_input[i] {
			if c == ' ' || c == '[' || c == ']' {
				continue
			}

			config[(j/4)] = append(config[(j/4)], string(c))
		}
	}
	dupConfig := [][]string{}
	for i, st := range config {
		dupConfig = append(dupConfig, []string{})
		dupConfig[i] = append(dupConfig[i], st...)
	}
	instruction := strings.Split(data[1], "\n")

	part_1(config, instruction)
	part_2(dupConfig, instruction)
}

func part_1(config [][]string, instruction []string) {
	for _, ins := range instruction {
		quantity, fromId, toId := parseIns(ins)
		for i := 1; i <= quantity; i++ {
			item := config[fromId][len(config[fromId])-1]
			config[fromId] = config[fromId][:len(config[fromId])-1]
			config[toId] = append(config[toId], item)
		}
	}

	for _, st := range config {
		fmt.Print(st[len(st)-1])
	}
	fmt.Println("")
}

func part_2(config [][]string, instruction []string) {
	for _, ins := range instruction {
		quantity, fromId, toId := parseIns(ins)
		items := config[fromId][len(config[fromId])-quantity:]
		config[fromId] = config[fromId][:len(config[fromId])-quantity]
		config[toId] = append(config[toId], items...)
	}

	for _, st := range config {
		fmt.Print(st[len(st)-1])
	}
	fmt.Println("")
}

func parseIns(ins string) (int, int, int) {
	reg := regexp.MustCompile(`\d+`)
	match := reg.FindAllString(ins, -1)

	quantity, _ := strconv.Atoi(match[0])
	from, _ := strconv.Atoi(match[1])
	to, _ := strconv.Atoi(match[2])

	return quantity, from-1, to-1
}
