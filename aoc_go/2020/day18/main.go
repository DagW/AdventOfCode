package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) (lines []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

func main() {
	sum := 0
	for _, m := range readFile("aoc_go/2020/day18/test") {
		sum += part1(m)
	}
	fmt.Println("Part1 test:", sum)
	sum = 0
	for _, m := range readFile("aoc_go/2020/day18/input") {
		sum += part1(m)
	}
	fmt.Println("Part1 input:", sum)
	sum = 0
	for _, m := range readFile("aoc_go/2020/day18/test") {
		sum += part2(m)
	}
	fmt.Println("Part2 test:", sum)
	sum = 0
	for _, m := range readFile("aoc_go/2020/day18/input") {
		sum += part2(m)
	}
	fmt.Println("Part2 input:", sum)
}

func part1(m string) int {
	for strings.Contains(m, "(") {
		// Get the deepest parenthesis
		parts := strings.Split(m, "(")
		last := parts[len(parts)-1]
		last = strings.Split(last, ")")[0]
		m = strings.ReplaceAll(m, "("+last+")", strconv.Itoa(ltrSolve(last)))
	}
	return ltrSolve(m)
}

func ltrSolve(m string) int {
	res := 0
	operator := "+"
	for i, p := range strings.Split(m, " ") {
		if i == 0 {
			v, _ := strconv.Atoi(p)
			res += v
		} else {
			if p == "+" {
				operator = "+"
			} else if p == "*" {
				operator = "*"
			} else {
				if operator == "+" {
					v, _ := strconv.Atoi(p)
					res += v
				} else if operator == "*" {
					v, _ := strconv.Atoi(p)
					res *= v
				}
			}
		}
	}
	//fmt.Println("  Solve", m, "=", res)
	return res
}

func part2(m string) int {
	//fmt.Println("-----", m)
	for strings.Contains(m, "(") {
		// Get the deepest parenthesis
		parts := strings.Split(m, "(")
		last := parts[len(parts)-1]
		last = strings.Split(last, ")")[0]
		m = strings.ReplaceAll(m, "("+last+")", strconv.Itoa(precSolve(last)))
	}
	return precSolve(m)
}

func precSolve(m string) int {
	res := 0
	//fmt.Println("---", m)
	for strings.Contains(m, "+") {
		// Get plus
		parts := strings.Split(m, "+")
		left := strings.Split(strings.Trim(parts[0], " "), " ")
		leftplus := left[len(left)-1]

		right := strings.Split(strings.Trim(parts[1], " "), " ")
		rightplus := right[0]
		m = strings.Replace(m, leftplus+" + "+rightplus+"", strconv.Itoa(ltrSolve(leftplus+" + "+rightplus)), 1)

	}
	res += ltrSolve(m)
	//fmt.Println(res)
	return res
}
