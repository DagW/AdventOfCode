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
        cloth[c_x : c_x + c_w, c_y : c_y + c_h] += 1
    return len(np.where(cloth > 1)[0])


def part2(inputArray):
    cloth = np.zeros((1000, 1000))
    for claim in inputArray:
        matches = re.compile("^#([\d]+) @ ([\d]+),([\d]+): ([\d]+)x([\d]+)$").match(claim)
        c_id = int(matches.group(1))
        c_x = int(matches.group(2))
        c_y = int(matches.group(3))
        c_w = int(matches.group(4))
        c_h = int(matches.group(5))
        cloth[c_x : c_x + c_w, c_y : c_y + c_h] += 1

    for claim in inputArray:
        matches = re.compile("^#([\d]+) @ ([\d]+),([\d]+): ([\d]+)x([\d]+)$").match(claim)
        c_id = int(matches.group(1))
        c_x = int(matches.group(2))
        c_y = int(matches.group(3))
        c_w = int(matches.group(4))
        c_h = int(matches.group(5))
        if sum(sum(cloth[c_x : c_x + c_w, c_y : c_y + c_h])) == c_w * c_h:
            return c_id


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"]) == 4
    assert part2(["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"]) == 3

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
