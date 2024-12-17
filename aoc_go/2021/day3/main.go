package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2(values [][]bool) (powerConsumption int) {
	oxygen := calcBits(calculateValue(values, false), false)
	co2scrubber := calcBits(calculateValue(values, true), false)

	return oxygen * co2scrubber
}

func calculateValue(values [][]bool, inverse bool) []bool {
	tmp := values
	for i := range values {
		if len(tmp) == 1 {
			break
		}
		tmp = reduceLeastCommon(i, tmp, inverse)
	}
	return tmp[0]
}

// reduceLeastCommon removes the arrays with the least common bit in position i
func reduceLeastCommon(i int, values [][]bool, inverse bool) (newValues [][]bool) {
	ones := 0
	for _, v := range values {
		if v[i] {
			ones++
		}
	}
	onesMostCommon := ones > (len(values)-1)/2
	if inverse {
		onesMostCommon = !onesMostCommon
	}
	for _, v := range values {
		if v[i] == onesMostCommon {
			newValues = append(newValues, v)
		}
	}
	return newValues
}

func calcBits(bs []bool, inverse bool) int {
	sum, valuePerBit := 0, 1
	for i := range bs {
		b := bs[len(bs)-1-i]
		if inverse {
			b = !b
		}
		if b {
			sum += valuePerBit
		}
		valuePerBit *= 2
	}
	return sum
}

func part1(values [][]bool) (powerConsumption int) {
	var bits []bool
	for i := 0; i < len(values[0]); i++ {
		count := 0
		for j := 0; j < len(values); j++ {
			if values[j][i] {
				count++
			}
		}
		bits = append(bits, count > len(values)/2)
	}
	return calcBits(bits, false) * calcBits(bits, true)
}

func readFile(filename string) (values [][]bool) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		var bits []bool
		for _, c := range line {
			bits = append(bits, c == '1')
		}
		values = append(values, bits)
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day3/test")))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day3/input")))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day3/test")))
	fmt.Println("input", part2(readFile("aoc_go/2021/day3/input")))
}
