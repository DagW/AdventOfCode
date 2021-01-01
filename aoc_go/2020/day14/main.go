package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
	fmt.Println(readFile("aoc_go/2020/day14/test"))
	fmt.Println(readFile("aoc_go/2020/day14/input"))
}
