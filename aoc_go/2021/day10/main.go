package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(filename string) (m [][]int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, s := range lines {
		fmt.Println("is corrupt?", s, isCorrupted(s))
	}
	return
}

func isCorrupted(s string) bool {
	curly, bracket, croc, parenthesis := 0, 0, 0, 0
	for _, c := range s {
		// fmt.Println(fmt.Sprintf("%c", c))
		switch c {
		case '{':
			curly++
		case '}':
			curly--
		case '(':
			parenthesis++
		case ')':
			parenthesis--
		case '<':
			croc++
		case '>':
			croc--
		case '[':
			bracket++
		case ']':
			bracket--
		}
		// fmt.Println(curly, bracket, croc, parenthesis)
	}
	return curly < 0 || bracket < 0 || croc < 0 || parenthesis < 0
}

func main() {
	fmt.Println("Part1")
	//fmt.Println("Test", recurse("<([]){()}[{}])"))
	fmt.Println("Test", part1("aoc_go/2021/day10/test"))
	// fmt.Println("Input", part1(readFile("aoc_go/2021/day9/input")))
}
