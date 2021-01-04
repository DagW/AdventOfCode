package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part2(mem *map[int]int, mask string, address, value int) {
	//fmt.Println("Adding", value, "to", address)
	result := "000000000000000000000000000000000000"

	strval := strconv.FormatInt(int64(address), 2)
	/*fmt.Print("Address: ")
	for i := 0; i < 36; i++ {
		if i >= 36-len(strval) {
			fmt.Print(string(strval[len(strval)-(36-i)]))
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()*/

	for i, c := range mask {
		if c == '0' {
			if i >= 36-len(strval) {
				result = result[:i] + string(strval[len(strval)-(36-i)]) + result[i+1:]
			}
		} else if c == '1' {
			result = result[:i] + string(mask[i]) + result[i+1:]
		} else if c == 'X' {
			result = result[:i] + string(mask[i]) + result[i+1:]
		}
	}
	//fmt.Println("Mask:   ", mask)
	//fmt.Println("Result: ", result)

	//Generate the permutations
	str := ""
	for i := 0; i < strings.Count(result, "X"); i++ {
		str += "1"
	}
	max, _ := strconv.ParseInt(str, 2, 64)
	items := []string{str}
	for i := int64(0); i < max; i++ {
		item := strconv.FormatInt(i, 2)
		for len(item) < len(str) {
			item = "0" + item
		}
		items = append(items, item)
	}
	//fmt.Println(items)
	results := []string{}
	for _, s := range items {
		copy := result
		for _, c := range s {
			new := string(c)
			copy = strings.Replace(copy, "X", new, 1)
			//i = i[1:]
		}
		results = append(results, copy)
	}
	for _, s := range results {
		a, _ := strconv.ParseInt(s, 2, 64)
		(*mem)[int(a)] = value
	}
}

func part1(mem *map[int]int, mask string, address, value int) {
	result := "000000000000000000000000000000000000"
	strval := strconv.FormatInt(int64(value), 2)
	for i, c := range mask {
		if c == '0' || c == '1' {
			result = result[:i] + string(mask[i]) + result[i+1:]
		} else {
			if i >= 36-len(strval) {
				result = result[:i] + string(strval[len(strval)-(36-i)]) + result[i+1:]
			}
		}
	}
	output, _ := strconv.ParseInt(result, 2, 64)
	(*mem)[address] = int(output)
}

func apply(input []string, f func(mem *map[int]int, mask string, address, value int)) int {
	mem := map[int]int{}
	var currentMask string
	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			currentMask = strings.Split(s, "= ")[1]
		} else if strings.HasPrefix(s, "mem") {
			i, _ := strconv.Atoi(strings.Split(strings.Split(s, "[")[1], "]")[0])
			v, _ := strconv.Atoi(strings.Split(s, "= ")[1])
			f(&mem, currentMask, i, v)
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
	fmt.Println(apply(readFile("test"), part1), 165)
	fmt.Println(apply(readFile("input"), part1), 8332632930672)
	fmt.Println(apply(readFile("test2"), part2), 208)
	fmt.Println(apply(readFile("input"), part2), 4753238784664)
}
