package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	size     int
	children []string
}

func getDirs(filename string) map[int][]string {

	var sizes = map[string]node{}

	file, err := os.ReadFile(filename)
	noError(err)
	path := []string{}
	lines := strings.Split(string(file), "\n")
	lines = append(lines, "$ exit")

	lno := 0
	currentCmd := ""
	result := []string{}
	for lno < len(lines) {
		l := lines[lno]
		if strings.HasPrefix(l, "$") {

			if len(result) > 0 {
				fmt.Println("Process", strings.Join(path, "/"), currentCmd, result)
				// Process result
				n := node{}
				for _, rl := range result {
					sp := strings.Split(rl, " ")
					if sp[0] == "dir" {
						n.children = append(n.children, strings.Join(append(path, sp[1]), "/"))
					} else {
						s, _ := strconv.Atoi(sp[0])
						n.size += s
					}
				}
				sizes[strings.Join(path, "/")] = n
				result = []string{}
			}

			currentCmd = strings.Split(l, " ")[1]
			if strings.HasPrefix(l, "$ cd ..") {
				path = path[0 : len(path)-1]
			} else if strings.HasPrefix(l, "$ cd") {
				path = append(path, strings.Split(l, "$ cd ")[1])
			}
		} else {
			result = append(result, l)
		}
		lno++
	}

	dirs := map[string]int{}
	done := false
	for !done {
		done = true
		for k, v := range sizes {
			if len(v.children) == 0 {
				dirs[k] = v.size
			} else if len(v.children) > 0 {
				done = false

				child := v.children[0]
				if len(sizes[child].children) == 0 {
					v.children = v.children[1:]
					s := sizes[child].size
					v.size += s
					sizes[k] = v
					dirs[child] = s
					break
				}
			}
		}
	}

	r := map[int][]string{}
	for k, v := range dirs {
		r[v] = append(r[v], k)
	}
	return r
}

func part1(filename string) int {
	dirs := getDirs(filename)

	resultSizes := 0

	keys := []int{}
	for k := range dirs {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		size, paths := k, dirs[k]
		sort.Strings(paths)
		for _, path := range paths {
			fmt.Println(" - ", path, size)
			if size < 100000 {
				resultSizes += size
			}
		}
	}

	return resultSizes
}

func part2(filename string) int {
	dirs := getDirs(filename)

	need := 70000000 - 30000000
	used := getTotal(dirs)
	diff := used - need

	keys := []int{}
	for k := range dirs {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		if v > diff {
			return v
		}
	}

	return diff
}

func getTotal(dirs map[int][]string) int {
	for size, paths := range dirs {
		for _, path := range paths {
			if path == "/" {
				return size
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(part1("test"))
	fmt.Println(part1("input"))
	fmt.Println(part2("test"))
	fmt.Println(part2("input"))
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
