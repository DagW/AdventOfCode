import re
import numpy as np
import pandas as pd
import string
import math
from collections import Counter


def readFile(filename):
    return open(filename).read().strip().split("\n")


def manhattanDistance(pointA, pointB):
    return math.fabs(pointA[0] - pointB[0]) + math.fabs(pointA[1] - pointB[1])


def part2(inputArray, regionsize=10000):
    coordinates = np.zeros((len(inputArray), 2))
    for index in range(0, len(inputArray)):
        coordinates[index][0] = inputArray[index].split(", ")[0]
        coordinates[index][1] = inputArray[index].split(", ")[1]

    c_w = coordinates.max(axis=0)[0] + 1
    c_h = coordinates.max(axis=0)[1] + 1
    coordMap = np.zeros((int(c_w), int(c_h)))
    coordMap[:] -= 1  # -1 for no-ones point

    # Fill the map with manhattan distance sums for each coordinate
    for x in range(0, len(coordMap)):
        for y in range(0, len(coordMap[0])):
            coordMap[x][y] = sumDistances([x, y], coordinates)
    # Check how many places are within the regionsize
    return len(np.where(coordMap < regionsize)[0])


def sumDistances(point, coordinates):
    sum = 0
    for index in range(0, len(coordinates)):
        sum += manhattanDistance(point, coordinates[index])
    return sum


def part1(inputArray):
    coordinates = np.zeros((len(inputArray), 2))
    for index in range(0, len(inputArray)):
        coordinates[index][0] = inputArray[index].split(", ")[0]
        coordinates[index][1] = inputArray[index].split(", ")[1]

    c_w = coordinates.max(axis=0)[0] + 1
    c_h = coordinates.max(axis=0)[1] + 1
    coordMap = np.zeros((int(c_w), int(c_h)))
    coordMap[:] -= 1  # -1 for no-ones point

    # Fill the map with coordinate row indexes
    for x in range(0, len(coordMap)):
        for y in range(0, len(coordMap[0])):
            coordMap[x][y] = getOwnerForPoint([x, y], coordinates)

    # now remove all indexes touching the walls
    for x in range(0, len(coordMap)):
        for y in range(0, len(coordMap[0])):
            if x == 0 or y == 0 or x == len(coordMap) - 1 or y == len(coordMap[0]) - 1:
                if coordMap[x][y] != -1:
                    value = coordMap[x][y]
                    coordMap[coordMap == value] = -1

    unique, counts = np.unique(coordMap, return_counts=True)
    emptyIndex = np.where(unique == -1)
    unique = np.delete(unique, emptyIndex, 0)
    counts = np.delete(counts, emptyIndex, 0)
    return counts[np.argmax(counts)]


def getOwnerForPoint(point, coordinates):
    distanceMap = {}
    for index in range(0, len(coordinates)):
        distanceMap[index] = manhattanDistance(point, coordinates[index])

    least_distances = Counter(distanceMap).most_common()[:-3:-1]
    if least_distances[0][1] == least_distances[1][1]:
        return -1
    else:
        return least_distances[0][0]


if __name__ == "__main__":
    assert part1(["1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"]) == 17
    assert part2(["1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9"], 32) == 16
    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
