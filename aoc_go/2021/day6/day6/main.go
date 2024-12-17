package day6

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func part1(fish []int, days int) (score int) {
	for day := 0; day < days; day++ {
		for i := 0; i < len(fish); i++ {
			if fish[i] == 0 {
				// New fish time
				fish = append(fish, 9)
				fish[i] = 6
			} else {
				fish[i]--
			}
		}
	}
	return len(fish)
}

func readFile(filename string) (values []int) {
	file, _ := ioutil.ReadFile(filename)
	nums := strings.Split(string(file), ",")
	for _, s := range nums {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	return values
}

func simulate2(data []int, days int) int {
	day0 := 0
	day1 := 0
	day2 := 0
	day3 := 0
	day4 := 0
	day5 := 0
	day6 := 0
	day7 := 0
	day8 := 0

	for j := range data {
		if data[j] == 0 {
			day0 += 1
		}
		if data[j] == 1 {
			day1 += 1
		}
		if data[j] == 2 {
			day2 += 1
		}
		if data[j] == 3 {
			day3 += 1
		}
		if data[j] == 4 {
			day4 += 1
		}
		if data[j] == 5 {
			day5 += 1
		}
		if data[j] == 6 {
			day6 += 1
		}
	}

	for i := 0; i < days; i++ {
		tmp := day0
		day0 = day1
		day1 = day2
		day2 = day3
		day3 = day4
		day4 = day5
		day5 = day6
		day6 = day7
		day7 = day8
		day6 = day6 + tmp
		day8 = tmp
	}
	return day0 + day1 + day2 + day3 + day4 + day5 + day6 + day7 + day8
}

func simulate3(data []int, days int) int {
	tmp := 0
	day0 := 0
	day1 := 0
	day2 := 0
	day3 := 0
	day4 := 0
	day5 := 0
	day6 := 0
	day7 := 0
	day8 := 0

	for j := range data {
		if data[j] == 0 {
			day0 += 1
		}
		if data[j] == 1 {
			day1 += 1
		}
		if data[j] == 2 {
			day2 += 1
		}
		if data[j] == 3 {
			day3 += 1
		}
		if data[j] == 4 {
			day4 += 1
		}
		if data[j] == 5 {
			day5 += 1
		}
		if data[j] == 6 {
			day6 += 1
		}
	}

	for i := 0; i < days; i++ {
		tmp = day0
		day0 = day1
		day1 = day2
		day2 = day3
		day3 = day4
		day4 = day5
		day5 = day6
		day6 = day7
		day7 = day8
		day6 = day6 + tmp
		day8 = tmp
	}
	return day0 + day1 + day2 + day3 + day4 + day5 + day6 + day7 + day8
}

func main() {
	/*fmt.Println("Part1")
	fmt.Println("Test", part1(readFile("aoc_go/2021/day6/test"), 80))
	fmt.Println("Input", part1(readFile("aoc_go/2021/day6/input"), 80))
	fmt.Println("Part2")
	fmt.Println("Test", part2(readFile("aoc_go/2021/day6/test"), 256))*/
	f := readFile("aoc_go/2021/day6/input")

	runs := 1000_000_000
	s := time.Now()
	fmt.Println("abdi", simulate2(f, runs))
	fmt.Println("abdi", time.Now().Sub(s))

	s = time.Now()
	fmt.Println("dag part2_no_modulo", part2_no_modulo(f, runs))
	fmt.Println(time.Now().Sub(s))

	/*s = time.Now()
	fmt.Println("dag original", part2(f, runs))
	fmt.Println(time.Now().Sub(s))

	s = time.Now()
	fmt.Println("dag no loop", part2_noloop(f, runs))
	fmt.Println(time.Now().Sub(s))

	s = time.Now()
	fmt.Println("dag pointers", part2_ptrs(f, runs))
	fmt.Println(time.Now().Sub(s))

	s = time.Now()
	fmt.Println("dag modulo", part2_modulo(f, runs))
	fmt.Println(time.Now().Sub(s))*/

}

/*
Part1
0 5 [0 1 1 2 1 0 0 0 0 0]
1 6 [1 1 2 1 0 0 0 0 0 1]
2 7 [1 2 1 0 0 0 1 0 1 1]
3 9 [2 1 0 0 0 1 1 1 1 2]
4 10 [1 0 0 0 1 1 3 1 2 1]
5 10 [0 0 0 1 1 3 2 2 1 0]
6 10 [0 0 1 1 3 2 2 1 0 0]
7 10 [0 1 1 3 2 2 1 0 0 0]
8 11 [1 1 3 2 2 1 0 0 0 1]
9 12 [1 3 2 2 1 0 1 0 1 1]
10 15 [3 2 2 1 0 1 1 1 1 3]
11 17 [2 2 1 0 1 1 4 1 3 2]
12 19 [2 1 0 1 1 4 3 3 2 2]
13 20 [1 0 1 1 4 3 5 2 2 1]
*/
func part2(fish []int, days int) int {
	fishBuckets := make([]int, 10) // 7 days plus 2 and one "zero day"
	for _, fish := range fish {
		fishBuckets[fish]++
	}
	for day := 0; day < days-1; day++ {
		renewed := fishBuckets[0]

		// Shift the buckets as the fish grow
		for i := 0; i < len(fishBuckets)-1; i++ {
			fishBuckets[i] = fishBuckets[i+1]
		}
		fishBuckets[6] += renewed       // parent fish
		fishBuckets[9] = fishBuckets[0] // newborn fish
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}
func part2_modulo(fish []int, days int) int {
	fishBuckets := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range fish {
		fishBuckets[fish]++
	}

	for day := 0; day < days; day++ {
		fishBuckets[(day+7)%9] += fishBuckets[day%9]
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}
func part2_no_modulo(fish []int, days int) int {
	fishBuckets := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range fish {
		fishBuckets[fish]++
	}

	day := 0
	for day < days {
		for i := 0; i < 9 && day < days; i++ {
			birthday := i + 7
			if birthday > 8 {
				birthday -= 9
			}
			fishBuckets[birthday] += fishBuckets[i]
			day++
		}
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}

func part2_try(fish []int, days int) int {
	fishBuckets := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range fish {
		fishBuckets[fish]++
	}

	for day := 0; day < days; day++ {
		fishBuckets[(day+7)%9] += fishBuckets[day%9]
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}

// Super slow, 25s
func part2_map(fish []int, days int) int {
	fishBuckets := map[int]int{}
	for _, fish := range fish {
		fishBuckets[fish]++
	}

	for day := 0; day < days; day++ {
		fishBuckets[(day+7)%len(fishBuckets)] += fishBuckets[day%(len(fishBuckets))]
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}

func part2_noloop(fish []int, days int) int {
	fishBuckets := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 7 days plus 2 and one "zero day"
	for _, fish := range fish {
		fishBuckets[fish]++
	}
	for day := 0; day < days-1; day++ {
		renewed := fishBuckets[0]

		// Shift the buckets as the fish grow
		//for i := 0; i < len(fishBuckets)-1; i++ {
		//	fishBuckets[i] = fishBuckets[i+1]
		//}
		fishBuckets[0] = fishBuckets[1]
		fishBuckets[1] = fishBuckets[2]
		fishBuckets[2] = fishBuckets[3]
		fishBuckets[3] = fishBuckets[4]
		fishBuckets[4] = fishBuckets[5]
		fishBuckets[5] = fishBuckets[6]
		fishBuckets[6] = fishBuckets[7]
		fishBuckets[7] = fishBuckets[8]
		fishBuckets[8] = fishBuckets[9]

		fishBuckets[6] += renewed       // parent fish
		fishBuckets[9] = fishBuckets[0] // newborn fish
	}

	total := 0
	for _, f := range fishBuckets {
		total += f
	}
	return total
}

func part2_ptrs(fish []int, days int) int {
	fishBuckets := [10]*int{} // 7 days plus 2 and one "zero day"
	for i := range fishBuckets {
		if fishBuckets[i] == nil {
			x := 0
			fishBuckets[i] = &x
		}
	}
	for _, fish := range fish {
		*fishBuckets[fish]++
	}
	for day := 0; day < days-1; day++ {
		renewed := *fishBuckets[0]

		// Shift the buckets as the fish grow
		//for i := 0; i < len(fishBuckets)-1; i++ {
		//	fishBuckets[i] = fishBuckets[i+1]
		//}
		fishBuckets[0] = fishBuckets[1]
		fishBuckets[1] = fishBuckets[2]
		fishBuckets[2] = fishBuckets[3]
		fishBuckets[3] = fishBuckets[4]
		fishBuckets[4] = fishBuckets[5]
		fishBuckets[5] = fishBuckets[6]
		fishBuckets[6] = fishBuckets[7]
		fishBuckets[7] = fishBuckets[8]
		fishBuckets[8] = fishBuckets[9]

		*fishBuckets[6] = *fishBuckets[6] + renewed // parent fish
		fishBuckets[9] = fishBuckets[0]             // newborn fish
	}

	total := 0
	for _, f := range fishBuckets {
		total += *f
	}
	return total
}
