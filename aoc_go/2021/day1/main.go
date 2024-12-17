package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// part1 counts increases from previous value
func part1(values []int) (increases int) {
	for i, v := range values {
		if i != 0 && v > values[i-1] {
			increases++
		}
	}
	return increases
}

// part2 uses a sliding window sum
func part2(values []int) (increases int) {
	a, b, c := 0, 0, 0
	for i, v := range values {
		if i > 2 && (b+c+v) > (a+b+c) {
			increases++
		}
		a, b, c = b, c, v
	}
	return increases
}

func readFile(filename string) (values []int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		values = append(values, v)
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day1/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day1/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day1/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day1/input")))
}
