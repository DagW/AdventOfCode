package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) (m map[string][]string) {
	m = map[string][]string{}
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, s := range lines {
		parts := strings.Split(s, "-")
		m[parts[0]] = append(m[parts[0]], parts[1])
		m[parts[1]] = append(m[parts[1]], parts[0])
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", solve(readFile("aoc_go/2021/day12/test"), part1))
	fmt.Println("Input", solve(readFile("aoc_go/2021/day12/input"), part1))
	fmt.Println("Part2")
	fmt.Println("Test", solve(readFile("aoc_go/2021/day12/test"), part2))
	fmt.Println("Input", solve(readFile("aoc_go/2021/day12/input"), part2))
}

func solve(file map[string][]string, visitableFunc func(target string, path []string) bool) int {

	start, end := "start", "end"

	var current []string
	var pathQueue [][]string
	completed := map[string][]string{}
	pathQueue = append(pathQueue, []string{start})

	for len(pathQueue) > 0 {
		current = make([]string, 0, len(pathQueue[0]))
		current = append(current, pathQueue[0]...)
		pathQueue = pathQueue[1:]

		last := current[len(current)-1]
		if last == end {
			completed[strings.Join(current, ",")] = current
			continue
		}

		for _, v := range file[last] {
			if visitableFunc(v, current) {
				pathQueue = append(pathQueue, append(current, v))
			}
		}
	}

	/*for k := range completed {
		fmt.Println(k)
	}*/

	return len(completed)
}

func part1(target string, path []string) bool {
	if target == strings.ToUpper(target) {
		return true
	} else {
		for _, p := range path {
			if p == target {
				return false
			}
		}
	}
	return true
}

func part2(target string, path []string) (success bool) {
	if target == "start" {
		return false
	} else if target == strings.ToUpper(target) {
		return true
	}

	visits := map[string]int{}
	for _, p := range path {
		visits[p]++
	}
	times, exists := visits[target]

	if !exists || times == 0 {
		// You can always visit it the first time
		return true
	} else if times >= 2 {
		// Never more than twice
		return false
	} else if times == 1 {
		// you can visit ONE small cave TWICE
		anotherCaveHasBeenVisitedTwice := false
		for k, v := range visits {
			if v > 1 && k != strings.ToUpper(k) {
				anotherCaveHasBeenVisitedTwice = true
				break
			}
		}
		if anotherCaveHasBeenVisitedTwice {
			return false
		}
	}

	return true
}
