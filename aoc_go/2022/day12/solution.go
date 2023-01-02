package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type pos struct {
	x, y int
}

func part1(input [][]rune) int {
	start := find(input, 'S')
	goal := find(input, 'E')
	input[start.y][start.x] = 'a'
	input[goal.y][goal.x] = 'z'

	visited := bfs(input, start)
	return visited[goal]
}

func bfs(input [][]rune, start pos) map[pos]int {
	var prio = map[rune]int{}
	for i, r := range "abcdefghijklmnopqrstuvwxyz" {
		prio[r] = i
	}

	visited := map[pos]int{}
	toVisit := map[pos]int{start: 0}
	for len(toVisit) > 0 {
		pos, steps := pop(toVisit)
		visited[pos] = steps

		for _, n := range getNeighbours(pos, len(input), len(input[0])) {
			if prio[input[n.y][n.x]] > prio[input[pos.y][pos.x]]+1 {
				continue
			}
			prevSteps, alreadyVisited := visited[n]
			if alreadyVisited && prevSteps < steps+1 {
				continue // This path is not better
			}
			// If its already in the list of things to visit, only add it if its better
			if tv, alreadyInVisiting := toVisit[n]; alreadyInVisiting {
				if tv > steps+1 {
					toVisit[n] = steps + 1
				}
			} else {
				toVisit[n] = steps + 1
			}
		}
	}

	return visited
}

func printScoreMap(input [][]rune, visited map[pos]int) {
	fmt.Println()
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			fmt.Print(fmt.Sprintf(" %d ", visited[pos{x: x, y: y}]))
		}
		fmt.Println()
	}
}

func part2(input [][]rune) int {
	start := find(input, 'S')
	input[start.y][start.x] = 'a'
	goal := find(input, 'E')
	input[goal.y][goal.x] = 'z'

	var results []int
	for _, start := range findAll(input, 'a') {

		visited := bfs(input, start)

		if x, has := visited[goal]; has {
			results = append(results, x)
		}
	}
	sort.Ints(results)
	fmt.Println(results)
	return results[0]
}

func pop(m map[pos]int) (pos, int) {
	for k, v := range m {
		delete(m, k)
		return k, v
	}
	noError(fmt.Errorf("toVisit should have contained an item"))
	return pos{}, 0
}

func getNeighbours(p pos, ymax, xmax int) (n []pos) {
	for _, neighbour := range []pos{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	} {
		npos := pos{x: p.x + neighbour.x, y: p.y + neighbour.y}
		if npos.y < 0 || npos.x < 0 || npos.x > xmax-1 || npos.y > ymax-1 {
			continue
		}
		n = append(n, npos)
	}
	return n
}

func find(input [][]rune, s rune) pos {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == s {
				return pos{y: y, x: x}
			}
		}
	}
	noError(fmt.Errorf("could not find wanted postition"))
	return pos{}
}
func findAll(input [][]rune, s rune) (r []pos) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == s {
				r = append(r, pos{y: y, x: x})
			}
		}
	}
	return r
}

func readFile(filename string) (output [][]rune) {
	file, err := os.ReadFile(filename)
	noError(err)
	for _, val := range strings.Split(string(file), "\n") {
		var row []rune
		for _, r := range val {
			row = append(row, r)
		}
		output = append(output, row)
	}
	return output
}

func main() {
	fmt.Println(part1(readFile("aoc_go/2022/day12/test")))
	fmt.Println(part1(readFile("aoc_go/2022/day12/input")))
	fmt.Println(part2(readFile("aoc_go/2022/day12/test")))
	fmt.Println(part2(readFile("aoc_go/2022/day12/input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
