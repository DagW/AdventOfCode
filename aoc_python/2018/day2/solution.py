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
    # Look for string pairs matching all but one char
    matchLenght = len(inputArray[0]) - 1
    for i in range(0, len(inputArray)):
        for j in range(0, len(inputArray)):
            # Dont compare strings to themselves
            if i != j:
                left = inputArray[i]
                right = inputArray[j]
                inverseDistance = sum([1 for leftChar, rightChar in zip(left, right) if leftChar == rightChar])
                if inverseDistance == matchLenght:
                    returnChars = [leftChar for leftChar, rightChar in zip(left, right) if leftChar == rightChar]
                    return ''.join(returnChars)
    # But probably, i should just costruct a matrix with above?


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"]) == 12
    assert part2(["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]) == "fgij"

    inputData = readFile("input")
    assert part1(inputData)
    assert part2(inputData)

