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

func numOccupied(m [][]int, y, x, r int) int {
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
	for t := 0; t < len(test); t++ {
		direction := test[t]
		for i := 0; i < r; i++ {
			ty, tx := y+(direction[1]*(i+1)), x+(direction[0]*(i+1))
			if ty < 0 || tx < 0 || ty > len(m)-1 || tx > len(m[1])-1 {
				break
			}
			if m[ty][tx] == occupied {
				numOccupied++
				break
			}
			if m[ty][tx] == empty {
				break
			}
		}
	}
	return numOccupied
}

func iteratePassengers(m [][]int, radius, limit int) ([][]int, int) {
	changes := 0
	// Make all the changes on a new map
	newMap := make([][]int, len(m))
	for y := range m {
		newMap[y] = make([]int, len(m[0]))
	}
	for y, row := range m {
		for x, col := range row {
			if col == empty && numOccupied(m, y, x, radius) == 0 {
				newMap[y][x] = occupied
				changes++
			} else if col == occupied && numOccupied(m, y, x, radius) >= limit {
				newMap[y][x] = empty
				changes++
			} else {
				newMap[y][x] = col
			}
		}
	}
	return newMap, changes
}

func solve(m [][]int, radius, limit int) int {
	changes := 1
	for changes > 0 {
		m, changes = iteratePassengers(m, radius, limit)
		count++
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
	fmt.Println(solve(test, 1, 4), 37)
	fmt.Println(solve(input, 1, 4), 2261)
	fmt.Println(solve(test, len(test), 5))
	fmt.Println(solve(input, len(input), 5))
}
