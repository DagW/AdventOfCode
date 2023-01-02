package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	num    int
	cycles int
	t      string
}

func part1(input []instruction) int {
	cycle := 0
	x := 1
	sum := 0
	for _, instruction := range input {
		// fmt.Println(instruction)
		cycles := 1
		if instruction.t == "addx" {
			cycles = 2
		}
		for i := 0; i < cycles; i++ {
			cycle++
			if (cycle-20)%40 == 0 {
				fmt.Println(cycle, "regval", x)
				sum += cycle * x
			}
		}
		if instruction.t == "noop" {
		} else if instruction.t == "addx" {
			x += instruction.num
		}
	}
	return sum
}

func part2(input []instruction) int {
	cycle := 0
	x := 1

	pixels := make([]bool, 240)

	for _, instruction := range input {
		cycles := 1
		if instruction.t == "addx" {
			cycles = 2
		}
		for i := 0; i < cycles; i++ {
			cycle++
			c := (cycle % 40) - 1
			if c == x || c == x-1 || c == x+1 {
				pixels[cycle-1] = true
			}
		}
		if instruction.t == "noop" {
		} else if instruction.t == "addx" {
			x += instruction.num
		}
	}

	c := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if pixels[c] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
			c++
		}
		fmt.Println()
	}
	return 0
}
func readFile(filename string) []instruction {
	file, err := os.ReadFile(filename)
	noError(err)
	var grid []instruction
	for _, val := range strings.Split(string(file), "\n") {
		splits := strings.Split(val, " ")
		i := 0
		if len(splits) > 1 {
			i, _ = strconv.Atoi(splits[1])
		}
		grid = append(grid, instruction{t: splits[0], num: i})
	}
	return grid
}

func main() {
	fmt.Println(part1(readFile("test")))
	fmt.Println(part1(readFile("input")))
	fmt.Println(part2(readFile("test")))
	fmt.Println(part2(readFile("input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
