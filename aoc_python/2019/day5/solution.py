def get_input():
    return 1


def get_input_part2():
    print("returning input 5")
    return 5


def set_output(data):
    print("output", data)


def part1(data):
    print("start", data)
    position = 0
    while True:
        if position > len(data):
            break

        instruction = data[position] % 100
        mode_1 = int(data[position] / 100) % 10
        mode_2 = int(data[position] / 1000) % 10
        mode_3 = int(data[position] / 10000) % 10

        print(instruction, mode_1, mode_2, mode_3)
        if instruction == 1:
            # Add
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]

            sum = pos1 + pos2
            data[data[position + 3]] = sum
            # Next instruction
            position += 4
        elif instruction == 2:
            # Multiply
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]

            sum = pos1 * pos2
            data[data[position + 3]] = sum
            # Next instruction
            position += 4
        elif instruction == 3:
            # Input
            indata = get_input()
            data[data[position + 1]] = indata
            # Next instruction
            position += 2
        elif instruction == 4:
            # Output
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]

            set_output(pos1)
            # Next instruction
            position += 2

        elif instruction == 99:
            # Stop
            break
    print("end", data)
    return data


def part2(data, input):
    print("start", data)
    position = 0
    while True:
        if position > len(data):
            break

        instruction = data[position] % 100
        mode_1 = int(data[position] / 100) % 10
        mode_2 = int(data[position] / 1000) % 10
        mode_3 = int(data[position] / 10000) % 10

        print("parse_modes", instruction, mode_1, mode_2, mode_3, data)
        if instruction == 1:
            # Add
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]

            sum = pos1 + pos2
            data[data[position + 3]] = sum
            # Next instruction
            position += 4
        elif instruction == 2:
            # Multiply
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]

            sum = pos1 * pos2
            data[data[position + 3]] = sum
            # Next instruction
            position += 4
        elif instruction == 3:
            # Input
            indata = input
            data[data[position + 1]] = indata
            # Next instruction
            position += 2
        elif instruction == 4:
            # Output
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]

            set_output(pos1)
            # Next instruction
            position += 2
        elif instruction == 5:  # Jump if > 0
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]
            # print("instruction ", instruction, "if", pos1, ">0 jump to", pos2)
            if pos1 > 0:
                position = pos2
            else:
                position += 3
        elif instruction == 6:  # Jump if 0
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]
            print("instruction ", instruction, "if", pos1, "== 0 jump to", pos2)
            if pos1 == 0:
                position = pos2
            else:
                position += 3
        elif instruction == 7:  # less than
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]
            print("instruction ", instruction, "if", pos1, "<", pos2, " set 1")
            if pos1 < pos2:
                data[data[position + 3]] = 1
            else:
                data[data[position + 3]] = 0
            position += 4

        elif instruction == 8:  # if ==
            if mode_1 == 1:  # Immediate mode
                pos1 = data[position + 1]
            else:  # Position mode
                pos1 = data[data[position + 1]]
            if mode_2 == 1:  # Immediate mode
                pos2 = data[position + 2]
            else:  # Position mode
                pos2 = data[data[position + 2]]
            print("instruction ", instruction, "if", pos1, "==", pos2, " set 1")
            if pos1 == pos2:
                data[data[position + 3]] = 1
            else:
                data[data[position + 3]] = 0
            position += 4
        elif instruction == 99:
            # Stop
            break
        else:
            print("Unknown instruction!", instruction)
            break
    print("end", data)
    return data


def readFile(filename):
    items = open(filename).read().strip().split(",")
    return list(map(int, items))


if __name__ == "__main__":
    # assert part1([1101, 100, -1, 4, 0]) == [1101, 100, -1, 4, 99]
    print(part1(readFile("input")))
    print(part2(readFile("input"), 5))
    # part2([3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8], 10)  # 1
    # part2([3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8])  # 0
    # part2([3, 3, 1108, -1, 8, 3, 4, 3, 99])  # 0
    # part2([3, 3, 1107, -1, 8, 3, 4, 3, 99])  # 1
    # part2([3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9], 0)
    # part2([3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9], 1)
    # part2([3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1], 0)
    # part2([3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1], 1)

    """part2([3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
           1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
           999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99], 1)
    part2([3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
           1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
           999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99], 8)
    part2([3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
           1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
           999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99], 10)"""
