import sys


class Halted(Exception):
    def __init__(self, last_output):
        self.last_output = last_output


class IntcodeComputer:
    def __init__(self, data):
        self.data = {}
        for index, item in enumerate(data):
            self.data[index] = item
        self.inputs = []
        self.position = 0
        self.relative_base = 0
        self.outputs = []

    def clear_outputs(self):
        self.outputs = []

    def get_data(self, position):
        if position < 0:
            raise ValueError("Cant address memory below 0 ({})".format(position))
        return self.data.get(position, 0)

    def set_data(self, position, item):
        self.data[position] = item

    def run(self, input_signals, start_position=None):
        # print(self.inputs, input_signals)
        self.inputs = self.inputs + input_signals
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
                if len(self.inputs) == 0:
                    # print("HALTING, no inputs!")
                    break
                indata = self.inputs.pop(0)
                # self.inputs[0] if len(self.inputs) == 1 else self.inputs.pop(0)
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
            elif instruction == 99:
                raise Halted(self.outputs)
            else:
                print("Unknown instruction!", instruction)
                sys.exit(1)
        return self.outputs

# def part1(program):
#    c = IntcodeComputer(program)
#    return c.run(1)
