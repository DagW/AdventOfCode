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

def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["abcdef","bababc","abbcde","abcccd","aabcdd","abcdee","ababab"]) == 12

    inputData = readFile("input")
    print(part1(inputData))
