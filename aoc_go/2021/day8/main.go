package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func part1(segments []segments) int {
	count := 0
	for _, s := range segments {
		for _, o := range s.output {
			switch len(o) {
			case 3:
				count++
			case 2:
				count++
			case 7:
				count++
			case 4:
				count++
			}
		}
	}
	return count
}

func part2(segments []segments) int {
	sum := 0
	for _, s := range segments {
		m := map[string]int{}
		for _, o := range s.input {
			m[o] = stringfunc(o)
		}
		fmt.Println(m)
		ns := 0
		for i, v2 := range s.output {
			for k, v := range m {
				fmt.Println(k, v2)
				if equalsIgnoreOrder(k, v2) {
					fmt.Println("--->", ns)
					ns += v * int(math.Pow10(i))
					break
				}
			}
		}

		sum += ns
	}
	return sum
}

func stringfunc(o string) int {
	switch len(o) {
	case 3:
		return 7
	case 2:
		return 1
	case 7:
		return 8
	case 4:
		return 4
	}


	/*switch {
	case containsIgnoreOrder("fbcad", o):
		return 3
	case containsIgnoreOrder("cdfbe", o):
		return 5
	case containsIgnoreOrder("cdfgeb", o):
		return 6
	case containsIgnoreOrder("cefabd", o):
		return 9
	case containsIgnoreOrder("gcdfa", o):
		return 2

	case containsIgnoreOrder("cagedb", o):
		return 0

	case containsIgnoreOrder("acedgfb", o):
		return 8

	case containsIgnoreOrder("eafb", o):
		return 4
	case containsIgnoreOrder("dab", o):
		return 7
	case containsIgnoreOrder("ab", o):
		return 1
	}*/
	panic(o)
}

func containsIgnoreOrder(a, b string) bool {
	m := map[rune]struct{}{}
	for _, o := range a {
		m[o] = struct{}{}
	}
	for _, o := range b {
		_, exists := m[o]
		if !exists {
			return false
		}
	}
	return true
}

func equalsIgnoreOrder(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return true
	}
	m := map[rune]struct{}{}
	for _, o := range a {
		m[o] = struct{}{}
	}
	for _, o := range b {
		_, exists := m[o]
		if !exists {
			return false
		}
	}
	return true
}

type segments struct {
	input  []string
	output []string
}

func readFile(filename string) (seg []segments) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, s := range lines {
		var in []string
		var out []string
		parts := strings.Split(s, "|")
		for _, v := range strings.Split(strings.Trim(parts[0], " "), " ") {
			in = append(in, v)
		}
		for _, v := range strings.Split(strings.Trim(parts[1], " "), " ") {
			out = append(out, v)
		}
		seg = append(seg, segments{in, out})
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day8/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day8/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2([]segments{
		{
			[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
		},
	}))
	fmt.Println("Test", part2(readFile("aoc_go/2021/day8/test")))
	//fmt.Println("Input", part2(readFile("aoc_go/2021/day8/input")))
}
