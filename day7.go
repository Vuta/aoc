package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
)

type Tree struct {
	name string
	parent *Tree
	size int
}

func main() {
	input, _ := os.ReadFile("input/day7.txt")
	input = input[:len(input)-1]
	history := strings.Split(string(input), "\n")

	// full path -> directory
	dirMap := make(map[string]*Tree)
	rootDir := &Tree{name: "/"}
	currentDir := rootDir
	dirMap["/"] = rootDir

	for i := 1; i < len(history); i++ {
		cmd := history[i]

		if isCd(cmd) {
			currentDir = handleCd(cmd, currentDir, dirMap)
		} else if isLs(cmd) {
			j := i+1
			next := history[j]
			for !isLs(next) && !isCd(next) {
				handleLs(next, currentDir, dirMap)

				if j == len(history)-1 {
					break
				}

				j++
				next = history[j]
			}

			i = j-1
		}
	}

	count := 0
	available := 70000000 - dirMap["/"].size
	need := 30000000
	min := 70000000
	for _, dir := range dirMap {
		if dir.size <= 100000 {
			count += dir.size
		}

		if dir.size + available >= need && dir.size < min {
			min = dir.size
		}
	}

	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", min)
}

func isCd(cmd string) bool {
	return cmd[:4] == "$ cd"
}

func isLs(cmd string) bool {
	return cmd[:4] == "$ ls"
}

func isDir(cmd string) bool {
	return cmd[:3] == "dir"
}

func parseFile(file string) int {
	rex := regexp.MustCompile(`(\d+) .+`)
	size := rex.FindStringSubmatch(file)[1]
	val, _ := strconv.Atoi(size)

	return val
}

func getFullPath(currentDir *Tree) string {
	name := currentDir.name
	temp := currentDir
	for temp.parent != nil {
		name = temp.parent.name + name
		temp = temp.parent
	}

	return name
}

func handleCd(cmd string, currentDir  *Tree, dirMap map[string]*Tree) *Tree {
	rex := regexp.MustCompile(`\$ cd (.+)`)
	dir := rex.FindStringSubmatch(cmd)[1]
	if dir == ".." {
		if currentDir != dirMap["/"] {
			currentDir = currentDir.parent
		}
	} else {
		fullPath := getFullPath(currentDir)+dir
		item, ok := dirMap[fullPath]
		if ok {
			currentDir = item
		} else {
			currentDir = &Tree{name: dir, parent: currentDir}
			dirMap[fullPath] = currentDir
		}
	}

	return currentDir
}

func handleLs(cmd string, currentDir *Tree, dirMap map[string]*Tree) {
	if isDir(cmd) {
		handleLsDir(cmd, currentDir, dirMap)
	} else {
		handleLsFile(cmd, currentDir)
	}
}

func handleLsFile(cmd string, currentDir *Tree) {
	nextSize := parseFile(cmd)
	currentDir.size += nextSize
	temp := currentDir
	for temp.parent != nil {
		temp.parent.size += nextSize
		temp = temp.parent
	}
}

func handleLsDir(next string, currentDir *Tree, dirMap map[string]*Tree) {
	rex := regexp.MustCompile(`dir (.+)`)
	dir := rex.FindStringSubmatch(next)[1]
	fullPath := getFullPath(currentDir)+dir
	_, ok := dirMap[fullPath]
	if !ok {
		dirMap[fullPath] = &Tree{name: dir, parent: currentDir}
	}
}
