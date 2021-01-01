package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	FREE     = iota
	OCCUPIED = iota
	FLOOR    = iota
)

func printMap(m [][]int) {
	fmt.Println("----")
	for _, r := range m {
		for _, c := range r {
			if c == FREE {
				fmt.Print("L")
			} else if c == OCCUPIED {
				fmt.Print("#")
			} else if c == FLOOR {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----")
}

func allAroundMeAreFree(m [][]int, y, x int) bool {
	test := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, t := range test {
		tx, ty := x-t[1], y-t[0]
		if ty > 0 || tx > 0 || ty < len(m) || tx < len(m[1]) {
			continue
		}
		fmt.Println(tx, ty)
		if m[ty][tx] == OCCUPIED {
			return false
		}
	}
	return true
}
func allAroundMeAreOccupied(m [][]int, y, x int) bool {
	test := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, t := range test {
		tx, ty := x-t[1], y-t[0]
		if ty > 0 || tx > 0 || ty < len(m) || tx < len(m[1]) {
			continue
		}
		fmt.Println(tx, ty)
		if m[ty][tx] == FREE {
			return false
		}
	}
	return true
}

func part1(m [][]int) int {
	fmt.Println("INPUT")
	printMap(m)
	fmt.Println("STARTING")
	changes := -1
	count := 0
	for changes != 0 && count != 10 {

		fmt.Println()
		fmt.Println(m)
		fmt.Println()
		fmt.Println("BEFORE", changes)
		printMap(m)
		count++
		changes = 0
		m2 := make([][]int, len(m))
		for i := range m {
			m2[i] = make([]int, len(m[i]))
			copy(m2[i], m[i])
		}
		for y, r := range m {
			for x, c := range r {
				if c == FREE && allAroundMeAreFree(m, y, x) {
					m2[y][x] = OCCUPIED
					changes++
				}
			}
		}
		for y, r := range m2 {
			for x, c := range r {
				if c == OCCUPIED && allAroundMeAreOccupied(m, y, x) {
					m2[y][x] = FREE
					changes++
				}
			}
		}

		fmt.Println(count, "AFTER", changes)
		m = m2
		printMap(m)
	}
	sum := 0
	for _, r := range m {
		for _, c := range r {
			if c == OCCUPIED {
				sum++
			}
		}
	}
	return sum
}

func readFile(filename string) (seats [][]int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		row := make([]int, 0)
		for _, c := range line {
			if c == '.' {
				row = append(row, FLOOR)
			} else if c == '#' {
				row = append(row, OCCUPIED)
			} else if c == 'L' {
				row = append(row, FREE)
			}
		}
		seats = append(seats, row)
	}
	return
}

func main() {
	fmt.Println(part1(readFile("aoc_go/2020/day11/test")))
	//fmt.Println(readFile("aoc_go/2020/day11/input"))
}
