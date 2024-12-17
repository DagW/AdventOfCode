package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	UP      = iota
	DOWN    = iota
	FORWARD = iota
)

type instruction struct {
	direction int
	amount    int
}

func part1(values []instruction) (m int) {
	var horizontal, depth int
	for _, v := range values {
		switch v.direction {
		case UP:
			depth -= v.amount
		case DOWN:
			depth += v.amount
		case FORWARD:
			horizontal += v.amount
		}
	}
	return horizontal * depth
}

func part2(values []instruction) (m int) {
	var horizontal, depth, aim int
	for _, v := range values {
		switch v.direction {
		case UP:
			aim -= v.amount
		case DOWN:
			aim += v.amount
		case FORWARD:
			horizontal += v.amount
			depth += v.amount * aim
		}
	}
	return horizontal * depth
}

func readFile(filename string) (values []instruction) {
	dirs := map[string]int{"forward": FORWARD, "up": UP, "down": DOWN}
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		v, _ := strconv.Atoi(parts[1])
		values = append(values, instruction{
			direction: dirs[parts[0]],
			amount:    v,
		})
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day2/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day2/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day2/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day2/input")))
}
