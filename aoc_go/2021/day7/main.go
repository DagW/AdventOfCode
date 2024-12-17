package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func minMax(ints []int) (int, int) {
	min, max := math.MaxInt, 0
	for _, v := range ints {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func bestCost(crabSubmarines []int, costFunc costFunction) int {
	min, max := minMax(crabSubmarines)
	bestPos, leastCost := -1, math.MaxInt
	for pos := min; pos <= max; pos++ {
		cost := costFunc(pos, crabSubmarines)
		if cost < leastCost {
			leastCost, bestPos = cost, pos
		}
	}
	_ = bestPos
	return leastCost
}

type costFunction func(pos int, crabSubmarines []int) int

func part2cost(pos int, crabSubmarines []int) int {
	cost := 0
	for _, cs := range crabSubmarines {
		a := abs(cs - pos)
		for j := 1; j <= a; j++ {
			cost += j
		}
	}
	return cost
}

func part1cost(pos int, crabSubmarines []int) int {
	cost := 0
	for _, cs := range crabSubmarines {
		cost += abs(cs - pos)
	}
	return cost
}

func part1(crabSubmarines []int) (score int) {
	return bestCost(crabSubmarines, part1cost)
}
func part2(crabSubmarines []int) (score int) {
	return bestCost(crabSubmarines, part2cost)
}

func readFile(filename string) (values []int) {
	file, _ := ioutil.ReadFile(filename)
	nums := strings.Split(string(file), ",")
	for _, s := range nums {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	return values
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day7/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day7/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day7/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day7/input")))
}
