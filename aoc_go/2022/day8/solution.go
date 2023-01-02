package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxIn(p int, in []int) bool {
	for _, v := range in {
		if p <= v {
			return false
		}
	}
	return true
}

func part1(input [][]int) int {
	visible := 0
	for y, row := range input {
		for x, col := range row {
			if y == 0 || x == 0 || y == len(input)-1 || x == len(row)-1 {
				visible++
				continue
			}

			left := row[0:x]
			right := []int{}
			if x < len(row) {
				right = row[x+1 : len(row)]
			}

			up, down := []int{}, []int{}
			for ny := 0; ny < len(input); ny++ {
				if ny < y {
					up = append(up, input[ny][x])
				} else if ny > y {
					down = append(down, input[ny][x])
				}
			}

			for _, arr := range [][]int{up, down, left, right} {
				if maxIn(col, arr) {
					visible++
					break
				}
			}
		}
	}

	return visible
}

func part2(input [][]int) int {
	highscore := 0
	for y, row := range input {
		for x, col := range row {

			left := 0
			for i := x - 1; i >= 0; i-- {
				t := input[y][i]
				left++
				if t >= col {
					break
				}
			}

			right := 0
			for i := x + 1; i < len(row); i++ {
				t := input[y][i]
				right++
				if t >= col {
					break
				}
			}

			up := 0
			for i := y - 1; i >= 0; i-- {
				t := input[i][x]
				up++
				if t >= col {
					break
				}
			}

			down := 0
			for i := y + 1; i < len(input); i++ {
				t := input[i][x]
				down++
				if t >= col {
					break
				}
			}
			// fmt.Println(x, y, "has score", left*right*up*down, left, right, up, down)
			if calc := left * right * up * down; calc > highscore {
				highscore = calc
			}
		}
	}

	return highscore
}

func readFile(filename string) [][]int {
	file, err := os.ReadFile(filename)
	noError(err)
	var grid [][]int
	for _, val := range strings.Split(string(file), "\n") {
		row := []int{}
		for _, r := range val {
			i, err := strconv.Atoi(string(r))
			noError(err)
			row = append(row, i)
		}
		grid = append(grid, row)
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
