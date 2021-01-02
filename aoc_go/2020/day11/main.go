package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	empty = iota
	occupied
	floor
)

func printMap(m [][]int) {
	fmt.Println("----")
	for _, r := range m {
		for _, c := range r {
			if c == empty {
				fmt.Print("L")
			} else if c == occupied {
				fmt.Print("#")
			} else if c == floor {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----")
}

func noOccupied(m [][]int, y, x, r int) bool {
	test := [8][2]int{
		{-1, 1},
		{0, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	for _, t := range test {
		//for multiplier := 1; multiplier < len(m); multiplier++ {
		tx, ty := x+t[1], y+t[0]
		//tx, ty = tx*multiplier, ty*multiplier
		if ty < 0 || tx < 0 ||
			ty > len(m)-1 || tx > len(m[1])-1 {
			continue
		}
		if m[ty][tx] == occupied {
			return false
		}
		//}
	}
	return true
}
func numOccupied(m [][]int, y, x, limit int) bool {
	test := [8][2]int{
		{-1, 1},
		{0, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	numOccupied := 0
	for _, t := range test {
		tx, ty := x+t[1], y+t[0]
		if ty < 0 || tx < 0 || ty > len(m)-1 || tx > len(m[1])-1 {
			continue
		}
		if m[ty][tx] == occupied {
			numOccupied++
		}
		if numOccupied >= limit {
			return true
		}
	}
	return false
}

func doRound(m [][]int, r, limit int) ([][]int, int) {
	changes := 0
	newMap := make([][]int, len(m))
	for y := range m {
		newMap[y] = make([]int, len(m[0]))
	}
	for y, row := range m {
		for x, col := range row {
			if col == empty && noOccupied(m, y, x, r) {
				newMap[y][x] = occupied
				changes++
			} else if col == occupied && numOccupied(m, y, x, limit) {
				newMap[y][x] = empty
				changes++
			} else {
				newMap[y][x] = col
			}

		}
	}
	return newMap, changes
}

func doRounds(m [][]int, r, limit int) int {
	changes := 1
	for changes > 0 {
		m, changes = doRound(m, r, limit)
	}
	num := 0
	for _, row := range m {
		for _, col := range row {
			if col == occupied {
				num++
			}
		}
	}
	return num
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
				row = append(row, floor)
			} else if c == '#' {
				row = append(row, occupied)
			} else if c == 'L' {
				row = append(row, empty)
			}
		}
		seats = append(seats, row)
	}
	return
}

func main() {
	test, input := readFile("test"), readFile("input")
	fmt.Println(doRounds(test, 1, 4), 37)
	fmt.Println(doRounds(input, 1, 4), 2261)
	//fmt.Println(readFile("aoc_go/2020/day11/input"))
	fmt.Println(doRounds(test, len(input), 5))
	fmt.Println(doRounds(input, len(input), 5))
}
