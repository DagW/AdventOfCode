package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	up    = iota
	down  = iota
	left  = iota
	right = iota
)

type image struct {
	id     string
	pixels [][]int
}

func (i *image) flip() {
	newPixels := make([][]int, len(i.pixels))
	for j, row := range i.pixels {
		newPixels[j] = make([]int, len(row))
	}
	for y, row := range i.pixels {
		for x, val := range row {
			newPixels[y][len(row)-x-1] = val
		}
	}
	i.pixels = newPixels
}

func (i *image) rotate() {
	newPixels := make([][]int, len(i.pixels))
	for j, row := range i.pixels {
		newPixels[j] = make([]int, len(row))
	}
	for y, row := range i.pixels {
		for x := range row {
			newPixels[y][x] = i.pixels[len(row)-x-1][y]
		}
	}
	i.pixels = newPixels
}

func (i *image) trim() {
	newPixels := make([][]int, len(i.pixels)-2)
	for j := range newPixels {
		newPixels[j] = make([]int, len(i.pixels)-2)
	}
	for y, row := range newPixels {
		for x := range row {
			newPixels[y][x] = i.pixels[y+1][x+1]
		}
	}
	i.pixels = newPixels
}

func (i *image) getEdge(edge int) []int {
	if edge == up {
		return i.pixels[0]
	} else if edge == right {
		var pixels []int
		for _, row := range i.pixels {
			pixels = append(pixels, row[len(row)-1])
		}
		return pixels
	} else if edge == down {
		return i.pixels[len(i.pixels)-1]
	} else {
		var pixels []int
		for _, row := range i.pixels {
			pixels = append(pixels, row[0])
		}
		return pixels
	}
}

func (i *image) fitsIn(frame [][]image, y int, x int) bool {
	type neighbour struct {
		direction int
		matchEdge []int
	}
	// Do we have neighbours
	var neighbours []neighbour
	if y != 0 && frame[y-1][x].id != "" {
		neighbours = append(neighbours, neighbour{up, frame[y-1][x].getEdge(down)})
	}
	if y != len(frame)-1 && frame[y+1][x].id != "" {
		neighbours = append(neighbours, neighbour{down, frame[y+1][x].getEdge(up)})
	}
	if x != 0 && frame[y][x-1].id != "" {
		neighbours = append(neighbours, neighbour{left, frame[y][x-1].getEdge(right)})
	}
	if x != len(frame[0])-1 && frame[y][x+1].id != "" {
		neighbours = append(neighbours, neighbour{right, frame[y][x+1].getEdge(left)})
	}

	// For each neighbour - do we fit
	fits := true
	for _, n := range neighbours {
		if !compare(i.getEdge(n.direction), n.matchEdge) {
			fits = false
			break
		}
	}

	return fits
}

func (i *image) mask(mask [][]int, y int, x int) {
	total := 0
	for _, row := range mask {
		for _, col := range row {
			if col == 1 {
				total++
			}
		}
	}
	matches := 0
	for my, row := range mask {
		for mx, col := range row {
			if col == 1 && i.pixels[y+my][x+mx] == 1 {
				matches++
			}
		}
	}
	if matches == total {
		for my, row := range mask {
			for mx, col := range row {
				if col == 1 {
					i.pixels[y+my][x+mx] = 2
				}
			}
		}
	}
}

func compare(edge []int, edge2 []int) bool {
	if len(edge) != len(edge2) {
		return false
	}
	for i, v := range edge {
		if edge2[i] != v {
			return false
		}
	}
	return true
}

func printImage(i image) {
	fmt.Println("--- Image", i.id, " --- (", len(i.pixels), "x", len(i.pixels[0]), ")")
	for _, row := range i.pixels {
		for _, col := range row {
			if col == 1 {
				fmt.Print("#")
			} else if col == 2 {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--- End ---")
	fmt.Println()
}

func printFrame(frame [][]image) {

	fmt.Println("--- Frame ---")
	for _, row := range frame {
		for _, c := range row {
			if c.id == "" {
				fmt.Print("EMPTY ")
			} else {
				fmt.Print(c.id, "  ")
			}
		}
		fmt.Println()
	}
	fmt.Println("--- End ---")
	fmt.Println()
}

func pop(m map[string]image) (image, map[string]image) {
	key := ""
	for k := range m {
		key = k
		break
	}
	img, ok := m[key]
	if ok {
		delete(m, key)
	}
	return img, m
}

func getNextPositions(frame [][]image) [][2]int {
	var positions [][2]int
	for y, row := range frame {
		for x := range row {
			if frame[y][x].id != "" {
				//Skip this one
			} else if y-1 >= 0 && frame[y-1][x].id != "" {
				positions = append(positions, [2]int{y, x})
			} else if y+1 < len(frame) && frame[y+1][x].id != "" {
				positions = append(positions, [2]int{y, x})
			} else if x-1 >= 0 && frame[y][x-1].id != "" {
				positions = append(positions, [2]int{y, x})
			} else if x+1 < len(frame) && frame[y][x+1].id != "" {
				positions = append(positions, [2]int{y, x})
			}
		}
	}
	return positions
}

func part1(images map[string]image) (int, [][]image) {
	// There must be a match to all frames.
	// So start with the first..
	side := int(math.Sqrt(float64(len(images)))) * 2
	frame := make([][]image, side)
	for i := 0; i < side; i++ {
		frame[i] = make([]image, side)
	}
	x, y := (side/2)-1, (side/2)-1

	current, images := pop(images)
	frame[y][x] = current

	// Locations that share a side
	// Aka all locations that have a neighbour, but not occupied
	for len(images) > 0 {

		positions := getNextPositions(frame)
		for _, pos := range positions {

		imageLoop:
			for k := range images {
				i := images[k]
				for f := 0; f < 2; f++ {
					for r := 0; r < 4; r++ {
						if i.fitsIn(frame, pos[0], pos[1]) {
							frame[pos[0]][pos[1]] = i
							delete(images, i.id)
							break imageLoop
						}
						i.rotate()
					}
					i.flip()
				}
			}
		}
	}

	//printFrame(frame)
	result := 1
	for y := 0; y < len(frame); y++ {
		for x := 0; x < len(frame); x++ {
			if frame[y][x].id != "" {
				if (y-1 < 0 || frame[y-1][x].id == "") && (x-1 < 0 || frame[y][x-1].id == "") {
					// Top left
					val, _ := strconv.Atoi(frame[y][x].id)
					result *= val
				} else if (y+1 > len(frame) || frame[y+1][x].id == "") && (x+1 > len(frame) || frame[y][x+1].id == "") {
					// Bottom right
					val, _ := strconv.Atoi(frame[y][x].id)
					result *= val
				} else if (y-1 < 0 || frame[y-1][x].id == "") && (x+1 > len(frame) || frame[y][x+1].id == "") {
					// Top right
					val, _ := strconv.Atoi(frame[y][x].id)
					result *= val
				} else if (y+1 > len(frame) || frame[y+1][x].id == "") && (x-1 < 0 || frame[y][x-1].id == "") {
					// Bottom left
					val, _ := strconv.Atoi(frame[y][x].id)
					result *= val
				}
			}
		}
	}
	return result, frame
}

func part2(frame [][]image) int {
	// Trim the edges
	for y, row := range frame {
		for x, image := range row {
			if image.id != "" {
				frame[y][x].trim()
			}
		}
	}

	// Expand it
	maxX, maxY, minX, minY := 0, 0, 10_000, 10_1000
	for y, row := range frame {
		for x, col := range row {
			if col.id != "" {
				switch {
				case x < minX:
					minX = x
					fallthrough
				case x > maxX:
					maxX = x
					fallthrough
				case y < minY:
					minY = y
					fallthrough
				case y > maxY:
					maxY = y
				}
			}
		}
	}
	subImages := maxY - minY + 1
	subImageSize := len(frame[minY][minX].pixels)
	pixels := make([][]int, subImages*subImageSize)
	for i := range pixels {
		pixels[i] = make([]int, subImages*subImageSize)
	}
	for y, row := range pixels {
		for x := range row {
			pixels[y][x] = frame[minY+(y/subImageSize)][minX+(x/subImageSize)].pixels[y%subImageSize][x%subImageSize]
		}
	}
	fullImage := image{pixels: pixels}
	//printImage(fullImage)

	dragon := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
	}
	for f := 0; f < 2; f++ {
		for r := 0; r < 4; r++ {
			//Erase the dragons..
			for y := 0; y < len(fullImage.pixels)-len(dragon); y++ {
				for x := 0; x < len(fullImage.pixels[0])-len(dragon[0]); x++ {
					//Apply the mask here..Match?
					fullImage.mask(dragon, y, x)
				}
			}
			fullImage.rotate()
		}
		fullImage.flip()
	}

	hashes := 0
	for _, row := range fullImage.pixels {
		for _, col := range row {
			if col == 1 {
				hashes++
			}
		}
	}
	return hashes
}

func readFile(filename string) (images map[string]image) {
	images = make(map[string]image)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	r := regexp.MustCompile(`Tile (\d{4}):`)
	img := image{}
	for _, line := range lines {
		if len(strings.Trim(line, " ")) == 0 {

			if img.id != "" {
				images[img.id] = img
			}

		}
		matches := r.FindStringSubmatch(line)
		if len(matches) > 0 {
			img = image{}
			img.id = matches[1]
		} else {
			var bs []int
			for _, c := range line {
				p := 0
				if c == '#' {
					p = 1
				}
				bs = append(bs, p)
			}
			img.pixels = append(img.pixels, bs)
		}
	}
	return
}

func main() {
	result, testFrame := part1(readFile("aoc_go/2020/day20/test"))
	fmt.Println("Part1 test->", result)
	result, inputFrame := part1(readFile("aoc_go/2020/day20/input"))
	fmt.Println("Part1 input->", result)

	fmt.Println(part2(testFrame))
	fmt.Println(part2(inputFrame))
}
