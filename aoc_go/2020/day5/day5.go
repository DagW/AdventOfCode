package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type seat struct {
	row, col, id int
}

func parsePasses(boardingPasses []string) []seat {
	var results []seat
	for _, bp := range boardingPasses {
		row, col := [2]int{1, 128}, [2]int{1, 8}
		for _, c := range bp[0:7] {
			if c == 'F' {
				//minus half the diff
				row[1] -= (row[1] - (row[0] - 1)) / 2
			} else {
				//plus half the diff
				row[0] += (row[1] - (row[0] - 1)) / 2
			}
		}
		for _, c := range bp[7:] {
			if c == 'L' {
				col[1] -= (col[1] - (col[0] - 1)) / 2
			} else {
				col[0] += (col[1] - (col[0] - 1)) / 2
			}
		}
		row[0], row[1], col[0], col[1] = row[0]-1, row[1]-1, col[0]-1, col[1]-1
		// fmt.Println(bp[0:7], bp[7:], row, col, (row[0]*8)+col[0])
		results = append(results, seat{
			row: row[0],
			col: col[0],
			id:  (row[0] * 8) + col[0],
		})
	}
	return results
}

func maxSeatId(arr []seat) int {
	max := 0
	for _, item := range arr {
		if item.id > max {
			max = item.id
		}
	}
	return max
}

func readFile(filename string) []string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test:", maxSeatId(parsePasses(readFile("aoc_go/2020/day5/test"))))
	passes := parsePasses(readFile("aoc_go/2020/day5/input"))
	fmt.Println("Input:", maxSeatId(passes))

	fmt.Println("Part2")
	occupiedSeats := [128][8]bool{}
	for _, pass := range passes {
		occupiedSeats[pass.row][pass.col] = true
	}

	for y, row := range occupiedSeats {
		if row == [8]bool{} ||
			occupiedSeats[y-1] == [8]bool{} ||
			occupiedSeats[y+1] == [8]bool{} {
			// Back and front of plane
		} else {
			for x, col := range row {
				if !col {
					fmt.Println("row", y, "seat", x, "is free", (y*8)+x)
				}
			}
		}
	}
}
