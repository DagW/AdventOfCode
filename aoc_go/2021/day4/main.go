package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part2(draws []int, boards [][][]int) (score int) {
	for _, d := range draws {
		for i := 0; i < len(boards); i++ {
			b := boards[i]
			for x := range b {
				for y, v := range b[x] {
					if v == d {
						b[x][y] = -1
					}
				}
			}
			if checkForWinners(b) {
				if len(boards) == 1 {
					return getScore(boards[0]) * d
				}
				boards = append(boards[:i], boards[i+1:]...)
				i--
			}
		}
	}
	return 0
}

func part1(draws []int, boards [][][]int) (score int) {
	for _, d := range draws {
		for _, b := range boards {
			for x := range b {
				for y, v := range b[x] {
					if v == d {
						b[x][y] = -1
					}
				}
			}
			if checkForWinners(b) {
				return getScore(b) * d
			}
		}
	}
	return 0
}

func getScore(b [][]int) int {
	sum := 0
	for x := range b {
		for _, v := range b[x] {
			if v != -1 {
				sum += v
			}
		}
	}
	return sum
}

func checkForWinners(b [][]int) bool {
	for x := range b {
		if b[x][0] == -1 && b[x][1] == -1 && b[x][2] == -1 && b[x][3] == -1 && b[x][4] == -1 {
			return true
		}
	}
	for y := range b[0] {
		if b[0][y] == -1 && b[1][y] == -1 && b[2][y] == -1 && b[3][y] == -1 && b[4][y] == -1 {
			return true
		}
	}
	return false
}

func readFile(filename string) (draws []int, boards [][][]int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	board := make([][]int, 0, 5)
	for i, line := range lines {
		if i == 0 {
			for _, s := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(s)
				draws = append(draws, n)
			}
		} else if line != "" {
			row := make([]int, 0, 5)
			for _, s := range strings.Split(strings.Trim(line, " "), " ") {
				if s != "" {
					n, _ := strconv.Atoi(s)
					row = append(row, n)
				}
			}
			board = append(board, row)
			if len(board) == 5 {
				boards = append(boards, board)
				board = make([][]int, 0, 5)
			}
		}
	}
	return draws, boards
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day4/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day4/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day4/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day4/input")))
}
