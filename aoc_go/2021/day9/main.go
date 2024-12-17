package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) (m [][]int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, s := range lines {
		var l []int
		for _, c := range s {
			n, _ := strconv.Atoi(string(c))
			l = append(l, n)
		}
		m = append(m, l)
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day9/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day9/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day9/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day9/input")))
}

func part1(file [][]int) int {
	sum := 0

	for r, row := range file {
		for c, value := range row {
			lowPoint := true
			for _, neighbour := range neighbours {
				nx, ny := c+neighbour[0], r+neighbour[1]
				if nx >= 0 && nx < len(row) && ny >= 0 && ny < len(file) {
					if value >= file[ny][nx] {
						lowPoint = false
					}
				}
			}
			if lowPoint {
				sum += value + 1
			}
		}
	}
	return sum
}

var neighbours = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func part2(file [][]int) int {
	var positions [][]int
	for r, row := range file {
		for c, value := range row {
			lowPoint := true
			for _, neighbour := range neighbours {
				nx, ny := c+neighbour[0], r+neighbour[1]
				if nx >= 0 && nx < len(row) && ny >= 0 && ny < len(file) {
					if value >= file[ny][nx] {
						lowPoint = false
					}
				}
			}
			if lowPoint {
				positions = append(positions, []int{r, c})
			}
		}
	}

	var res []int
	for _, pos := range positions {
		copied := make([][]int, 0, len(file))
		for _, c := range file {
			n := append([]int{}, c...)
			copied = append(copied, n)
		}
		res = append(res, fillNeighbours(copied, pos))
	}

	n1, res := getMax(res)
	n2, res := getMax(res)
	n3, res := getMax(res)

	return n1 * n2 * n3
}

func getMax(res []int) (int, []int) {
	max, maxPos := 0, 0
	for i, v := range res {
		if v > max {
			max = v
			maxPos = i
		}
	}
	return max, append(res[:maxPos], res[maxPos+1:]...)
}

func fillNeighbours(m [][]int, pos []int) int {
	value := m[pos[0]][pos[1]]
	count := 1
	for _, n := range neighbours {
		ny, nx := pos[0]+n[0], pos[1]+n[1]
		if nx >= 0 && nx < len(m[0]) && ny >= 0 && ny < len(m) {
			neighbour := m[ny][nx]
			if neighbour > value && neighbour < 9 {
				count += fillNeighbours(m, []int{ny, nx})
			}
		}
	}
	m[pos[0]][pos[1]] = 10
	return count
}
