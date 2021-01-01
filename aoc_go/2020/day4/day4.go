package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func snip(s string, inPassport string) string {
	if strings.Contains(inPassport, s) {
		return strings.Trim(strings.Split(strings.Split(inPassport, s+":")[1], " ")[0], " ")
	}
	return ""
}

type validationItem struct {
	name          string
	mustBePresent bool
}

func (v validationItem) verify(input string) bool {
	if !v.mustBePresent {
		return true
	}
	return len(snip(v.name, input)) > 0
}

func validatePassports(passports []string, validators []validationItem) int {
	numValid := 0
	for _, p := range passports {
		valid := true
		for _, v := range validators {
			if !v.verify(p) {
				valid = false
				break
			}
		}
		if valid {
			numValid++
		}
	}
	return numValid
}

func readFile(filename string) (passports []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	str := ""
	for _, line := range lines {
		if len(strings.Trim(line, " ")) == 0 {
			passports = append(passports, str)
			str = ""
		} else {
			str += " " + line
		}

	}
	return
}

func main() {
	p1 := []validationItem{
		{"cid", false},
		{"iyr", true},
		{"byr", true},
		{"eyr", true},
		{"hgt", true},
		{"hcl", true},
		{"ecl", true},
		{"pid", true},
	}
	fmt.Println("Part1", validatePassports(readFile("aoc_go/2020/day4/day4.test"), p1), 2)
	fmt.Println("Part1", validatePassports(readFile("aoc_go/2020/day4/day4.input"), p1), 256)

}
