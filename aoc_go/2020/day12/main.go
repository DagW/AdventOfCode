package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(values []string) int {
	y, x, sy, sx := 1, 10, 0, 0
	for _, v := range values {
		m, _ := strconv.Atoi(v[1:])
		switch v[0] {
		case 'N':
			y += m
		case 'S':
			y -= m
		case 'E':
			x += m
		case 'W':
			x -= m
		case 'L':
			switch m {
			case 90:
				y, x = x, -y
			case 180:
				y, x = -y, -x
			case 270:
				y, x = -x, y
			}
		case 'R':
			switch m {
			case 90:
				y, x = -x, y
			case 180:
				y, x = -y, -x
			case 270:
				y, x = x, -y
			}
		case 'F':
			sy += y * m
			sx += x * m
		}
	}

	return abs(sy) + abs(sx)
}
func part1(values []string) int {
	y, x, dir := 0, 0, 90
	for _, v := range values {
		m, _ := strconv.Atoi(v[1:])

		switch v[0] {
		case 'N':
			y += m
		case 'S':
			y -= m
		case 'E':
			x += m
		case 'W':
			x -= m
		case 'L':
			dir -= m
			if dir < 0 {
				dir += 360
			}
			dir %= 360
		case 'R':
			dir += m
			dir %= 360
		case 'F':
			switch dir {
			case 0:
				y += m
			case 90:
				x += m
			case 180:
				y -= m
			case 270:
				x -= m
			}
		}
	}

	return abs(y) + abs(x)
}

func readFile(filename string) (values []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		values = append(values, line)
	}
	return
}

func main() {
	fmt.Println(part1(readFile("test")))
	fmt.Println(part1(readFile("input")))
	fmt.Println(part2(readFile("test")))
	fmt.Println(part2(readFile("input")))
}
