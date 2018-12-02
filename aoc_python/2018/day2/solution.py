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
    # Create dict of how the rows compare to each of the other rows
    # How many chars are matching?
    inverseDistances = {}
    for string in inputArray:
        inverseDistances[string] = {}
        for string2 in [i for i in inputArray if i != string]:
            inverseDistances[string][string2] = sum([1 for left, right in zip(string, string2) if left == right])

    # If they match all but one char, return all matching chars
    stringLenght = len(inputArray[0]) - 1
    match = ""
    for left, matchList in inverseDistances.items():
        for right, value in matchList.items():
            if value == stringLenght:
                for charIndex in range(0, len(left)):
                    if left[charIndex] == right[charIndex]:
                        match += str(left[charIndex])
                return match
    return None


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"]) == 12
    assert part2(["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]) == "fgij"

    inputData = readFile("input")
    assert part1(inputData)
    assert part2(inputData)

