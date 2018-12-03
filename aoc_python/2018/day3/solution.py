import re
import numpy as np


def part1(inputArray):
    cloth = np.zeros((1000, 1000))
    for claim in inputArray:
        matches = re.compile("^#([\d]+) @ ([\d]+),([\d]+): ([\d]+)x([\d]+)$").match(claim)
        c_id = int(matches.group(1))
        c_x = int(matches.group(2))
        c_y = int(matches.group(3))
        c_w = int(matches.group(4))
        c_h = int(matches.group(5))

        for x in range(0, len(cloth)):
            for y in range(0, len(cloth[0])):
                cloth_position = cloth[x][y]
                if x >= c_x and x < c_x + c_w and y >= c_y and y < c_y + c_h:
                    if cloth_position == 0:
                        cloth[x][y] = c_id
                    else:
                        cloth[x][y] = -1
    sum = 0
    for x in range(0, len(cloth)):
        for y in range(0, len(cloth[0])):
            if cloth[x][y] == -1:
                sum += 1
    print("cloth collisions:", sum)
    return sum


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2", "#2 @ 253,106: 10x25"]) == 4

    inputData = readFile("input")
    assert part1(inputData)
