def part1(inputArray):
    twoTimes = 0
    threeTimes = 0
    for string in inputArray:
        letters = {}
        for letterIndex in range(0, len(string)):
            letters[string[letterIndex]] = letters.get(string[letterIndex], 0) + 1
        if 2 in letters.values():
            twoTimes += 1
        if 3 in letters.values():
            threeTimes += 1
    return twoTimes * threeTimes


def part2(inputArray):
    # Create table of how the rows compare to each of the other rows
    # How many chars are matching?
    stringmap = {}
    for string in inputArray:
        stringmap[string] = {}
        for string2 in [i for i in inputArray if i != string]:
            stringmap[string][string2] = 0
            for charIndex in range(0, len(string)):
                if string[charIndex] == string2[charIndex]:
                    stringmap[string][string2] += 1

    # If they match all but one char,
    # return all matching chars
    stringLenght = len(inputArray[0]) - 1
    match = ""
    for left, matchList in stringmap.items():
        for right, value in matchList.items():
            if value == stringLenght:
                for charIndex in range(0, len(left)):
                    if left[charIndex] == right[charIndex]:
                        match += str(left[charIndex])
                return match


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"]) == 12
    assert part2(["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]) == "fgij"

    inputData = readFile("input")
    assert part1(inputData) == 6150
    assert part2(inputData) == "rteotyxzbodglnpkudawhijsc"
