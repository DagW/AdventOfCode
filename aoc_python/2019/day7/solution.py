import itertools
import sys


class IntcodeComputer:
    def __init__(self, data, phase_setting):
        self.data = data
        self.inputs = [phase_setting]
        self.position = 0

    def run(self, inputSignal, start_position=None):
        self.inputs.append(inputSignal)
        self.position = start_position if start_position else self.position

        # print("intcode_computer", inputs)
        # print("start", data, "inputSignal", inputSignal)
        output = None
        while self.data[self.position] != 99:
            instruction = self.data[self.position] % 100
            mode_1 = self.data[self.position] // 100 % 10
            mode_2 = self.data[self.position] // 1000 % 10
            # mode_3 = self.data[self.position] // 10000 % 10 # Not used.. yet?

            # Position mode, else immediate mode
            try:
                pos1 = self.data[self.data[self.position + 1]] if mode_1 == 0 else self.data[self.position + 1]
                pos2 = self.data[self.data[self.position + 2]] if mode_2 == 0 else self.data[self.position + 2]
            except:
                # We might not always get pos2
                pass

            # print(self.position, "parse_modes", instruction, pos1, pos2, self.data)
            if instruction == 1:
                # Add
                self.data[self.data[self.position + 3]] = pos1 + pos2
                # Next instruction
                self.position += 4
            elif instruction == 2:
                # Multiply
                self.data[self.data[self.position + 3]] = pos1 * pos2
                # Next instruction
                self.position += 4
            elif instruction == 3:
                # Input
                if len(self.inputs) > 1:
                    indata = self.inputs.pop(0)
                else:
                    indata = self.inputs[0]
                    # Always pos mode
                self.data[self.data[self.position + 1]] = indata
                # Next instruction
                self.position += 2
            elif instruction == 4:
                # Output
                output = pos1
                # Next instruction
                self.position += 2
            elif instruction == 5:  # Jump if > 0
                # print("instruction ", instruction, "if", pos1, ">0 jump to", pos2)
                if pos1 > 0:
                    self.position = pos2
                else:
                    self.position += 3
            elif instruction == 6:  # Jump if 0
                # print("instruction ", instruction, "if", pos1, "== 0 jump to", pos2)
                if pos1 == 0:
                    self.position = pos2
                else:
                    self.position += 3
            elif instruction == 7:  # less than
                # print("instruction ", instruction, "if", pos1, "<", pos2, " set 1")
                if pos1 < pos2:
                    self.data[self.data[self.position + 3]] = 1
                else:
                    self.data[self.data[self.position + 3]] = 0
                self.position += 4

            elif instruction == 8:  # if ==
                # print("instruction ", instruction, "if", pos1, "==", pos2, " set 1")
                if pos1 == pos2:
                    self.data[self.data[self.position + 3]] = 1
                else:
                    self.data[self.data[self.position + 3]] = 0
                self.position += 4
            else:
                print("Unknown instruction!", instruction)
                sys.exit(1)
        # print("Halted", output)
        return output


def amplifiers(program, phase_sequence, first_input=0):
    amp_a = IntcodeComputer(program, phase_sequence[0])
    amp_b = IntcodeComputer(program, phase_sequence[1])
    amp_c = IntcodeComputer(program, phase_sequence[2])
    amp_d = IntcodeComputer(program, phase_sequence[3])
    amp_e = IntcodeComputer(program, phase_sequence[4])

    output = amp_e.run(amp_d.run(amp_c.run(amp_b.run(amp_a.run(first_input)))))
    return output


def part1(program):
    maxresult = 0
    for perm in list(itertools.permutations([0, 1, 2, 3, 4])):
        maxresult = max(maxresult, amplifiers(program, perm))
    # print("maxresult", maxresult)
    return maxresult


def readFile(filename):
    current_array = open(filename).read().strip().split(",")
    return [int(numeric_string) for numeric_string in current_array]


if __name__ == "__main__":
    print("Part1")
    assert amplifiers(
        [3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0],
        [4, 3, 2, 1, 0]
    ) == 43210
    assert amplifiers(
        [3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0],
        [0, 1, 2, 3, 4]
    ) == 54321
    assert amplifiers(
        [3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31,
         31, 4, 31, 99, 0, 0, 0],
        [1, 0, 4, 3, 2]
    ) == 65210

    assert part1(
        [3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0]
    ) == 43210
    assert part1(
        [3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0]
    ) == 54321
    assert part1(
        [3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31,
         31, 4, 31, 99, 0, 0, 0]
    ) == 65210
    print(part1(readFile("input")))
