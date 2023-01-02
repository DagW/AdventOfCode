package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	num       int
}
type position struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func notBiggerThanOne(a int) int {
	if a < 0 {
		if a < -1 {
			return -1
		}
		return a
	}
	if a > 1 {
		return 1
	}
	return a
}

func printMap(visited map[position]struct{}, ps ...position) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			c := "."
			p := position{x: x, y: 9 - y}
			if _, v := visited[p]; v {
				c = "#"
			}
			for i, p2 := range ps {
				if p == p2 {
					c = fmt.Sprintf("%d", i)
					break
				}
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(input []instruction) int {
	visited := map[position]struct{}{
		{}: {},
	}
	head, tail := position{}, position{}
	for _, instruction := range input {
		for i := 0; i < instruction.num; i++ {
			switch instruction.direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
				if abs(head.x-tail.x) > 0 && abs(head.y-tail.y) > 0 {
					tail.x += notBiggerThanOne(head.x - tail.x)
					tail.y += notBiggerThanOne(head.y - tail.y)
				} else if abs(head.x-tail.x) > 0 {
					tail.x += notBiggerThanOne(head.x - tail.x)
				} else if abs(head.y-tail.y) > 0 {
					tail.y += notBiggerThanOne(head.y - tail.y)
				}
				visited[tail] = struct{}{}
			}
		}
	}

	printMap(visited, head, tail)
	return len(visited)
}

func part2(input []instruction) int {
	visited := map[position]struct{}{{}: {}}
	head, tails := position{}, make([]position, 9)
	for _, instruction := range input {
		for i := 0; i < instruction.num; i++ {
			switch instruction.direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			for i, tail := range tails {
				last := head
				if i > 0 {
					last = tails[i-1]
				}
				if abs(last.x-tail.x) > 1 || abs(last.y-tail.y) > 1 {
					if abs(last.x-tail.x) > 0 && abs(last.y-tail.y) > 0 {
						tail.x += notBiggerThanOne(last.x - tail.x)
						tail.y += notBiggerThanOne(last.y - tail.y)
					} else if abs(last.x-tail.x) > 0 {
						tail.x += notBiggerThanOne(last.x - tail.x)
					} else if abs(last.y-tail.y) > 0 {
						tail.y += notBiggerThanOne(last.y - tail.y)
					}
					if i == 8 {
						visited[tail] = struct{}{}
					}
				}

				tails[i] = tail
			}
			// printMap(visited, append([]position{head}, tails...)...)
		}
	}
	return len(visited)
}

func readFile(filename string) []instruction {
	file, err := os.ReadFile(filename)
	noError(err)
	var grid []instruction
	for _, val := range strings.Split(string(file), "\n") {
		splits := strings.Split(val, " ")
		i, err := strconv.Atoi(splits[1])
		noError(err)
		grid = append(grid, instruction{direction: splits[0], num: i})
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
