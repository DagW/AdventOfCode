package main

import (
	"testing"
)

func TestRecursionPart1(t *testing.T) {
	var testFiles = []struct {
		data     string
		expected int
	}{
		{readInput("test1"), 31},
		{readInput("test2"), 165},
		{readInput("test3"), 13312},
		{readInput("test4"), 180697},
		{readInput("test5"), 2210736},
	}
	for i, tt := range testFiles {
		if got := SolvePart1(tt.data, findOreForReactionsWithRecursion); got != tt.expected {
			t.Errorf("SolvePart1 recursive %d = %d; wanted %d", i, got, tt.expected)

		}
	}
}

func TestRecursionPart2(t *testing.T) {
	var testFiles = []struct {
		data     string
		expected int
	}{
		{readInput("test3"), 82892753},
		{readInput("test4"), 5586022},
		{readInput("test5"), 460664},
	}
	for i, tt := range testFiles {
		if got := SolvePart2(tt.data, findOreForReactionsWithRecursion); got != tt.expected {
			t.Errorf("SolvePart1 recursive %d = %d; wanted %d", i, got, tt.expected)

		}
	}
}

func TestQueuePart1(t *testing.T) {
	var testFiles = []struct {
		data     string
		expected int
	}{
		{readInput("test1"), 31},
		{readInput("test2"), 165},
		{readInput("test3"), 13312},
		{readInput("test4"), 180697},
		{readInput("test5"), 2210736},
	}
	for i, tt := range testFiles {
		if got := SolvePart1(tt.data, findOreForReactionsWithQueue); got != tt.expected {
			t.Errorf("SolvePart1 recursive %d = %d; wanted %d", i, got, tt.expected)

		}
	}
}

func TestQueuePart2(t *testing.T) {
	var testFiles = []struct {
		data     string
		expected int
	}{
		{readInput("test3"), 82892753},
		{readInput("test4"), 5586022},
		{readInput("test5"), 460664},
	}
	for i, tt := range testFiles {
		if got := SolvePart2(tt.data, findOreForReactionsWithQueue); got != tt.expected {
			t.Errorf("SolvePart1 recursive %d = %d; wanted %d", i, got, tt.expected)

		}
	}
}

func BenchmarkRecursion(b *testing.B) {
	data := readInput("input")
	for n := 0; n < b.N; n++ {
		SolvePart1(data, findOreForReactionsWithRecursion)
	}
}

func BenchmarkQueue(b *testing.B) {
	data := readInput("input")
	for n := 0; n < b.N; n++ {
		SolvePart1(data, findOreForReactionsWithQueue)
	}
}
