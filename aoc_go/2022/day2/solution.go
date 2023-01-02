package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(input [][]string) int {
	var c = map[string]int{"X": 1, "Y": 2, "Z": 3}
	var r = map[string]map[string]int{
		"A": {"X": 3, "Y": 6, "Z": 0},
		"B": {"X": 0, "Y": 3, "Z": 6},
		"C": {"X": 6, "Y": 0, "Z": 3},
	}

	score := 0
	for _, v := range input {
		score += r[v[0]][v[1]] + c[v[1]]
	}
	return score
}

func part2(input [][]string) int {
	var c = map[string]int{"X": 0, "Y": 3, "Z": 6}
	var r = map[string]map[string]int{
		"A": {"X": 3, "Y": 1, "Z": 2},
		"B": {"X": 1, "Y": 2, "Z": 3},
		"C": {"X": 2, "Y": 3, "Z": 1},
	}

	score := 0
	for _, v := range input {
		score += r[v[0]][v[1]] + c[v[1]]
	}
	return score
}

func readFile(filename string) [][]string {
	file, err := os.ReadFile(filename)
	noError(err)
	var rps [][]string
	for _, val := range strings.Split(string(file), "\n") {
		parts := strings.Split(val, " ")
		rps = append(rps, parts)
	}
	return rps
}

func main() {
	fmt.Println("Part1 test", part1(readFile("aoc_go/2022/day2/test")))
	fmt.Println("Part1", part1(readFile("aoc_go/2022/day2/input")))
	fmt.Println("Part2 test", part2(readFile("aoc_go/2022/day2/test")))
	fmt.Println("Part2", part2(readFile("aoc_go/2022/day2/input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
