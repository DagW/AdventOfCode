package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func contains(ints []int, target int) bool {
	for _, i := range ints {
		if i == target {
			return true
		}
	}
	return false
}

/*func toDag(ints []int, target int) int {
	ints = append(ints, 0)
	ints = append(ints, target)
	dag := map[int][]int{}
	for _, v := range ints {
		if contains(ints, v+1){
			dag[v] = append(dag[v], v+1)
		}
		if contains(ints, v+2){
			dag[v] = append(dag[v], v+2)
		}
		if contains(ints, v+3){
			dag[v] = append(dag[v], v+3)
		}
	}
	fmt.Println(dag)

	return 0
}*/

func backtrack(chain []int, ints []int) ([]int, error) {
	if len(ints) == 0 {
		chain = append(chain, chain[len(chain)-1]+3)
		return chain, nil
	}
	for i, value := range ints {
		if value-chain[len(chain)-1] == 3 || value-chain[len(chain)-1] == 1 {
			chain = append(chain, value)
			var newInts []int
			newInts = append(newInts, ints[:i]...)
			newInts = append(newInts, ints[i+1:]...)
			temp, err := backtrack(chain, newInts)
			if err == nil {
				return temp, nil
			}
			chain = chain[:len(chain)-1]
		}
	}
	return chain, errors.New("couldnt fint the chain")
}

func part1(ints []int) int {
	chain, _ := backtrack([]int{0}, ints)
	fmt.Println(chain)
	onej, threej, last := 0, 0, 0
	for _, c := range chain {
		if c-last == 1 {
			onej++
		} else if c-last == 3 {
			threej++
		}
		last = c
	}
	fmt.Println(onej, threej)
	return onej * threej
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
	fmt.Println(part1(readFile("test")))
	fmt.Println(part1(readFile("input")))
	/*fmt.Println(part2(readFile("aoc_go/2020/day10/test"), 22))
	fmt.Println(part2(readFile("aoc_go/2020/day10/input"), 170))*/
}
