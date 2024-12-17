package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

const gridsize = 300
const (
	PWR       = 0
	MAXWINDOW = 1
	MAXVALUE  = 2
)

func TestPart2(t *testing.T) {
	start := time.Now()
	sn := 7315
	fmt.Println("serialNumber:", sn)
	assert.Equal(t, [3]int{242, 13, 9}, part2(sn))
	fmt.Println("done", time.Since(start))
}

func ()  {
	
}


// 113 seconds
func part2(serialNumber int) interface{} {
	grid := [gridsize][gridsize][3]int{}
	for x := range grid {
		for y := range grid[0] {
			// each grid x,y should save its own max and max window size
			pwr := computePowerLevel(x, y, serialNumber)
			grid[x][y][PWR], grid[x][y][MAXWINDOW], grid[x][y][MAXVALUE] = pwr, 1, pwr
		}
	}
	maxVal, maxWin, maxPos := 0, 0, [2]int{}
	for x := range grid {
		for y := range grid[0] {
			for i := 2; i < 300; i++ {
				sum := 0
				for sX := x + 1; sX < min(i, gridsize); sX++ {
					for sY := y + 1; sY < min(i, gridsize); sY++ {
						sum += grid[sX][sY][PWR]
					}
				}
				if sum > grid[x][y][MAXVALUE] {
					windowSize := i
					grid[x][y][MAXVALUE] = sum
					grid[x][y][MAXWINDOW] = windowSize
					if sum > maxVal {
						maxVal = sum
						maxWin = windowSize
						maxPos = [2]int{x, y}
						fmt.Println("Found max value", maxPos, maxWin, maxVal)
					}
				}
			}
		}
	}
	return [3]int{maxPos[0], maxPos[1], maxWin}
}

func part2_1(serialNumber int) interface{} {
	maxWindowSize := -1
	max := [3]int{}
	grid := [gridsize][gridsize]int{}
	for x := range grid {
		for y := range grid[0] {
			grid[x][y] = computePowerLevel(x, y, serialNumber)
		}
	}

	sumGrid := [gridsize][gridsize]int{}
	for i := 1; i < gridsize; i++ {
		// Sliding window sum mask
		t := [3]int{}
		for x := range grid {
			for y := range grid[0] {

				sum := 0
				for sX := x; sX < min(len(grid)-1, x+i); sX++ {
					for sY := y; sY < min(len(grid[0])-1, y+i); sY++ {
						sum += grid[sX][sY]
					}
				}
				sumGrid[x][y] = sum

				// Keep the max
				if sumGrid[x][y] > t[2] {
					t[0], t[1], t[2] = x, y, sumGrid[x][y]
				}
			}
		}
		if t[2] > max[2] {
			fmt.Println("Found a great window size:", i, t)
			max[0], max[1], max[2] = t[0], t[1], t[2]
			maxWindowSize = i
		}
	}
	max[2] = maxWindowSize
	return max
}

func part1(serialNumber int, windowSize int) [3]int {
	grid := [gridsize][gridsize]int{}
	for x := range grid {
		for y := range grid[0] {
			grid[x][y] = computePowerLevel(x, y, serialNumber)
		}
	}

	// Sliding window sum mask
	max := [3]int{}
	for x := range grid {
		for y := range grid[0] {

			for sX := x; sX < min(len(grid)-1, x+windowSize); sX++ {
				for sY := y; sY < min(len(grid[0])-1, y+windowSize); sY++ {
					grid[x][y] += grid[sX][sY]
				}
			}

			// Keep the max
			if grid[x][y] > max[2] {
				max[0], max[1], max[2] = x, y, grid[x][y]
			}
		}
	}

	return max
}

func TestPart1(t *testing.T) {
	require.Equal(t, [3]int{33, 45, 33}, part1(18, 3))
	require.Equal(t, [3]int{21, 61, 34}, part1(42, 3))
	require.Equal(t, [3]int{21, 72, 34}, part1(7315, 3))
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(i, 3)
	}
}

func computePowerLevel(x, y, serialNumber int) int {
	rackID := x + 10
	powerLevel := y * rackID
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel /= 100
	powerLevel %= 10
	powerLevel -= 5
	return powerLevel
}

func TestComputePowerLevel(t *testing.T) {
	require.Equal(t, computePowerLevel(122, 79, 57), -5)
	require.Equal(t, computePowerLevel(217, 196, 39), 0)
	require.Equal(t, computePowerLevel(101, 153, 71), 4)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
