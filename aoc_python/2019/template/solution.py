def part1(data):
    return 1


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["0,0,0,0", "3,0,0,0"]) == 2

    print(part1(readFile("input")))
