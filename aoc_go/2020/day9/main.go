package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) (numbers []int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	return
}

type sums []int

func (s sums) contains(num int) bool {
	for _, n := range s {
		if n == num {
			return true
		}
	}
	return false
}

func getSums(offset int, numbers []int, preambleLength int) sums {
	var sums sums
	for i, n1 := range numbers[offset-preambleLength : offset] {
		for j, n2 := range numbers[offset-preambleLength : offset] {
			if i != j {
				sums = append(sums, n1+n2)
			}
		}
	}
	return sums
}

func part1(numbers []int, preambleLength int) int {
	for i := preambleLength; i < len(numbers); i++ {
		sums := getSums(i, numbers, preambleLength)
		if !sums.contains(numbers[i]) {
			return numbers[i]
		}
	}
	return -1
}

func main() {
	fmt.Println("part1", part1(readFile("aoc_go/2020/day9/test"), 25))
	fmt.Println("part1", part1(readFile("aoc_go/2020/day9/test2"), 5))

	input := readFile("aoc_go/2020/day9/input")
	ans := part1(input, 25)
	fmt.Println("part1", ans)

outer:
	for i := range input {
		sum := 0
		min, max := int(^uint(0)>>1), 0
		for j := i; j < len(input); j++ {
			sum += input[j]
			if input[j] > max {
				max = input[j]
			}
			if input[j] < min {
				min = input[j]
			}
			if sum == ans {
				fmt.Println("part2", min+max)
				break outer
			}
		}
	}

}
