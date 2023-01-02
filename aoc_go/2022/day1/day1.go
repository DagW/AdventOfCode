package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(input [][]int) int {
	maxCals := 0
	for _, e := range input {
		sum := 0
		for _, v := range e {
			sum += v
		}
		if sum > maxCals {
			maxCals = sum
		}
	}
	return maxCals
}

func part2(input [][]int) int {
	var elves []int
	for _, e := range input {
		sum := 0
		for _, v := range e {
			sum += v
		}
		elves = append(elves, sum)
	}
	sort.Ints(elves)
	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}

func readFile(filename string) [][]int {
	file, err := os.ReadFile(filename)
	noError(err)
	var elves [][]int
	var elf []int
	for _, val := range strings.Split(string(file), "\n") {
		if val == "" {
			elves = append(elves, elf)
			elf = []int{}
			continue
		}
		i, err := strconv.Atoi(val)
		noError(err)
		elf = append(elf, i)
	}
	return elves
}

func main() {
	fmt.Println("Part1 test", part1(readFile("aoc_go/2022/day1/day1.test")))
	fmt.Println("Part1", part1(readFile("aoc_go/2022/day1/day1.input")))
	fmt.Println("Part2", part2(readFile("aoc_go/2022/day1/day1.input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
