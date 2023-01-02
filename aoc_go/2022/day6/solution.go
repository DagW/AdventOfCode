package main

import (
	"fmt"
	"os"
)

func allUnique(s string) bool {
	e := map[rune]struct{}{}
	for _, r := range s {
		_, exists := e[r]
		if exists {
			return false
		}
		e[r] = struct{}{}
	}
	return true
}

func part1(input string) int {
	for i := range input {
		if i > 4 {
			key := input[i-4 : i]
			if allUnique(key) {
				return i
			}
		}
	}
	return -1
}

func part2(input string) int {
	for i := range input {
		if i > 14 {
			key := input[i-14 : i]
			if allUnique(key) {
				return i
			}
		}
	}
	return -1
}

func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	noError(err)
	return string(file)
}

func main() {
	fmt.Println("Part1", part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	fmt.Println("Part1", part1(readFile("input")))
	fmt.Println("Part2", part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	fmt.Println("Part2", part2(readFile("input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
