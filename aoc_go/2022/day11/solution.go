package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	num       int
	items     []int
	operation string
	ifTrueTo  int
	ifFalseTo int
	test      int
}

func its(s string) int {
	i, err := strconv.Atoi(strings.Trim(s, " "))
	noError(err)
	return i
}

func part1(input []monkey) int {
	counter := map[int]int{}
	for i := 0; i < 20; i++ {
		for monkeyIndex, m := range input {
			for _, item := range m.items {
				n := strings.ReplaceAll(m.operation, "old", fmt.Sprintf("%d", item))
				pieces := strings.Split(n, " ")
				nval := 0
				if pieces[1] == "*" {
					nval = its(pieces[0]) * its(pieces[2])
				} else if pieces[1] == "+" {
					nval = its(pieces[0]) + its(pieces[2])
				}
				nval /= 3

				if nval%m.test == 0 {
					input[m.ifTrueTo].items = append(input[m.ifTrueTo].items, nval)
				} else {
					input[m.ifFalseTo].items = append(input[m.ifFalseTo].items, nval)
				}
				counter[monkeyIndex]++
			}
			input[monkeyIndex].items = []int{}
		}
	}
	var monkeyBusiness []int
	for i, m := range input {
		fmt.Println(i, m.items, "inspected", counter[i])
		monkeyBusiness = append(monkeyBusiness, counter[i])
	}
	sort.Ints(monkeyBusiness)

	return monkeyBusiness[len(monkeyBusiness)-1] * monkeyBusiness[len(monkeyBusiness)-2]
}

func part2(input []monkey) int {
	allMods := 1
	for _, m := range input {
		allMods *= m.test
	}
	counter := map[int]int{}
	for i := 0; i < 10000; i++ {
		for monkeyIndex, m := range input {
			for _, item := range m.items {
				n := strings.ReplaceAll(m.operation, "old", fmt.Sprintf("%d", item))
				pieces := strings.Split(n, " ")
				nval := 0
				if pieces[1] == "*" {
					nval = its(pieces[0]) * its(pieces[2])
				} else if pieces[1] == "+" {
					nval = its(pieces[0]) + its(pieces[2])
				}

				nval %= allMods

				if nval%m.test == 0 {
					input[m.ifTrueTo].items = append(input[m.ifTrueTo].items, nval)
				} else {
					input[m.ifFalseTo].items = append(input[m.ifFalseTo].items, nval)
				}
				counter[monkeyIndex]++
			}
			input[monkeyIndex].items = []int{}
		}
	}
	var monkeyBusiness []int
	for i, m := range input {
		fmt.Println(i, m.items, "inspected", counter[i])
		monkeyBusiness = append(monkeyBusiness, counter[i])
	}
	sort.Ints(monkeyBusiness)

	return monkeyBusiness[len(monkeyBusiness)-1] * monkeyBusiness[len(monkeyBusiness)-2]
}

func readFile(filename string) []monkey {
	file, err := os.ReadFile(filename)
	noError(err)
	var monkeys []monkey

	var currentMonkey monkey
	for _, val := range strings.Split(string(file), "\n") {
		if strings.HasPrefix(val, "Monkey ") {
			// Start a new monkey
			noError(err)
			currentMonkey.num = its(strings.ReplaceAll(strings.Split(val, "Monkey ")[1], ":", ""))
		} else if val == "" {
			// End the current monkey
			monkeys = append(monkeys, currentMonkey)
			currentMonkey = monkey{}
		} else if strings.Contains(val, "Starting items:") {
			items := strings.Split(strings.Split(val, "items: ")[1], ",")
			for _, item := range items {
				currentMonkey.items = append(currentMonkey.items, its(item))
			}
		} else if strings.Contains(val, "Operation:") {
			op := strings.Split(val, "Operation: new = ")[1]
			currentMonkey.operation = op
		} else if strings.Contains(val, "Test:") {
			op := strings.Split(val, "Test: divisible by")[1]
			currentMonkey.test = its(op)
		} else if strings.Contains(val, "If true:") {
			op := strings.Split(val, "If true: throw to monkey ")[1]
			currentMonkey.ifTrueTo = its(op)
		} else if strings.Contains(val, "If false:") {
			op := strings.Split(val, "If false: throw to monkey ")[1]
			currentMonkey.ifFalseTo = its(op)
		}
	}
	return monkeys
}

func main() {
	fmt.Println(part1(readFile("aoc_go/2022/day11/test")))
	fmt.Println(part1(readFile("aoc_go/2022/day11/input")))
	fmt.Println(part2(readFile("aoc_go/2022/day11/test")))
	fmt.Println(part2(readFile("aoc_go/2022/day11/input")))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
