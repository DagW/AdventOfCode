package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(input []int) int {
	for _, item1 := range input {
		for _, item2 := range input {
			if item1 != item2 && (item1+item2) == 2020 {
				return item1 * item2
			}
		}
	}
	return -1
}

func part2(input []int) int {
	for _, item1 := range input {
		for _, item2 := range input {
			for _, item3 := range input {
				if (item1 != item2) && (item1 != item3) && (item2 != item3) && (item1+item2+item3) == 2020 {
					return item1 * item2 * item3
				}
			}
		}
	}
	return -1
}

func readFile(filename string) []int {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var program []int
	for _, val := range strings.Split(string(file), "\n") {
		i, _ := strconv.Atoi(val)
		program = append(program, i)
	}
	return program
}

func main() {
	fmt.Println("Part1", part1([]int{1721, 979, 366, 299, 675, 1456}))
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day1/input")))
	fmt.Println("Part2", part2([]int{1721, 979, 366, 299, 675, 1456}))
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day1/input")))
}
