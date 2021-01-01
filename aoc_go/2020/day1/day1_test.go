package main

import "testing"

func TestPart1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	got := part1(input)
	want := 514579
	if got != want {
		t.Errorf("part1() = %d; want %d", got, want)
	}
}
func TestPart2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	got := part2(input)
	want := 241861950
	if got != want {
		t.Errorf("part2() = %d; want %d", got, want)
	}
}