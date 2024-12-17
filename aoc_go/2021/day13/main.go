package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) (m [][2]int, ins []string) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	t := 0
	for _, s := range lines {
		if s == "" {
			t = 1
		}
		if t == 0 {
			parts := strings.Split(s, ",")
			y, _ := strconv.Atoi(parts[0])
			x, _ := strconv.Atoi(parts[1])
			m = append(m, [2]int{y, x})
		} else if s != "" {
			ins = append(ins, strings.Split(s, "fold along ")[1])
		}
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", solve(readFile("aoc_go/2021/day13/test")))
	fmt.Println("Input", solve(readFile("aoc_go/2021/day13/input")))
}

func solve(input [][2]int, instructions []string) int {
	m := toMap(input)

	for _, instruction := range instructions {
		fmt.Println(instruction)
		ps := strings.Split(instruction, "=")
		n, _ := strconv.Atoi(ps[1])
		if ps[0] == "y" {
			newMap := make([][]int, len(m))
			for i := range m {
				newMap[i] = append(newMap[i], m[i][:n]...)
				for j := range m[i][n:] {
					if m[i][n+j] > 0 {
						newMap[i][n-j] += m[i][n+j]
					}
				}
			}
			m = newMap
		} else {
			// Row folding
			newMap := make([][]int, n)
			for i := range m {
				if i < n {
					newMap[i] = append(newMap[i], m[i]...)
				} else if i > n {
					for j := range m[i] {
						if m[i][j] > 0 {
							newMap[n+(n-i)][j] += m[i][j]
						}
					}
				}
			}
			m = newMap
		}
		// break - part1
	}
	return printMap(m)
}

func printMap(m [][]int) (sum int) {
	for x := range m[0] {
		for y := range m {
			if m[y][x] > 0 {
				sum++
			}
			if m[y][x] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return sum
}

func toMap(input [][2]int) [][]int {
	maxY, maxX := getMax(input)
	m := make([][]int, maxY+1)
	for i := 0; i < maxY+1; i++ {
		m[i] = make([]int, maxX+1)
	}
	for _, p := range input {
		m[p[0]][p[1]] = 1
	}
	return m
}

func getMax(input [][2]int) (y int, x int) {
	for _, v := range input {
		if v[0] > y {
			y = v[0]
		}
		if v[1] > x {
			x = v[1]
		}
	}
	return

}
