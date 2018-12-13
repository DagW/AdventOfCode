import re
import numpy as np
import pandas as pd
import string
import math
from collections import Counter


def readFile(filename):
    return open(filename).read().strip().split("\n")


def part1(inputArray):
    initState = inputArray[0].split(": ")[1]
    lastState = "."*10 + initState + "."*20
    print(lastState, 0)

    for generation in range(1, 21):
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
        print(lastState, generation)

    sum = 0
    for index in range(0, len(lastState)):
        if lastState[index] == "#":
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
