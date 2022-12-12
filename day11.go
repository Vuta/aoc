package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Monkey struct {
	id int
	items []int
	operate func(old int) int
	test func(item int) (int, int)
	divisible int
}

func (m *Monkey) throw(item int, monkeys []*Monkey) {
	val, targetId := m.test(item)
	target := monkeys[targetId]

	target.items = append(target.items, val)
}

func main() {
	input, _ := os.ReadFile("input/day11.txt")
	input = input[:len(input)-1]

	part_1(string(input))
	part_2(string(input))
}

func part_2(input string) {
	text := strings.Split(string(input), "\n\n")
	monkeys := make([]*Monkey, len(text))
	result := make(map[*Monkey]int)
	moduloProduct := 1

	for _, t := range text {
		monkey := parse(t)
		monkeys[monkey.id] = monkey
		moduloProduct *= monkey.divisible
		result[monkey] = 0
	}

	// part 2
	for j := 0; j < 10000; j++ {
		for _, monkey := range monkeys {
			len := len(monkey.items)
			for i := 0; i < len; i++ {
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				newItem := monkey.operate(item) % moduloProduct
				monkey.throw(newItem, monkeys)
				result[monkey]++
			}
		}
	}

	// rely on stdio to perform the multiply because I'm just too lazy hehe
	fmt.Println(result)
}

func part_1(input string) {
	text := strings.Split(string(input), "\n\n")
	monkeys := make([]*Monkey, len(text))
	result := make(map[*Monkey]int)

	for _, t := range text {
		monkey := parse(t)
		monkeys[monkey.id] = monkey
		result[monkey] = 0
	}

	// part 1
	for j := 0; j < 20; j++ {
		for _, monkey := range monkeys {
			len := len(monkey.items)
			for i := 0; i < len; i++ {
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				newItem := monkey.operate(item) / 3
				monkey.throw(newItem, monkeys)
				result[monkey]++
			}
		}
	}

	// rely on stdio to perform the multiply because I'm just too lazy hehe
	fmt.Println(result)
}

func parse(input string) *Monkey {
	attrs := strings.Split(input, "\n")
	id := parseId(attrs[0])
	items := parseItems(attrs[1])
	operationFunc := parseOperation(attrs[2])
	test, divisible := parseTest(attrs[3:])

	monkey := &Monkey{id: id, items: items, operate: operationFunc, test: test, divisible: divisible}

	return monkey
}

func parseId(input string) int {
	idRow := strings.Split(input, " ")
	idCell := idRow[len(idRow)-1]
	id, _ := strconv.Atoi(idCell[:len(idCell)-1])

	return id
}

func parseItems(input string) []int {
	itemsRow := strings.Split(input, "Starting items: ")[1]
	items := []int{}
	for _, item := range strings.Split(itemsRow, ", ") {
		val, _ := strconv.Atoi(item)
		items = append(items, val)
	}

	return items
}

func parseOperation(input string) func(int) int {
	operationRow := strings.Split(
		strings.Split(input, "Operation: ")[1],
		" = ",
	)
	operation := strings.Split(operationRow[1], " ")
	return ops(operation[1], operation[2])
}

func ops(operator string, operand string) func(int) int {
	val, err := strconv.Atoi(operand)
	fun := func(x int) int { return 0 }

	switch operator {
	case "+":
		fun = func(old int) int { return old + val }
	case "*":
		if err != nil {
			fun = func(old int) int { return old * old }
		} else {
			fun = func(old int) int { return old * val }
		}
	}

	return fun
}

func parseTest(input []string) (func(item int) (int, int), int) {
	conditionRow := strings.Split(strings.Split(input[0], "Test: ")[1], " by ")
	fun := func(item int) (int, int) { return -1, -1 }
	operand, _ := strconv.Atoi(conditionRow[1])

	switch conditionRow[0] {
	case "divisible":
		fun = func(item int) (int, int) {
			if item % operand == 0 {
				id, _ := strconv.Atoi(strings.Split(input[1], "If true: throw to monkey ")[1])

				return item, id
			} else {
				id, _ := strconv.Atoi(strings.Split(input[2], "If false: throw to monkey ")[1])

				return item, id
			}
		}
	}

	return fun, operand
}
