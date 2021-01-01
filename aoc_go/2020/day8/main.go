package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	RUNNING         = iota
	ALREADYEXECUTED = iota
	DONE            = iota
)

type console struct {
	accumulator int
	position    int
	program     []instruction
}

func (c *console) run() int {
	if c.position >= len(c.program) {
		return DONE
	}
	if c.program[c.position].executed {
		return ALREADYEXECUTED
	}
	c.program[c.position].executed = true
	c.program[c.position].operation(c, c.program[c.position].argument)
	return RUNNING
}

func (c *console) nop(int) {
	c.position++
}

func (c *console) acc(a int) {
	c.accumulator += a
	c.position++
}

func (c *console) jmp(a int) {
	c.position += a
}

type instruction struct {
	argument  int
	operation func(*console, int)
	executed  bool
}

func readFile(filename string) (instructions []instruction) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	r := regexp.MustCompile(`(nop|jmp|acc) ([+\-]\d+)`)
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		operation := (*console).nop
		switch strings.Trim(matches[1], " ") {
		case "acc":
			operation = (*console).acc
		case "jmp":
			operation = (*console).jmp
		}
		argument, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{
			argument:  argument,
			operation: operation,
		})
	}
	return
}

//Ughh... backed myself into a corner
//w making part1 reusable
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	for _, fn := range []string{"aoc_go/2020/day8/test", "aoc_go/2020/day8/input"} {

		c := console{
			accumulator: 0,
			position:    0,
			program:     readFile(fn),
		}
		for c.run() == RUNNING {
		}
		fmt.Println("Part1", c.accumulator)
	}

	for _, fn := range []string{"aoc_go/2020/day8/test", "aoc_go/2020/day8/input"} {

		pos := 0
		instructions := readFile(fn)
	outer:
		for pos < len(instructions) {
			newInstructions := make([]instruction, len(instructions))
			copy(newInstructions, instructions)

			if strings.Contains(GetFunctionName(newInstructions[pos].operation), ".jmp") {
				newInstructions[pos].operation = (*console).nop
			} else if strings.Contains(GetFunctionName(newInstructions[pos].operation), ".nop") {
				newInstructions[pos].operation = (*console).jmp
			}
			c := console{
				accumulator: 0,
				position:    0,
				program:     newInstructions,
			}
			for {
				r := c.run()
				if r == RUNNING {
				} else if r == DONE {
					fmt.Println("Part2", c.accumulator)
					break outer
				} else if r == ALREADYEXECUTED {
					break
				}
			}
			pos++
		}
	}
}
