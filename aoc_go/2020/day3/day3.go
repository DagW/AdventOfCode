package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2(m [][]bool) int {
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	total := 1
	for _, slope := range slopes {
		total *= countCollisions(m, slope[0], slope[1])
	}
	return total
}

func part1(m [][]bool) int {
	return countCollisions(m, 3, 1)
}

func countCollisions(m [][]bool, slopeX int, slopeY int) int {
	jumps, collisions, x, y := 0, 0, 0, 0
	for y < len(m) {
		//fmt.Println(x, y, jumps, collisions, "istree=", m[y][x])
		if m[y][x] {
			collisions++
		}
		x += slopeX
		y += slopeY
		x = x % len(m[0])
		jumps++
	}
	return collisions
}
func readFile(filename string) [][]bool {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	var m [][]bool
	for _, val := range lines {
		var linechars []bool
		for _, c := range val {
			linechars = append(linechars, c == '#')
		}
		m = append(m, linechars)
	}
	return m
}

func main() {

	fmt.Println("Part1", part1(readFile("aoc_go/2020/day3/day3.test0")), 7)
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day3/day3.input")), 181)
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day3/day3.test0")), 336)
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day3/day3.input")), 1260601650)

}
