package day16

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/DagW/AdventOfCode/common"
	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	tcs := []struct {
		input    string
		expected int
	}{
		{
			input: `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
			expected: 7036,
		},
		{
			input: `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`,
			expected: 11048,
		},
		{
			input:    string(common.Must(os.ReadFile("input"))),
			expected: 11048,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			require.Equal(t, tc.expected, part1(tc.input))
		})
	}
}

const (
	wall  rune = '#'
	start rune = 'S'
	end   rune = 'E'
)

type chart struct {
	m          [][]rune
	start, end position
}
type position [2]int
type state struct {
	current position
	visited []position
	path    string
}

func parse(input string) chart {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")

	c := chart{m: make([][]rune, len(rows))}
	for i, row := range rows {
		c.m[i] = make([]rune, len(row))
	}

	for y, row := range rows {
		for x, col := range row {
			c.m[y][x] = col
			if col == start {
				c.start[0], c.start[1] = y, x
			} else if col == end {
				c.end[0], c.end[1] = y, x
			}
		}
	}
	return c
}

func (c chart) print(s state) {
	for y, r := range c.m {
		for x, c := range r {
			if slices.Contains(s.visited, position{y, x}) {
				fmt.Print("x")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func (c chart) done(current position) bool {
	return current[0] == c.end[0] && current[1] == c.end[1]
}

func Dijkstras(c *chart) string {
	type result struct {
		score int
		path  string
	}

	visited := map[position]result{}
	for y, row := range c.m {
		for x := range row {
			visited[position{y, x}] = result{score: math.MaxInt32}
		}
	}

	pop := func(m map[position]result) (position, result) {
		// pops the lowest score
		lowest := result{score: math.MaxInt32}
		lowestKey := position{}
		for k, v := range m {
			if v.score < lowest.score {
				lowest = v
				lowestKey = k
			}
		}
		delete(m, lowestKey)
		return lowestKey, lowest
	}

	toVisit := map[position]result{c.start: {score: 0, path: "r"}}
	for len(toVisit) > 0 {
		current, currentResult := pop(toVisit)

		// Add current position to visited
		visited[current] = result{score: currentResult.score, path: currentResult.path}

		// Search the neighbours
		neighbours := map[string]position{
			"l": {current[0], current[1] - 1},
			"r": {current[0], current[1] + 1},
			"u": {current[0] - 1, current[1]},
			"d": {current[0] + 1, current[1]},
		}
		for key, neighbour := range neighbours {
			sign := c.m[neighbour[0]][neighbour[1]]
			if sign != wall {
				newPath := currentResult.path + key
				newScore := score(newPath)
				if r, alreadyVisited := visited[neighbour]; alreadyVisited {
					if newScore < r.score {
						// If we have a better score this pass, search it again
						toVisit[neighbour] = result{score: newScore, path: newPath}
					}
				} else {
					toVisit[neighbour] = result{score: newScore, path: newPath}
				}
			}
		}
	}

	return visited[c.end].path
}

func (s state) Dfs(c *chart, d int) (state, bool) {
	s.visited = append(s.visited, s.current)

	if c.done(s.current) {
		return s, true
	}

	var results []state
	neighbours := map[string]position{
		"l": {s.current[0], s.current[1] - 1},
		"r": {s.current[0], s.current[1] + 1},
		"u": {s.current[0] - 1, s.current[1]},
		"d": {s.current[0] + 1, s.current[1]},
	}
	for key, neighbour := range neighbours {
		sign := c.m[neighbour[0]][neighbour[1]]
		if sign != wall && !slices.Contains(s.visited, neighbour) {
			opt := s
			opt.current = neighbour
			opt.path += key
			if state, ok := opt.Dfs(c, d+1); ok {
				results = append(results, state)
			}
		}
	}

	return minState(results...)
}

func (s state) Bfs(c *chart, d int) (state, bool) {
	s.visited = append(s.visited, s.current)

	if c.done(s.current) {
		return s, true
	}

	neighbours := map[string]position{
		"l": {s.current[0], s.current[1] - 1},
		"r": {s.current[0], s.current[1] + 1},
		"u": {s.current[0] - 1, s.current[1]},
		"d": {s.current[0] + 1, s.current[1]},
	}
	var opts []state
	for key, neighbour := range neighbours {
		sign := c.m[neighbour[0]][neighbour[1]]
		if sign != wall && !slices.Contains(s.visited, neighbour) {
			opt := s
			opt.current = neighbour
			opt.path += key
			opts = append(opts, opt)
		}
	}

	var results []state
	for _, opt := range opts {
		if state, ok := opt.Bfs(c, d+1); ok {
			results = append(results, state)
		}
	}

	return minState(results...)
}

func minState(results ...state) (state, bool) {
	if len(results) < 1 {
		return state{}, false
	}
	best := results[0]
	for _, v := range results {
		if score(v.path) <= score(best.path) {
			best = v
		}
	}
	return best, true
}

func score(path string) int {
	turns := 0
	last := rune(path[0])
	for _, c := range path {
		if c != last {
			turns++
			last = c
		}
	}
	return len(path) - 1 + (turns * 1000)
}

func part1(input string) int {
	chart := parse(input)

	/*start := time.Now()
	s := state{current: chart.start, path: "r"}
	r, _ := s.Dfs(&chart, 0)
	fmt.Println("DFS: ", time.Since(start), score(r.path))
	chart.print(r)

	start = time.Now()
	s = state{current: chart.start, path: "r"}
	r, _ = s.Bfs(&chart, 0)
	fmt.Println("DFS: ", time.Since(start), score(r.path))
	chart.print(r)*/

	start := time.Now()
	path := Dijkstras(&chart)
	fmt.Println("Dijkstras: ", time.Since(start), score(path))

	return score(path)
}
