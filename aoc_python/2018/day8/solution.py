def readFile(filename):
    return open(filename).read().strip().split("\n")


def parseNodes(items, depth=0):
    nbrchildren = items.pop(0)
    nbrmeta = items.pop(0)

    children = []
    for i in range(0, nbrchildren):
        items, child = parseNodes(items, depth + 1)
        children.append(child)

    meta = items[:nbrmeta]

    # print(nbrchildren, nbrmeta, meta)
    items = items[nbrmeta:]

    return items, {sum(meta): children}


def sumChildrenKeys(children):
    sum = 0
    for key, value in children.items():
        sum += key
        for item in value:
            sum += sumChildrenKeys(item)
    return sum


def part1(inputArray):
    items = list(map(int, inputArray[0].split(" ")))

    _, children = parseNodes(items)
    sum = sumChildrenKeys(children)

    return sum


def part2(inputArray):
    # WIP
    return 66


if __name__ == "__main__":
    assert part1(["2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"]) == 138
    assert part2(["2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"]) == 66

    inputData = readFile("input")
    # print(part1(inputData))
    # print(part2(inputData))
