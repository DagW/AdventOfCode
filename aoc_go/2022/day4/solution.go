package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pts struct {
	min, max int
}

func part1(input [][]pts) int {
	sum := 0
	for _, p := range input {
		if p[0].min >= p[1].min && p[0].max <= p[1].max {
			sum++
		} else if p[1].min >= p[0].min && p[1].max <= p[0].max {
			sum++
		}
	}
	return sum
}

func part2(input [][]pts) int {
	sum := 0
	for _, p := range input {
		c := map[int]struct{}{}
		for i := p[0].min; i <= p[0].max; i++ {
			c[i] = struct{}{}
		}
		for i := p[1].min; i <= p[1].max; i++ {
			if _, found := c[i]; found {
				sum++
				break
			}
		}
	}
	return sum
}

func readFile(filename string) [][]pts {
	file, err := os.ReadFile(filename)
	noError(err)
	var rps [][]pts
	for _, val := range strings.Split(string(file), "\n") {
		cms := strings.Split(val, ",")
		p1 := strings.Split(cms[0], "-")
		min, err := strconv.Atoi(p1[0])
		noError(err)
		max, err := strconv.Atoi(p1[1])
		noError(err)
		i1 := pts{min: min, max: max}

		p2 := strings.Split(cms[1], "-")
		min, err = strconv.Atoi(p2[0])
		noError(err)
		max, err = strconv.Atoi(p2[1])
		noError(err)
		i2 := pts{min: min, max: max}

		rps = append(rps, []pts{i1, i2})
	}
	return rps
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
