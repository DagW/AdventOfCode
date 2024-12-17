package main

import (
	"fmt"
	"testing"
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
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if result := part1(tc.input); result != tc.expected {
				t.Fatalf("%d: expected %d should have been %d", i, tc.expected, result)
			}
		})
	}
}

func part1(input string) int {
	return len(input)
}
