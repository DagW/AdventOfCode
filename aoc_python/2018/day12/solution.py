import re
import numpy as np
import pandas as pd
import string
import math
from collections import Counter


def readFile(filename):
    return open(filename).read().strip().split("\n")


def part2(inputArray):
    initState = inputArray[0].split(": ")[1]
    lastState = "." * 10 + initState + "." * 200

    lastVal = sumForRow(lastState)
    lastDiff = 0

    # After about 100 generations the values just move to the right
    # Check nbr of plants stabilizes, then take the diff to predict answer
    for generation in range(0, 100):
        lastState = processGeneration(lastState, inputArray)
        # nbrPlants = np.sum([1 for char in lastState if char == "#"])
        # print(generation, nbrPlants, sumForRow(lastState)-lastVal, sumForRow(lastState))
        newVal = sumForRow(lastState)
        lastDiff = newVal - lastVal
        lastVal = newVal

        # print("Generation", generation, "value:",lastVal)

    predictedValue = lastVal + (lastDiff * (50000000000 - 100))
    return predictedValue


def part1(inputArray):
    initState = inputArray[0].split(": ")[1]
    lastState = "." * 10 + initState + "." * 20
    # print(lastState, 0)

    for generation in range(0, 20):
        lastState = processGeneration(lastState, inputArray)
        # print(lastState, generation)

    sum = sumForRow(lastState)
    return sum


def processGeneration(lastState, inputArray):
    newstate = list("." * len(lastState))

    for potIndex in range(3, len(lastState) - 2):
        for ruleIndex in range(2, len(inputArray)):
            currentRule = inputArray[ruleIndex]
            pots = currentRule.split(" => ")[0]
            result = currentRule.split(" => ")[1]
            pos = (
                lastState[potIndex - 2]
                + lastState[potIndex - 1]
                + lastState[potIndex]
                + lastState[potIndex + 1]
                + lastState[potIndex + 2]
            )

            if pos == pots:
                newstate[potIndex] = result

    lastState = "".join(newstate)
    return lastState


def sumForRow(row):
    sum = 0
    for index in range(0, len(row)):
        if row[index] == "#":
            sum += index - 10
    return sum


if __name__ == "__main__":
    assert (
        part1(
            [
                "initial state: #..#.#..##......###...###",
                "",
                "...## => #",
                "..#.. => #",
                ".#... => #",
                ".#.#. => #",
                ".#.## => #",
                ".##.. => #",
                ".#### => #",
                "#.#.# => #",
                "#.### => #",
                "##.#. => #",
                "##.## => #",
                "###.. => #",
                "###.# => #",
                "####. => #",
            ]
        )
        == 325
    )

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
