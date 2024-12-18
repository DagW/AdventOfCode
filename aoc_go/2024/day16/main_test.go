package day16

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/DagW/AdventOfCode/common"
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
			if result := part1(tc.input); result != tc.expected {
				t.Fatalf("%d: expected %d should have been %d", i, tc.expected, result)
			}
		})
	}
}

const (
	wall  rune = '#'
	start rune = 'S'
	end   rune = 'E'
)

type position [2]int
type state struct {
	points              int
	direction           int
	start, end, current position
	m                   [][]rune
	visited             []position
}

func parse(input string) state {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")

	s := state{m: make([][]rune, len(rows))}
	for i, row := range rows {
		s.m[i] = make([]rune, len(row))
	}

	for y, row := range rows {
		for x, col := range row {
			s.m[y][x] = col
			if col == start {
				s.start[0], s.start[1] = y, x
			} else if col == end {
				s.end[0], s.end[1] = y, x
			}
		}
	}
	s.current = s.start
	s.points = 0
	return s
}

func (s state) print() {
	for y, r := range s.m {
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

func (s state) done() bool {
	return s.current[0] == s.end[0] && s.current[1] == s.current[1]
}

func (s state) walk(d int) state {
	s.visited = append(s.visited, s.current)

	if s.m[s.current[0]][s.current[1]] == end {
		return s
	}

	var results []state

	if s.current[1] > 0 {
		left := s.m[s.current[0]][s.current[1]-1]
		pos := position{s.current[0], s.current[1] - 1}
		if left != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.direction = 0
			if s.direction != opt.direction {
				opt.points += 1000
			} else {
				opt.points += 1
			}
			state := opt.walk(d + 1)
			if state.points != -1 {
				results = append(results, state)
			}
		}
	}
	if s.current[1] < len(s.m[s.current[0]]) {
		right := s.m[s.current[0]][s.current[1]+1]
		pos := position{s.current[0], s.current[1] + 1}
		if right != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.direction = 1
			if s.direction != opt.direction {
				opt.points += 1000
			} else {
				opt.points += 1
			}
			state := opt.walk(d + 1)
			if state.points != -1 {
				results = append(results, state)
			}
		}
	}
	if s.current[0] > 0 {
		up := s.m[s.current[0]-1][s.current[1]]
		pos := position{s.current[0] - 1, s.current[1]}
		if up != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.direction = 2
			if s.direction != opt.direction {
				opt.points += 1000
			} else {
				opt.points += 1
			}
			state := opt.walk(d + 1)
			if state.points != -1 {
				results = append(results, state)
			}
		}
	}
	if s.current[0] < len(s.m) {
		down := s.m[s.current[0]+1][s.current[1]]
		pos := position{s.current[0] + 1, s.current[1]}
		if down != wall && !slices.Contains(s.visited, pos) {
			opt := s
			opt.current = pos
			opt.direction = 3
			if s.direction != opt.direction {
				opt.points += 1000
			} else {
				opt.points += 1
			}
			state := opt.walk(d + 1)
			if state.points != -1 {
				results = append(results, state)
			}
		}
	}

	if len(results) < 1 {
		return state{points: -1}
	}

	m := results[0]
	for _, v := range results {
		if v.points < m.points {
			m = v
		}
	}
	return m
}

func part1(input string) int {
	s := parse(input)
	r := s.walk(0)
	r.print()
	fmt.Println("SCORE", r.points)
	return r.points
}
