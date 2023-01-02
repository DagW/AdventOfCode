package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part2(stacks []string, instructions [][]int) string {
	for _, instruction := range instructions {
		amount := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1

		popped := stacks[from][:amount]
		stacks[from] = stacks[from][amount:]
		stacks[to] = string(popped) + stacks[to]
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			result += string(stack[0])
		}
	}
	return result
}
func part1(stacks []string, instructions [][]int) string {
	for _, instruction := range instructions {
		amount := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1
		for i := 0; i < amount; i++ {
			popped := stacks[from][0]
			stacks[from] = stacks[from][1:]
			stacks[to] = string(popped) + stacks[to]
		}
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			result += string(stack[0])
		}
	}
	return result
}

func readFile(filename string) ([]string, [][]int) {
	file, err := os.ReadFile(filename)
	noError(err)
	stacks := make([]string, 100, 100)
	var instructions [][]int
	for _, val := range strings.Split(string(file), "\n") {
		if strings.Contains(val, "move") {
			re := regexp.MustCompile("[0-9]+")
			instructions = append(instructions, tointslice(re.FindAllString(val, -1)))
		} else {
			idx := 0
			for {
				if idx+1 > len(val) {
					break
				}
				if val[idx] == '[' {
					stacks[idx/4] += string(val[idx+1])
				}
				idx += 4
			}
		}
	}
	return stacks, instructions
}

func main() {
	fmt.Println("Part1", part1(readFile("test")))
	fmt.Println("Part1", part1(readFile("input")))
	fmt.Println("Part2", part2(readFile("test")))
	fmt.Println("Part2", part2(readFile("input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}

func tointslice(strings []string) (r []int) {
	r = make([]int, len(strings))
	for i, s := range strings {
		c, err := strconv.Atoi(s)
		noError(err)
		r[i] = c
	}
	return r
}
