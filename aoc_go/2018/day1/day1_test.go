package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 3, SolvePart1("1122"))
	assert.Equal(t, 4, SolvePart1("1111"))
	assert.Equal(t, 0, SolvePart1("1234"))
	assert.Equal(t, 9, SolvePart1("91212129"))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 6, SolvePart2("1212"))
	assert.Equal(t, 0, SolvePart2("1221"))
	assert.Equal(t, 4, SolvePart2("123425"))
	assert.Equal(t, 12, SolvePart2("123123"))
	assert.Equal(t, 4, SolvePart2("12131415"))
}
