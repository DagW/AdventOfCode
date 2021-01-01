package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type policyPwd struct {
	minTimes, maxTimes int
	letter             rune
	text               string
}

func (p policyPwd) part1() bool {
	count := 0
	for _, c := range p.text {
		if c == p.letter {
			count++
		}
	}
	if count < p.minTimes || count > p.maxTimes {
		return false
	}
	return true
}

func (p policyPwd) part2() bool {
	if rune(p.text[p.minTimes-1]) == p.letter && rune(p.text[p.maxTimes-1]) != p.letter {
		return true
	}
	if rune(p.text[p.minTimes-1]) != p.letter && rune(p.text[p.maxTimes-1]) == p.letter {
		return true
	}
	return false
}

func validate(input []policyPwd, f func(p policyPwd) bool) int {
	numValid := 0
	for _, p := range input {
		if f(p) {
			numValid++
		}
	}
	return numValid
}

func readFile(filename string) []policyPwd {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var program []policyPwd
	r := regexp.MustCompile(`^(?P<min>[\d]+)-(?P<max>[\d]+) (?P<letter>[a-z]): (?P<pwd>[\w]+)$`)
	for _, val := range strings.Split(string(file), "\n") {
		matches := r.FindStringSubmatch(val)
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		ppwd := policyPwd{
			minTimes: min,
			maxTimes: max,
			letter:   []rune(matches[3])[0],
			text:     matches[4],
		}
		program = append(program, ppwd)
	}
	return program
}

func main() {
	fmt.Println("Part1", validate([]policyPwd{
		{1, 2, 'l', "hello"},
		{1, 3, 'a', "abcde"},
		{1, 2, 'b', "cdefg"},
		{2, 9, 'c', "ccccccccc"},
	}, policyPwd.part1))
	fmt.Println("Part1", validate(readFile("aoc_go/2020/day2/day3.input"),
		policyPwd.part1))

	fmt.Println("Part2", validate([]policyPwd{
		{1, 2, 'l', "hello"},
		{1, 3, 'a', "abcde"},
		{1, 2, 'b', "cdefg"},
		{2, 9, 'c', "ccccccccc"},
	}, policyPwd.part2))
	fmt.Println("Part2", validate(readFile("aoc_go/2020/day2/day3.input"),
		policyPwd.part2))
}
