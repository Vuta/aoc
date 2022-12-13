package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"sort"
)

type Pair struct {
	left, right any
}

type Pairs []Pair

func (pairs Pairs) Len() int {
	return len(pairs)
}

func (pairs Pairs) Swap(i, j int) {
	pairs[i], pairs[j] = pairs[j], pairs[i]
}

func (pairs Pairs) Less(i, j int) bool {
	return compare(pairs[i].left, pairs[j].left) == -1
}

func main() {
	input, _ := os.ReadFile("input/day13.txt")
	input = input[:len(input)-1]

	fmt.Println(part_1(string(input)))
	fmt.Println(part_2(string(input)))
}

func part_1(input string) int {
	pairs := []Pair{}
	data := strings.Split(input, "\n\n")
	for _, text := range data {
		lines := strings.Split(text, "\n")

		p := Pair{}
		json.Unmarshal([]byte(lines[0]), &(p.left))
		json.Unmarshal([]byte(lines[1]), &(p.right))

		pairs = append(pairs, p)
	}

	sum := 0
	for id, pair := range pairs {
		if compare(pair.left, pair.right) == -1 {
			sum += id+1
		}
	}

	return sum
}

func part_2(input string) int {
	pairs := Pairs{}
	data := strings.Split(input, "\n")
	i := 1
	for _, text := range data {
		if text == "" {
			continue
		}

		p := Pair{}
		i++
		json.Unmarshal([]byte(text), &(p.left))

		pairs = append(pairs, p)
	}

	var divider1, divider2 any
	json.Unmarshal([]byte(`[[2]]`), &divider1)
	json.Unmarshal([]byte(`[[6]]`), &divider2)

	pairs = append(pairs, Pair{left: divider1})
	pairs = append(pairs, Pair{left: divider2})

	sort.Sort(pairs)

	a, b := 0, 0
	for i, p := range pairs {
		x, _ := json.Marshal(p.left)

		if string(x) == `[[2]]` {
			a = i+1
		}

		if string(x) == `[[6]]` {
			b = i+1
		}
	}

	return a * b
}

func compare(l, r any) int {
	lVal, ok1 := l.(float64)
	rVal, ok2 := r.(float64)

	if ok1 && ok2 {
		if lVal == rVal {
			return 0 // next
		}

		if lVal < rVal {
			return -1 // true
		}

		return 1 // false
	}

	if ok1 {
		return compare([]interface{}{lVal}, r)
	}

	if ok2 {
		return compare(l, []interface{}{rVal})
	}

	lList := l.([]interface{})
	rList := r.([]interface{})
	if len(lList) == 0 && len(rList) == 0 {
		return 0
	}

	if len(lList) == 0 {
		return -1
	}

	if len(rList) == 0 {
		return 1
	}

	if lList[0] == nil && rList[0] != nil {
		return -1
	}

	if lList[0] != nil && rList[0] == nil {
		return 1
	}

	flag := compare(lList[0], rList[0])
	if flag == 0 {
		return compare(lList[1:], rList[1:])
	} else {
		return flag
	}

	return -2
}
