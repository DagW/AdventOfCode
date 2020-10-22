package main

import "fmt"

type IntCodeComputer struct {
	data         map[int]int
	inputs       []int
	outputs      []int
	position     int
	relativeBase int
}

func (icc *IntCodeComputer) program(program []int) {
	icc.data = make(map[int]int)
	for i, item := range program {
		icc.data[i] = item
	}
}
func (icc *IntCodeComputer) popInput() int {
	input := icc.inputs[0]
	icc.inputs = icc.inputs[1:]
	return input
}
func (icc *IntCodeComputer) clearOutputs() {
	icc.outputs = make([]int, 0)
}
func (icc *IntCodeComputer) getData(position int) int {
	return icc.data[position]
}
func (icc *IntCodeComputer) setData(position int, value int) {
	icc.data[position] = value
}
func (icc *IntCodeComputer) run(inputSignals []int) []int {
	icc.inputs = append(icc.inputs, inputSignals...)
	instruction := -1
	run := true
	for run {
		instruction = icc.getData(icc.position) % 100
		mode1 := (icc.getData(icc.position) / 100) % 10
		mode2 := (icc.getData(icc.position) / 1000) % 10
		mode3 := (icc.getData(icc.position) / 10000) % 10
		var pos1, pos2, pos3 int
		switch mode1 {
		case 0:
			pos1 = icc.getData(icc.getData(icc.position + 1))
		case 1:
			pos1 = icc.getData(icc.position + 1)
		case 2:
			pos1 = icc.getData(icc.getData(icc.position+1) + icc.relativeBase)
		}
		switch mode2 {
		case 0:
			pos2 = icc.getData(icc.getData(icc.position + 2))
		case 1:
			pos2 = icc.getData(icc.position + 2)
		case 2:
			pos2 = icc.getData(icc.getData(icc.position+2) + icc.relativeBase)
		}
		switch mode3 {
		case 0:
			pos3 = icc.getData(icc.position + 3)
		case 1:
			pos3 = icc.position + 3
		case 2:
			pos3 = icc.getData(icc.position+3) + icc.relativeBase
		}

		switch instruction {
		case 1:
			icc.setData(pos3, pos1+pos2)
			icc.position += 4
		case 2:
			icc.setData(pos3, pos1*pos2)
			icc.position += 4
		case 3:
			if len(icc.inputs) == 0 {
				fmt.Println("Halting, no more inputs")
				run = false
			} else {
				indata := icc.popInput()
				if mode1 == 2 {
					icc.setData(icc.getData(icc.position+1)+icc.relativeBase, indata)
				} else {
					icc.setData(icc.getData(icc.position+1), indata)
				}
				icc.position += 2
			}
		case 4:
			icc.outputs = append(icc.outputs, pos1)
			icc.position += 2
		case 5:
			if pos1 > 0 {
				icc.position = pos2
			} else {
				icc.position = icc.position + 3
			}
		case 6:
			if pos1 == 0 {
				icc.position = pos2
			} else {
				icc.position = icc.position + 3
			}
		case 7:
			val := 0
			if pos1 < pos2 {
				val = 1
			}
			icc.setData(pos3, val)
			icc.position += 4
		case 8:
			val := 0
			if pos1 == pos2 {
				val = 1
			}
			icc.setData(pos3, val)
			icc.position += 4
		case 9:
			icc.relativeBase += pos1
			icc.position += 2
		case 99:
			fmt.Println("Halting, instruction 99")
			run = false
		default:
			fmt.Println("Empty! wtf")
			run = false
		}
	}
	return icc.outputs
}
