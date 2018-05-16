package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("data")
	check(err)

	fmt.Println(string(dat))
	fmt.Println(SolvePart1(string(dat)))
	fmt.Println(SolvePart2(string(dat)))
}

func SolvePart1(input string) int {
	sum := 0
	for pos, char := range input {
		nextpos := (pos + 1) % len(input)
		if string(input[pos]) == string(input[nextpos]) {
			i, err := strconv.Atoi(string(char))
			check(err)
			sum += i
		}
	}
	return sum
}

func SolvePart2(input string) int {
	sum := 0
	for pos, char := range input {
		nextpos := (pos + len(input)/2) % len(input)
		if string(input[pos]) == string(input[nextpos]) {
			i, err := strconv.Atoi(string(char))
			check(err)
			sum += i
		}
	}
	return sum
}
