package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func apply(mem *map[int]int64, mask string, address, value int) {
	result := "000000000000000000000000000000000000"
	strval := strconv.FormatInt(int64(value), 2)
	/*fmt.Print("Value: ")
	for i := 0; i < 36; i++ {
		if i >= 36-len(strval) {
			fmt.Print(string(strval[len(strval)-(36-i)]))
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()*/

	for i, c := range mask {
		if c == '0' || c == '1' {
			result = result[:i] + string(mask[i]) + result[i+1:]
		} else {
			if i >= 36-len(strval) {
				result = result[:i] + string(strval[len(strval)-(36-i)]) + result[i+1:]
			}
		}
	}
	//fmt.Println("Mask: ", mask)
	//fmt.Println("Resu: ", result)
	output, _ := strconv.ParseInt(result, 2, 64)

	(*mem)[address] = output
}

func part1(input []string) int {
	mem := map[int]int64{}
	var currentMask string
	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			currentMask = strings.Split(s, "= ")[1]
		} else if strings.HasPrefix(s, "mem") {
			i, _ := strconv.Atoi(strings.Split(strings.Split(s, "[")[1], "]")[0])
			v, _ := strconv.Atoi(strings.Split(s, "= ")[1])
			apply(&mem, currentMask, i, v)
		}
	}
	sum := 0
	for k := range mem {
		sum += int(mem[k])
	}
	return sum
}

func readFile(filename string) (values []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

func main() {
	fmt.Println(part1(readFile("test")))
	fmt.Println(part1(readFile("input")))
}
