import sys


class ComputerHalted(Exception):
    pass


class IntcodeComputer:
    def __init__(self, data):
        self.data = {}
        for index, item in enumerate(data):
            self.data[index] = item
        self.inputs = []
        self.position = 0
        self.relative_base = 0
        self.outputs = []

    def get_data(self, position):
        if position < 0:
            raise ValueError("Cant address memory below 0 ({})".format(position))
        return self.data.get(position, 0)

    def set_data(self, position, item):
        self.data[position] = item

    def run(self, input_signal, start_position=None):
        self.inputs.append(input_signal)
        self.position = start_position if start_position else self.position

        while True:
            instruction = self.get_data(self.position) % 100
            mode_1 = self.get_data(self.position) // 100 % 10
            mode_2 = self.get_data(self.position) // 1000 % 10
            mode_3 = self.get_data(self.position) // 10000 % 10

            if mode_1 == 0:  # Pos mode
                pos1 = self.get_data(self.get_data(self.position + 1))
            elif mode_1 == 1:  # Immediate mode
                pos1 = self.get_data(self.position + 1)
            elif mode_1 == 2:  # Rel mode
                pos1 = self.get_data(self.get_data(self.position + 1) + self.relative_base)

            try:
                if mode_2 == 0:
                    pos2 = self.get_data(self.get_data(self.position + 2))
                elif mode_2 == 1:
                    pos2 = self.get_data(self.position + 2)
                elif mode_2 == 2:
                    pos2 = self.get_data(self.get_data(self.position + 2) + self.relative_base)
            except Exception:
                pass

            try:
                if mode_3 == 0:
                    pos3 = self.get_data(self.position + 3)
                elif mode_3 == 1:
                    pos3 = self.position + 3
                elif mode_3 == 2:
                    pos3 = self.get_data(self.position + 3) + self.relative_base
            except Exception:
                pass

            if instruction == 1:
                self.set_data(pos3, pos1 + pos2)
                self.position += 4
            elif instruction == 2:
                self.set_data(pos3, pos1 * pos2)
                self.position += 4
            elif instruction == 3:
                indata = self.inputs.pop(0)
                if mode_1 == 2:
                    self.set_data(self.get_data(self.position + 1) + self.relative_base, indata)
                else:
                    self.set_data(self.get_data(self.position + 1), indata)
                self.position += 2
            elif instruction == 4:
                output = pos1
                self.position += 2
                self.outputs.append(output)
            elif instruction == 5:
                self.position = pos2 if pos1 > 0 else self.position + 3
            elif instruction == 6:
                self.position = pos2 if pos1 == 0 else self.position + 3
            elif instruction == 7:
                self.set_data(pos3, 1 if pos1 < pos2 else 0)
                self.position += 4
            elif instruction == 8:
                self.set_data(pos3, 1 if pos1 == pos2 else 0)
                self.position += 4
            elif instruction == 9:
                self.relative_base += pos1
                self.position += 2
            elif instruction == 99:  # Halt
                # raise ComputerHalted
                break
            else:
                print("Unknown instruction!", instruction)
                sys.exit(1)
        return self.outputs


def part1(program):
    c = IntcodeComputer(program)
    return c.run(1)


def part2(program):
    c = IntcodeComputer(program)
    return c.run(2)


def readFile(filename):
    current_array = open(filename).read().strip().split(",")
    return [int(numeric_string) for numeric_string in current_array]


if __name__ == "__main__":
    print("Part1")
    assert part1([109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99]) \
           == [109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99]
    assert part1([1102, 34915192, 34915192, 7, 4, 7, 99, 0]) == [1219070632396864]
    assert part1([104, 1125899906842624, 99]) == [1125899906842624]
    print(part1(readFile("input")))
    print("Part2")
    print(part2(readFile("input")))
