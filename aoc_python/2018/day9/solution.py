import numpy as np


class dlMarbleList:
    def __init__(self, value):
        self.current = {"value": None, "next": None, "prev": None}
        self.current["value"] = value
        self.current["next"] = self.current
        self.current["prev"] = self.current

    def getCurrent(self):
        return self.current

    def addMarble(self, invalue):
        """newCurrent = {
            "value": invalue,
            "next": self.current["next"]["next"],
            "prev": self.current["next"]
        }
        # Insert new obj..
        self.current["next"]["next"] = newCurrent
        self.current["next"]["next"]["prev"] = newCurrent
        self.current = newCurrent"""

        prev = self.current["next"]
        next = prev["next"]
        newCurrent = {"value": invalue, "next": next, "prev": prev}
        prev["next"] = newCurrent
        next["prev"] = newCurrent
        self.current = newCurrent

    def unlink(self):
        removed = self.current
        # print("Removing value ",self.current["value"])
        self.current["prev"]["next"] = self.current["next"]
        self.current["next"]["prev"] = self.current["prev"]
        self.current = self.current["next"]
        # print("New value ",self.current["value"])
        return removed

    def shiftClockwise(self):
        # print("Shifting from",self.current["value"], end=' ')
        self.current = self.current["next"]
        # print("to",self.current["value"])
        return self.current

    def shiftCounterClockwise(self):
        # print("Shifting from",self.current["value"], end=' ')
        self.current = self.current["prev"]
        # print("to",self.current["value"])
        return self.current

    def printMarbles(self, length):
        values = []
        for i in range(length + 1):
            values.append(self.shiftClockwise()["value"])
        return values


def part1(inputArray):
    # https://oeis.org/A006257
    players = inputArray[0]
    lastMarbleWorth = inputArray[1]

    scores = np.zeros(players)

    # print(scores)
    dlist = dlMarbleList(0)
    # print(dlist.current)
    # print(0, [0])
    for marble in range(1, lastMarbleWorth + 1):
        player = (marble) % len(scores)
        if marble % 23 == 0:
            scores[player] += marble
            for i in range(7):
                dlist.shiftCounterClockwise()
                # print(player, dlist.printMarbles(marble), '-', dlist.getCurrent()["value"])
            removed = dlist.unlink()
            scores[player] += removed["value"]
            # print(player, dlist.printMarbles(marble), '-', dlist.getCurrent()["value"])
        else:
            # print("Adding marble", marble)
            dlist.addMarble(marble)
            # print(player, dlist.printMarbles(marble), '-', dlist.getCurrent()["value"])

    # dlist.addMarble(5)
    # print(dlist.current)
    # dlist.printMarbles(marble)
    return max(scores)


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    # part1([10,30])
    # part1([10,1618])
    assert part1([10, 1618]) == 8317
    assert part1([13, 7999]) == 146373
    assert part1([17, 1104]) == 2764
    assert part1([21, 6111]) == 54718
    assert part1([30, 5807]) == 37305

    inputData = readFile("input")
    print(part1([491, 71058]))
    print(part1([491, 71058 * 100]))
