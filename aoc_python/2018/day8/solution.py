def part1(inputArray):
    items = list(map(int, inputArray[0].split(" ")))

    _, children = parseNodes(items)
    sum = sumChildrenKeys(children)

    return sum


def parseNodes(items, depth=0):
    nbrchildren = items.pop(0)
    nbrmeta = items.pop(0)

    children = []
    for i in range(0, nbrchildren):
        items, child = parseNodes(items, depth + 1)
        children.append(child)

    meta = ",".join(str(x) for x in items[:nbrmeta])

    # return the input array without myself and child nodes
    items = items[nbrmeta:]
    return items, {meta: children}


def sumChildrenKeys(children):
    retSum = 0
    for key, value in children.items():
        retSum += sum([int(c) for c in key.split(",")])
        for item in value:
            retSum += sumChildrenKeys(item)
    return retSum


def part2(inputArray):
    items = list(map(int, inputArray[0].split(" ")))

    _, children = parseNodes(items)
    return sumPart2(children)


def sumPart2(nodes):
    retSum = 0

    for meta, childNodes in nodes.items():
        if len(childNodes) == 0:
            retSum += sum([int(c) for c in meta.split(",")])
        for index in meta.split(","):
            try:
                retSum += sumPart2(childNodes[int(index) - 1])
            except:
                pass
    return retSum


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"]) == 138
    assert part2(["2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"]) == 66

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
