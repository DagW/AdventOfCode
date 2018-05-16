package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	values := [][]string{
		{"5", "1", "9", "5"},
		{"7", "5", "3"},
		{"2", "4", "6", "8"},
	}
	assert.Equal(t, 18, SolvePart1(values))
}

func TestSolvePart2(t *testing.T) {
	values := [][]string{
		{"5", "9", "2", "8"},
		{"9", "4", "7", "3"},
		{"3", "8", "6", "5"},
	}
	assert.Equal(t, 9, SolvePart2(values))
}
