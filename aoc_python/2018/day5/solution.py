import re
import numpy as np
import pandas as pd
import string

def part2(inputArray):
    inputString = inputArray[0]
    minLength = len(inputString)+1
    for char in list(string.ascii_lowercase):
        tempString = inputString.replace(char, "")
        tempString = tempString.replace(char.upper(), "")
        newLength = part1([tempString])

        if newLength < minLength:
            minLength = newLength

    return minLength

def part1(inputArray):
    inputString = inputArray[0]

    permutations = []
    for char in list(string.ascii_lowercase):
        permutations.append(char+char.upper())
        permutations.append(char.upper()+char)

    changes = 1
    while changes > 0:
        changes -= 1
        for permutation in permutations:
            before = len(inputString)
            inputString = inputString.replace(permutation, "")
            if len(inputString) != before:
                changes += 1

    return len(inputString)


def part1slow(inputArray):
    # How to improve
    # Try generating list of aA/Ab/... and do string replace
    inputString = inputArray[0]
    loop = True
    while loop:
        changes = 0
        prevChar = inputString[0]
        for index in range(1, len(inputString)):
            prevChar = inputString[index-1]
            curChar = inputString[index]
            if prevChar != curChar:
                if (
                    prevChar.lower() == curChar or
                    prevChar.upper() == curChar 
                    ):
                    inputString = inputString[:index-1] + inputString[index+1:]
                    changes += 1
                    break

        if changes == 0:
            loop = False

    return len(inputString)

def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1([
        "dabAcCaCBAcCcaDA"
    ]) == 10
    assert part2([
        "dabAcCaCBAcCcaDA"
    ]) == 4

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
