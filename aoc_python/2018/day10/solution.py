import numpy as np
import re
import math
import sys


class positionClass:
    def __init__(self, positionX, positionY, velocityX, velocityY):
        self.positionX = positionX
        self.positionY = positionY
        self.velocityX = velocityX
        self.velocityY = velocityY

    def run_one_second(self):
        self.positionX += self.velocityX
        self.positionY += self.velocityY


def testMap(dots):
    minx = 100
    miny = 100
    maxx = 0
    maxy = 0
    for dot in dots:
        if dot.positionX > maxx:
            maxx = dot.positionX
        if dot.positionY > maxy:
            maxy = dot.positionY

        if dot.positionX < minx:
            minx = dot.positionX
        if dot.positionY < miny:
            miny = dot.positionY
            
    map = np.zeros((int(math.fabs(maxy - miny)) + 1, int(math.fabs(maxx - minx)) + 1))

    for dot in dots:
        map[int(math.fabs(dot.positionY - miny)), int(math.fabs(dot.positionX - minx))] = 1

    return map


def part1(inputArray):
    dots = []
    for row in inputArray:
        # EX position=< 9,  1> velocity=< 0,  2>
        matches = re.compile(
            "^.*<([ \-\d]+),([ \-\d]+)>.*<([ \-\d]+),([ \-\d]+)>$"
        ).match(row)
        positionX = int(matches.group(1))
        positionY = int(matches.group(2))
        velocityX = int(matches.group(3))
        velocityY = int(matches.group(4))
        dots.append(positionClass(positionX, positionY, velocityX, velocityY))

    lastArea = sys.maxsize
    iteration = 0
    map = None
    while True:
        iteration += 1
        for dot in dots:
            dot.run_one_second()

        newMap = testMap(dots)
        newArea = newMap.shape[0] * newMap.shape[1]

        if newArea > lastArea:
            break
        else:
            map = newMap
            lastArea = newArea
    print(iteration - 1)
    for x in range(len(map)):
        for y in range(len(map[0])):
            if map[x][y] == 0:
                print(".", end="")
            else:
                print("#", end="")
        print()


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    part1(readFile("test"))
    part1(readFile("input"))
