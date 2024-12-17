package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) (m [][]int) {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	for _, s := range lines {
		var l []int
		for _, c := range s {
			n, _ := strconv.Atoi(string(c))
			l = append(l, n)
		}
		m = append(m, l)
	}
	return
}

func main() {
	fmt.Println("Part1")
	//fmt.Println("Test mini", part1(readFile("aoc_go/2021/day11/test_mini"), 100))
	fmt.Println("Test", part1(readFile("aoc_go/2021/day11/test"), 100))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day11/input"), 100))
	fmt.Println("Part2")
	//fmt.Println("Test mini", part1(readFile("aoc_go/2021/day11/test_mini"), 100))
	fmt.Println("Test", part2(readFile("aoc_go/2021/day11/test")))
	fmt.Println("Input", part2(readFile("aoc_go/2021/day11/input")))
}

var neighbours = [][2]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}}

func part1(file [][]int, steps int) int {
	flashes := 0

	/*fmt.Println("Before any steps:")
	for _, row := range file {
		fmt.Println(row)
	}*/

	for i := 0; i < steps; i++ {
		// Increase all by 1, collect first flashes
		flashing := map[[2]int]struct{}{}
		has_flashed := map[[2]int]struct{}{}
		for y, row := range file {
			for x := range row {
				file[y][x]++
				if file[y][x] > 9 {
					flashing[[2]int{y, x}] = struct{}{}
				}
			}
		}
		// Flashes, increase neighbours
		for len(flashing) > 0 {
			var o [2]int
			for k := range flashing {
				o = k
				break
			}
			has_flashed[o] = struct{}{}

			for _, n := range neighbours {
				ny, nx := o[0]+n[0], o[1]+n[1]
				if ny >= 0 && ny < len(file) && nx >= 0 && nx < len(file[0]) {
					file[ny][nx]++

					_, hasFlashed := has_flashed[[2]int{ny, nx}]

					if file[ny][nx] > 9 && !hasFlashed {
						flashing[[2]int{ny, nx}] = struct{}{}
					}
				}
			}
			delete(flashing, o)
		}
		// Reset flashed ones
		for o := range has_flashed {
			file[o[0]][o[1]] = 0
		}

		flashes += len(has_flashed)
		/*fmt.Println("After step ", i+1, "(flashes=", flashes, ")")
		for _, row := range file {
			fmt.Println(row)
		}
		fmt.Println()*/
	}
	return flashes
}

func part2(file [][]int) int {


	for i := 0; true; i++ {
		// Increase all by 1, collect first flashes
		flashing := map[[2]int]struct{}{}
		has_flashed := map[[2]int]struct{}{}
		for y, row := range file {
			for x := range row {
				file[y][x]++
				if file[y][x] > 9 {
					flashing[[2]int{y, x}] = struct{}{}
				}
			}
		}
		// Flashes, increase neighbours
		for len(flashing) > 0 {
			var o [2]int
			for k := range flashing {
				o = k
				break
			}
			has_flashed[o] = struct{}{}

			for _, n := range neighbours {
				ny, nx := o[0]+n[0], o[1]+n[1]
				if ny >= 0 && ny < len(file) && nx >= 0 && nx < len(file[0]) {
					file[ny][nx]++

					_, hasFlashed := has_flashed[[2]int{ny, nx}]

					if file[ny][nx] > 9 && !hasFlashed {
						flashing[[2]int{ny, nx}] = struct{}{}
					}
				}
			}
			delete(flashing, o)
		}
		// Reset flashed ones
		for o := range has_flashed {
			file[o[0]][o[1]] = 0
		}

		flashes := len(has_flashed)
		if flashes == len(file)*len(file[0]){
			return i+1
		}
	}
	return -1
}
