package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

/*
Too slow, even in paralell :-)
func part2(ints []int, powerlevel int) int {
	cache := cache{}
	resultChannel := make(chan []int)
	go backtrack(&cache, []int{0}, ints, powerlevel, resultChannel)
	num := 0
	for c := range resultChannel {
		if num%1000000 == 0 {
			fmt.Println(num, c)
		}
		num++
	}
	return num
}
func backtrack(chain []int, ints []int, target int, resultChannel chan []int) {
	if chain[len(chain)-1] == target-3 {
		chain = append(chain, chain[len(chain)-1]+3)
		resultChannel <- chain
	}
	for i, value := range ints {
		last := chain[len(chain)-1]
		if value-last >= 1 && value-last <= 3 {
			chain = append(chain, value)
			var newInts []int
			newInts = append(newInts, ints[:i]...)
			newInts = append(newInts, ints[i+1:]...)

			if len(chain) == 1 {
				go backtrack(chain, newInts, target, resultChannel)
			} else {
				backtrack(chain, newInts, target, resultChannel)
			}
			chain = chain[:len(chain)-1]
		}
	}
	if len(chain) == 1 {
		// If we tried all starting combinations
		// Signal done
		close(resultChannel)
	}
}
*/

type cache struct {
	store map[int]int
}

func (c *cache) get(value int) (int, bool) {
	if c.store == nil {
		c.store = make(map[int]int)
	}
	item, ok := c.store[value]
	return item, ok
}

func (c *cache) put(value int, result int) {
	if c.store == nil {
		c.store = make(map[int]int)
	}

	c.store[value] = result
}

func countingDfs(cache *cache, ints []int, last int, target int) int {
	if item, ok := cache.get(last); ok {
		return item
	}
	if last+3 == target {
		return 1
	}
	sum := 0
	for i, v := range ints {
		if v-last >= 1 && v-last <= 3 {
			var newInts []int
			newInts = append(newInts, ints[:i]...)
			newInts = append(newInts, ints[i+1:]...)
			ts := countingDfs(cache, newInts, v, target)
			cache.put(v, ts)
			sum += ts
		}
	}
	return sum
}

func part2(ints []int, target int) int {
	n := countingDfs(&cache{}, ints, 0, target)
	return n
}

func part1(ints []int) ([]int, int) {
	chain := []int{0}
	for len(ints) > 0 {
		last := chain[len(chain)-1]
		for i, v := range ints {
			if v == last+1 || v == last+3 {
				chain = append(chain, v)
				ints = append(ints[:i], ints[i+1:]...)
				break
			}
		}
	}
	chain = append(chain, chain[len(chain)-1]+3)
	onej, threej, last := 0, 0, 0
	for _, c := range chain {
		if c-last == 1 {
			onej++
		} else if c-last == 3 {
			threej++
		}
		last = c
	}
	return chain, onej * threej
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
	sort.Ints(values)
	return
}

func main() {

	chain, multiplied := part1(readFile("test"))
	fmt.Println("P1T", multiplied, 35)
	fmt.Println("P2T", part2(readFile("test"), chain[len(chain)-1]))

	chain, multiplied = part1(readFile("test2"))
	fmt.Println("P1T2", multiplied, 220)
	fmt.Println("P2T2", part2(readFile("test2"), chain[len(chain)-1]))

	chain, multiplied = part1(readFile("input"))
	fmt.Println("P1I", multiplied, 35)
	fmt.Println("P2I", part2(readFile("input"), chain[len(chain)-1]))
}
