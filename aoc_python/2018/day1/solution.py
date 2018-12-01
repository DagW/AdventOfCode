def part1(inputArray):
    sum = 0
    for item in inputArray:
        sum += int(item)
    return sum


def part2(inputArray):
    currentFrequency = 0
    seenFrequencies = {0}
    while True:
        for item in inputArray:
            currentFrequency += int(item)
            if currentFrequency in seenFrequencies:
                return currentFrequency
            seenFrequencies.add(currentFrequency)


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1([+1, -2, +3, +1]) == 3
    assert part1([+1, +1, +1]) == 3
    assert part1([+1, +1, -2]) == 0
    assert part1([-1, -2, -3]) == -6

    assert part2([+1, -1]) == 0
    assert part2([+3, +3, +4, -2, -4]) == 10
    assert part2([-6, +3, +8, +5, -6]) == 5
    assert part2([+7, +7, -2, -7, -4]) == 14

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
