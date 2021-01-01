package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	amount   int
	contains []bag
}

func part1(bags []bag) int {
	toVisit := []string{"shiny gold"}
	results := map[string]struct{}{}
	for len(toVisit) > 0 {
		var popped string
		popped, toVisit = toVisit[0], toVisit[1:]
		for _, bag := range bags {
			for _, c := range bag.contains {
				if c.color == popped {
					results[bag.color] = struct{}{}
					toVisit = append(toVisit, bag.color)
				}
			}
		}
	}
	return len(results)
}

func part2(bags []bag) int {
	num := 0
	toVisit := []bag{{
		color:    "shiny gold",
		amount:   1,
		contains: nil,
	}}
	for len(toVisit) > 0 {
		var popped bag
		popped, toVisit = toVisit[0], toVisit[1:]
		for _, b := range bags {
			if b.color == popped.color {
				for _, c := range b.contains {
					c.amount *= popped.amount
					toVisit = append(toVisit, c)
				}
				num += popped.amount
			}
		}
	}
	// And remove the gold bag
	return num - 1
}

func readFile(filename string) []bag {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	r := regexp.MustCompile(`(\d+) (\w+ \w+) bag.?.?`)
	var bags []bag
	for _, line := range lines {
		pieces := strings.Split(line, "contain")
		var contains []bag
		for _, s := range strings.Split(pieces[1], ",") {
			if s != " no other bags." {
				matches := r.FindStringSubmatch(s)
				if len(matches) > 0 {
					amount, _ := strconv.Atoi(matches[1])
					contains = append(contains, bag{
						color:    strings.Trim(matches[2], " ,"),
						amount:   amount,
						contains: nil,
					})
				}
			}
		}
		bags = append(bags, bag{
			color:    strings.Trim(strings.Split(pieces[0], " bag")[0], " ,"),
			amount:   1,
			contains: contains,
		})
	}
	return bags
}

func main() {
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day7/test")))
	fmt.Println("Part1", part1(readFile("aoc_go/2020/day7/input")))
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day7/test")))
	fmt.Println("Part2", part2(readFile("aoc_go/2020/day7/input")))
}
