package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part2(values []int) int {

	/*fmt.Println("Part2 Chinese remainders")
	fmt.Println(values)
	return 0*/

	largest := 0
	largestPos := 0
	for i, l := range values {
		if l > largest {
			largest = l
			largestPos = i
		}
	}
	t := largest - largestPos
	for {
		found := true
		for i, val := range values {
			if val == -1 {
			} else {
				timeUntilDeparture := t % val
				timeLeft := val - timeUntilDeparture
				v := timeLeft % val
				if v != i {
					found = false
					break
				}
			}
		}
		if found {
			return t
		}
		t += largest
	}
}

func part1(earliestDeparture int, values []int) int {
	minTime := int(^uint(0) >> 1)
	id := 0
	for _, val := range values {
		timeUntilDeparture := earliestDeparture % val
		timeLeft := val - timeUntilDeparture
		if timeLeft < minTime {
			minTime = timeLeft
			id = val
		}
	}
	return id * minTime
}

func readFileP2(filename string) (values []int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if i != 0 {
			for _, bus := range strings.Split(line, ",") {
				busval, err := strconv.Atoi(bus)
				if err != nil {
					busval = -1
				}
				values = append(values, busval)
			}
		}
	}
	return
}

func readFileP1(filename string) (earliestDeparture int, values []int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if i == 0 {
			earliestDeparture, _ = strconv.Atoi(line)
		} else {
			for _, bus := range strings.Split(line, ",") {
				busval, err := strconv.Atoi(bus)
				if err != nil {
					continue
				}
				values = append(values, busval)
			}
		}
	}
	return
}

func main() {
	fmt.Println("Part1")
	fmt.Println(part1(readFileP1("test")), 295)
	fmt.Println(part1(readFileP1("input")), 6559)
	fmt.Println("Part2")
	fmt.Println(part2(readFileP2("test")), 1068781)
	//fmt.Println(part2(readFileP2("input")), "?")
	fmt.Println("Done")
}
