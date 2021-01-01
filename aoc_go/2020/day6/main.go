package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) [][]string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var groups [][]string
	var group []string
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups, group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)
	return groups
}

type empty struct{}

func part1(groups [][]string) (sum int) {
	for _, group := range groups {
		var groupSet = make(map[rune]empty)
		for _, answers := range group {
			for _, c := range answers {
				groupSet[c] = empty{}
			}
		}
		sum += len(groupSet)
	}
	return
}

func part2(groups [][]string) (sum int) {
	for _, group := range groups {
		var groupSet = make(map[rune]int)
		for _, answers := range group {
			for _, c := range answers {
				groupSet[c] = groupSet[c] + 1
			}
		}
		for _, v := range groupSet {
			if v == len(group) {
				sum += 1
			}
		}
	}
	return
}
func main() {
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day6/test")))
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day6/input")))
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day6/test")))
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day6/input")))
}
