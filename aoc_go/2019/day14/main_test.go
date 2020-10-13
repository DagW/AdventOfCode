package main

import (
	"testing"
)

func TestRecursionPart1(t *testing.T) {
	var testFiles = []struct {
		file     string
		expected int
	}{
		{"test1", 31},
		{"test2", 165},
		{"test3", 13312},
		{"test4", 180697},
		{"test5", 2210736},
	}
	for _, tt := range testFiles {
		if got := SolvePart1(readInput(tt.file), findOreForReactionsWithRecursion); got != tt.expected {
			t.Errorf("SolvePart1 recursive %s = %d; wanted %d", tt.file, got, tt.expected)

		}
	}
}

func TestRecursionPart2(t *testing.T) {
	var testFiles = []struct {
		file     string
		expected int
	}{
		{"test3", 82892753},
		{"test4", 5586022},
		{"test5", 460664},
	}
	for _, tt := range testFiles {
		if got := SolvePart2(readInput(tt.file), findOreForReactionsWithRecursion); got != tt.expected {
			t.Errorf("SolvePart1 recursive %s = %d; wanted %d", tt.file, got, tt.expected)

		}
	}
}

func TestQueuePart1(t *testing.T) {
	var testFiles = []struct {
		file     string
		expected int
	}{
		{"test1", 31},
		{"test2", 165},
		{"test3", 13312},
		{"test4", 180697},
		{"test5", 2210736},
	}
	for _, tt := range testFiles {
		if got := SolvePart1(readInput(tt.file), findOreForReactionsWithQueue); got != tt.expected {
			t.Errorf("SolvePart1 recursive %s = %d; wanted %d", tt.file, got, tt.expected)

		}
	}
}

func TestQueuePart2(t *testing.T) {
	var testFiles = []struct {
		file     string
		expected int
	}{
		{"test3", 82892753},
		{"test4", 5586022},
		{"test5", 460664},
	}
	for _, tt := range testFiles {
		if got := SolvePart2(readInput(tt.file), findOreForReactionsWithQueue); got != tt.expected {
			t.Errorf("SolvePart1 recursive %s = %d; wanted %d", tt.file, got, tt.expected)

		}
	}
}

func BenchmarkRecursion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolvePart1(readInput("input"), findOreForReactionsWithRecursion)
	}
}

func BenchmarkQueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SolvePart1(readInput("input"), findOreForReactionsWithQueue)
	}
}
