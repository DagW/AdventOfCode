package day16

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

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
	direction int
	current   position
	visited   []position
	decision  string
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

func (s state) walk(c *chart, d int) (state, bool) {
	s.visited = append(s.visited, s.current)

	if c.done(s.current) {
		return s, true
	}

	var results []state

	if s.current[1] > 0 {
		left := c.m[s.current[0]][s.current[1]-1]
		pos := position{s.current[0], s.current[1] - 1}
		if left != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.decision += "l"
			if state, ok := opt.walk(c, d+1); ok {
				results = append(results, state)
			}
		}
	}
	if s.current[1] < len(c.m[s.current[0]]) {
		right := c.m[s.current[0]][s.current[1]+1]
		pos := position{s.current[0], s.current[1] + 1}
		if right != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.decision += "r"
			if state, ok := opt.walk(c, d+1); ok {
				results = append(results, state)
			}
		}
	}
	if s.current[0] > 0 {
		up := c.m[s.current[0]-1][s.current[1]]
		pos := position{s.current[0] - 1, s.current[1]}
		if up != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.decision += "u"
			if state, ok := opt.walk(c, d+1); ok {
				results = append(results, state)
			}
		}
	}
	if s.current[0] < len(c.m) {
		down := c.m[s.current[0]+1][s.current[1]]
		pos := position{s.current[0] + 1, s.current[1]}
		if down != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.decision += "d"
			if state, ok := opt.walk(c, d+1); ok {
				results = append(results, state)
			}
		}
	}

	return func(s ...state) (state, bool) {
		if len(s) < 1 {
			return state{}, false
		}
		best := s[0]
		for _, v := range s {
			if v.score() <= best.score() {
				best = v
			}
		}
		return best, true
	}(results...)
}

func (s state) score() int {
	turns := 1
	last := rune(s.decision[0])
	for _, c := range s.decision {
		if c != last {
			turns++
			last = c
		}
	}
	return len(s.decision) + (turns * 1000)
}

func part1(input string) int {
	chart := parse(input)

	s := state{current: chart.start}
	r, ok := s.walk(&chart, 0)
	fmt.Println("SCORE", ok, r.score())
	chart.print(r)
	return r.score()
}
