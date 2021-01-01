package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	UNKNOWN = 0
	WALL    = 2
	OPEN    = 1
	ITEM    = 3
)

func fromAsciiToInstruction(text string) []int {
	var ints []int
	for _, c := range text {
		ints = append(ints, int(c))
	}
	ints = append(ints, 10)
	return ints
}
func toAscii(numbers []int) string {
	var runes []rune
	for _, number := range numbers {
		runes = append(runes, rune(number))
	}
	return string(runes)
}

func main() {
	file, _ := ioutil.ReadFile("aoc_go/2019/day25.input")
	var program []int
	for _, val := range strings.Split(string(file), ",") {
		i, _ := strconv.Atoi(val)
		program = append(program, i)
	}
	//var shipmap map[struct{ x, y int }]int

	icc := IntCodeComputer{}
	icc.program(program)
	response := toAscii(icc.run(nil))
	fmt.Println(strings.Split(strings.Split(response, "Doors here lead:\n")[1], "Command"))

	//parseResponse(response, &shipmap)

	/*fmt.Println(toAscii(icc.run(fromAsciiToInstruction("north"))))
	icc.clearOutputs()
	fmt.Println(toAscii(icc.run(fromAsciiToInstruction("take wreath"))))
	icc.clearOutputs()
	fmt.Println(toAscii(icc.run(fromAsciiToInstruction("inv"))))*/
}

/*func parseResponse(response string, m *map[struct{ x, y int }]int) {

}*/
