package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(input [][]rune) int {
	var prio = map[rune]int{}
	for i, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		prio[r] = i + 1
	}

	sum := 0
	for _, backpack := range input {
		l, r := backpack[0:len(backpack)/2], backpack[len(backpack)/2:]
		fm := map[rune]struct{}{}
		for _, v := range l {
			fm[v] = struct{}{}
		}
		for _, v := range r {
			if _, ok := fm[v]; ok {
				sum += prio[v]
				break
			}
		}
	}
	return sum
}

func deduplicate(in []rune) (out []rune) {
	d := map[rune]struct{}{}
	for _, r := range in {
		if _, exists := d[r]; !exists {
			d[r] = struct{}{}
			out = append(out, r)
		}
	}
	return out
}

func part2(input [][]rune) int {
	var prio = map[rune]int{}
	for i, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		prio[r] = i + 1
	}

	sum := 0
	for i := 0; i < len(input); i += 3 {
		fm := map[rune]int{}

		for _, rs := range [][]rune{
			deduplicate(input[i+0]),
			deduplicate(input[i+1]),
			deduplicate(input[i+2]),
		} {
			for _, v := range rs {
				fm[v] = fm[v] + 1
			}
		}

		for k, v := range fm {
			if v == 3 {
				sum += prio[k]
			}
		}
	}
	return sum
}

func readFile(filename string) [][]rune {
	file, err := os.ReadFile(filename)
	noError(err)
	var rps [][]rune
	for _, val := range strings.Split(string(file), "\n") {
		var r []rune
		for _, v := range val {
			r = append(r, v)
		}
		rps = append(rps, r)
	}
	return rps
}

func main() {
	fmt.Println("Part1", part1(readFile("test")))
	fmt.Println("Part1", part1(readFile("input")))
	fmt.Println("Part2", part2(readFile("test")))
	fmt.Println("Part2", part2(readFile("input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
