package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func part2(input [][]position) int {
	m := toMap(input)
	// printMap(m)

	start := position{x: 500, y: 0}
	_, maxY, _, _ := edges(m)

	c := 0
	for {
		sandPosition := start
		landed := false
		for i := 0; i < 1000; i++ {
			sandPosition.y++
			if sandPosition.y >= maxY+2 {
				// Hit bottom!
				sandPosition.y--
				landed = true
				break
			}
			if _, below := m[sandPosition]; !below {
				// Grab it
				continue
			}
			sandPosition.x--
			if _, left := m[sandPosition]; !left {
				// Grab it
				continue
			}
			sandPosition.x += 2
			if _, right := m[sandPosition]; !right {
				// Grab it
				continue
			}
			// reset, we could not move. Land
			sandPosition.x--
			sandPosition.y--
			landed = true
			break
		}

		if landed {
			m[sandPosition] = struct{}{}
		} else {
			// infinite
			break
		}
		c++

		// printMap(m)
		if sandPosition == start {
			break
		}
	}
	return c
}

func part1(input [][]position) int {
	m := toMap(input)
	printMap(m)

	c := 0
	for {
		sandPosition := position{x: 500, y: 0}
		landed := false
		for i := 0; i < 1000; i++ {
			sandPosition.y++
			if _, below := m[sandPosition]; !below {
				// Grab it
				continue
			}
			sandPosition.x--
			if _, left := m[sandPosition]; !left {
				// Grab it
				continue
			}
			sandPosition.x += 2
			if _, right := m[sandPosition]; !right {
				// Grab it
				continue
			}
			// reset, we could not move. Land
			sandPosition.x--
			sandPosition.y--
			landed = true
			break
		}

		if landed {
			m[sandPosition] = struct{}{}
		} else {
			// infinite
			break
		}
		c++
		// printMap(m)
	}
	return c
}

func printMap(m map[position]struct{}) {
	minY, maxY, minX, maxX := edges(m)
	fmt.Println("Edges", minY, maxY, minX, maxX)
	for y := minY - 5; y < maxY+5; y++ {
		for x := minX - 5; x < maxX+5; x++ {
			s := " ."
			if _, exists := m[position{y: y, x: x}]; exists {
				s = " #"
			} else if y == 0 && x == 500 {
				s = " v"
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

func edges(m map[position]struct{}) (int, int, int, int) {
	minY, maxY, minX, maxX := math.MaxInt64, math.MinInt64, math.MaxInt64, math.MinInt64
	for pos := range m {
		maxX = max(maxX, pos.x)
		minX = min(minX, pos.x)
		maxY = max(maxY, pos.y)
		minY = min(minY, pos.y)
	}
	return minY, maxY, minX, maxX
}

func toMap(input [][]position) map[position]struct{} {
	m := map[position]struct{}{}
	for _, positions := range input {
		last := position{}
		for i, p := range positions {
			if i > 0 {
				for y := min(last.y, p.y); y <= max(last.y, p.y); y++ {
					for x := min(last.x, p.x); x <= max(last.x, p.x); x++ {
						m[position{x: x, y: y}] = struct{}{}
					}
				}
			}
			last = p
		}
	}
	return m
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func readFile(filename string) [][]position {
	file, err := os.ReadFile(filename)
	noError(err)
	var row [][]position
	for _, val := range strings.Split(string(file), "\n") {
		var positions []position
		for _, v := range strings.Split(val, " -> ") {
			strs := strings.Split(v, ",")
			positions = append(positions, position{
				x: toInt(strs[0]),
				y: toInt(strs[1]),
			})
		}
		row = append(row, positions)
	}
	return row
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	noError(err)
	return i
}

func main() {
	fmt.Println(part1(readFile("aoc_go/2022/day14/test")))
	fmt.Println(part1(readFile("aoc_go/2022/day14/input")))
	fmt.Println(part2(readFile("aoc_go/2022/day14/test")))
	fmt.Println(part2(readFile("aoc_go/2022/day14/input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
