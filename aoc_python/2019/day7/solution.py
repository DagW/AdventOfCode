import itertools
import sys


class IntcodeComputer:
    def __init__(self, data, phase_setting):
        self.data = data
        self.inputs = [phase_setting]
        self.position = 0

    def run(self, input_signal, start_position=None):
        self.inputs.append(input_signal)
        self.position = start_position if start_position else self.position

        while True:
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

            # print(self.position, "parse_modes", instruction)
            if instruction == 1:  # Add
                # print("intruction add", pos1, pos2, (pos1 + pos2))
                self.data[self.data[self.position + 3]] = pos1 + pos2
                # Next instruction
                self.position += 4
            elif instruction == 2:  # Multiply
                # print("intruction multiply", pos1, pos2, (pos1 * pos2))
                self.data[self.data[self.position + 3]] = pos1 * pos2
                # Next instruction
                self.position += 4
            elif instruction == 3:  # Input
                indata = self.inputs.pop(0)
                # print("instruction input", indata)
                self.data[self.data[self.position + 1]] = indata
                # Next instruction
                self.position += 2
            elif instruction == 4:  # Output
                # print("instruction output", pos1)
                output = pos1
                self.position += 2
                return output
            elif instruction == 5:  # Jump if > 0
                # print("instruction ", instruction, "if", pos1, ">0 jump to", pos2)
                self.position = pos2 if pos1 > 0 else self.position + 3
            elif instruction == 6:  # Jump if 0
                # print("instruction ", instruction, "if", pos1, "== 0 jump to", pos2)
                self.position = pos2 if pos1 == 0 else self.position + 3
            elif instruction == 7:  # less than
                # print("instruction ", instruction, "if", pos1, "<", pos2, " set 1")
                self.data[self.data[self.position + 3]] = 1 if pos1 < pos2 else 0
                self.position += 4
            elif instruction == 8:  # if ==
                # print("instruction ", instruction, "if", pos1, "==", pos2, " set 1")
                self.data[self.data[self.position + 3]] = 1 if pos1 == pos2 else 0
                self.position += 4
            elif instruction == 99:  # Halt
                raise ComputerHalted
            else:
                print("Unknown instruction!", instruction)
                sys.exit(1)


def amplifiers(program, phase_sequence, first_input=0):
    amp_a = IntcodeComputer(program, phase_sequence[0])
    amp_b = IntcodeComputer(program, phase_sequence[1])
    amp_c = IntcodeComputer(program, phase_sequence[2])
    amp_d = IntcodeComputer(program, phase_sequence[3])
    amp_e = IntcodeComputer(program, phase_sequence[4])

    output = amp_e.run(amp_d.run(amp_c.run(amp_b.run(amp_a.run(first_input)))))
    return output


def part1(program):
    max_result = 0
    for perm in list(itertools.permutations([0, 1, 2, 3, 4])):
        max_result = max(max_result, amplifiers(program, perm))
    return max_result


class ComputerHalted(Exception):
    pass


def amplifiers_p2(program, phase_sequence):
    amplifiers = [IntcodeComputer(program.copy(), x) for x in phase_sequence]

    signal = 0
    while True:
        for amp in amplifiers:
            try:
                signal = amp.run(signal)
            except ComputerHalted:
                return signal


def part2(program):
    max_result = 0
    for perm in list(itertools.permutations([5, 6, 7, 8, 9])):
        max_result = max(max_result, amplifiers_p2(program, perm))
    return max_result


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
    print(part1(readFile("input")), "should be", 101490)

    print("Part2")
    assert amplifiers_p2(
        [3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28,
         -1, 28, 1005, 28, 6, 99, 0, 0, 5],
        [9, 8, 7, 6, 5]
    ) == 139629729
    assert amplifiers_p2(
        [3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1,
         53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0,
         0, 10],
        [9, 7, 8, 5, 6]
    ) == 18216
    print(part2(readFile("input")), "should be", 61019896)
