package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

type RowStats struct {
	Min int
	Max int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("data")
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	check(err)

	fmt.Println(SolvePart1(csvData))
	fmt.Println(SolvePart2(csvData))
}

func SolvePart1(input [][]string) int {
	sum := 0
	for _, row := range input {
		rowstat := RowStats{Min: math.MaxInt32, Max: math.MinInt32}
		for _, column := range row {
			column, err := strconv.Atoi(column)
			check(err)
			if column < rowstat.Min {
				rowstat.Min = column
			}
			if column > rowstat.Max {
				rowstat.Max = column
			}
		}
		sum += rowstat.Max - rowstat.Min
	}
	return sum
}

func SolvePart2(input [][]string) int {
	sum := 0
	for _, row := range input {
		for _, column := range row {
			column, err := strconv.Atoi(column)
			check(err)
			for _, column2 := range row {
				column2, err := strconv.Atoi(column2)
				check(err)

				if (column%column2) == 0 && column != column2 {
					sum += column / column2
					break
				}
			}
		}
	}
	return sum
}
