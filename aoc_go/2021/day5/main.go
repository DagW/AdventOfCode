package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func getRange(x, y int) (res []int) {
	min, max, reversed := x, y, false
	if x > y {
		min, max, reversed = y, x, true
	}
	for i := min; i <= max; i++ {
		res = append(res, i)
	}
	if reversed {
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
		}
	}
	return res
}

func solve(ranges []int, withDiagonal bool) (score int) {
	m := map[point]int{}
	for i := 0; i < len(ranges); i += 4 {
		x1, y1, x2, y2 := ranges[i], ranges[i+1], ranges[i+2], ranges[i+3]

		adds := map[point]struct{}{}
		if x1 == x2 || y1 == y2 {
			for _, x := range getRange(x1, x2) {
				adds[point{x, y1}] = struct{}{}
			}
			for _, y := range getRange(y1, y2) {
				adds[point{x1, y}] = struct{}{}
			}
		} else if withDiagonal {
			r1, r2 := getRange(x1, x2), getRange(y1, y2)
			for i := 0; i < len(r1); i++ {
				adds[point{r1[i], r2[i]}] = struct{}{}
			}
		}
		for k := range adds {
			m[k]++
		}
		// fmt.Println(x1, y1, "->", x2, y2)
		// printMap(m)
	}
	pts := 0
	for _, v := range m {
		if v > 1 {
			pts++
		}
	}
	return pts
}

func printMap(m map[point]int) {
	maxX, maxY := 0, 0
	for k := range m {
		if k.x >= maxX {
			maxX = k.x
		}
		if k.y >= maxY {
			maxY = k.y
		}
	}
	m2 := make([][]int, maxY+1)
	for i := 0; i < len(m2); i++ {
		m2[i] = make([]int, maxX+1)
	}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			m2[y][x] = m[point{x, y}]
		}
	}
	for i := range m2 {
		fmt.Println(m2[i])
	}
}

func readFile(filename string) (ranges []int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		for _, s := range strings.Split(strings.Replace(line, " ", "", -1), "->") {
			for _, s2 := range strings.Split(s, ",") {
				n, _ := strconv.Atoi(s2)
				ranges = append(ranges, n)
			}
		}
	}
	return ranges
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", solve(readFile("aoc_go/2021/day5/test"), false))
	fmt.Println("Input", solve(readFile("aoc_go/2021/day5/input"), false))
	fmt.Println("Part2")
	fmt.Println("Test", solve(readFile("aoc_go/2021/day5/test"), true))
	fmt.Println("Input", solve(readFile("aoc_go/2021/day5/input"), true))
}
