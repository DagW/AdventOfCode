def part1(data):
    position = 0
    while True:
        if position > len(data):
            break
        instruction = data[position]
        if instruction == 1:
            # Add
            pos1 = data[position + 1]
            pos2 = data[position + 2]
            sum = data[pos1] + data[pos2]
            data[data[position + 3]] = sum
        elif instruction == 2:
            # Multiply
            pos1 = data[position + 1]
            pos2 = data[position + 2]
            sum = data[pos1] * data[pos2]
            data[data[position + 3]] = sum
        elif instruction == 99:
            # Stop
            break
        # Next instruction
        position += 4
    # print(data)
    return data


def readFile(filename):
    items = open(filename).read().strip().split(",")
    return list(map(int, items))


if __name__ == "__main__":
    assert part1([1, 0, 0, 0, 99]) == [2, 0, 0, 0, 99]
    assert part1([2, 3, 0, 3, 99]) == [2, 3, 0, 6, 99]
    assert part1([2, 4, 4, 5, 99, 0]) == [2, 4, 4, 5, 99, 9801]
    assert part1([1, 1, 1, 4, 99, 5, 6, 0, 99]) == [30, 1, 1, 4, 2, 5, 6, 0, 99]

    # Gravity assist initiation
    inputData = readFile("input")
    inputData[1] = 12
    inputData[2] = 2
    print(part1(inputData)[0])
