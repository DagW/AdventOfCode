package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func transform(sn int, loops int) int {
	value := 1
	for i := 0; i < loops; i++ {
		value *= sn
		value %= 20201227
	}
	return value
}

func loop(sn int, input int) int {
	value := 1
	loop := 0
	for value != input {
		value *= sn
		value %= 20201227
		loop++
	}
	return loop
}

func readFile(filename string) (values []int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		values = append(values, v)
	}
	return
}

func main() {
	sn := 7
	values := readFile("aoc_go/2020/day25/input")

	loops := loop(sn, values[0])
	fmt.Println("SN:", sn, "and", values[0], "gives Loop-number", loops)
	result := transform(values[1], loops)
	fmt.Println("SN:", values[1], "for", loops, "loops, gives ", result)

}
