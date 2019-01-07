import numpy as np
import copy
from pprint import pprint
import datetime

# Above, left, right, below
searcharray = [(-1, 0), (0, -1), (0, +1), (+1, 0)]


def getUnits(grid):
    units = []

    # Determine unit order, l2r, u2d
    for row in range(0, len(grid)):
        for column in range(0, len(grid[row])):
            point = grid[row][column]
            if point[0] in "EG":
                location = (row, column)
                units.append(location)

    return units


def getHitpoints(grid):
    sumHp = 0
    for row in range(0, len(grid)):
        for column in range(0, len(grid[row])):
            if "," in grid[row][column]:
                sumHp += int(grid[row][column].split(",")[1])
    return sumHp


def printGrid(grid, separator=" "):
    for row in range(0, len(grid)):
        for column in range(0, len(grid[row])):
            try:
                print(grid[row][column][0], end=separator)
            except:
                print(grid[row][column], end=separator)
        print("")


def attack(grid, source, targets):
    # Default is first in reading order
    target = targets[0]
    # print("Select target from ", targets, "with the least amount of HP")
    minHp = 200
    for ltarget in targets:
        hp = int(grid[ltarget[0]][ltarget[1]].split(",")[1])
        if hp < minHp:
            minHp = hp
            target = ltarget

    sourceUnit = grid[source[0]][source[1]].split(",")
    targetUnit = grid[target[0]][target[1]].split(",")
    sourceUnitType = sourceUnit[0]
    sourceHP = int(sourceUnit[1])
    sourceAP = int(sourceUnit[2])
    targetUnitType = targetUnit[0]
    targetHP = int(targetUnit[1])
    targetAP = int(targetUnit[2])

    targetHP -= sourceAP

    targetString = targetUnitType + "," + str(targetHP) + "," + str(targetAP)
    if targetHP <= 0:
        targetString = "."
        # print(targetUnit, "died")
    grid[target[0]][target[1]] = targetString


def calcDistance(grid, source, target):
    # print("calcDistance from", source,target)
    # and return distance to the target, and the next step on the closest path
    distance = -1
    nextStep = None

    visited = {}
    visited[source] = 0

    tovisit = [source]

    foundTarget = False

    while len(tovisit) > 0:
        currentPos = tovisit.pop()
        for modifier in searcharray:
            modPos = (currentPos[0] + modifier[0], currentPos[1] + modifier[1])
            modPosChar = grid[modPos[0]][modPos[1]]

            if modPos == target:
                foundTarget = True

            if modPosChar == ".":
                prevDist = -1
                if modPos in visited:
                    prevDist = visited[modPos]
                    if visited[currentPos] + 1 < prevDist:
                        visited[modPos] = visited[currentPos] + 1
                        tovisit.append(modPos)
                else:
                    visited[modPos] = visited[currentPos] + 1
                    tovisit.append(modPos)

    if not foundTarget:
        return -1, None

    # Get distance to target
    if target in visited:
        distance = visited[target]

    # Find the path
    gridcopy = copy.deepcopy(grid)
    for visitedPoint, pdistance in visited.items():
        gridcopy[visitedPoint[0]][visitedPoint[1]] = pdistance

    path = [target]
    count = 0
    if distance > 1:
        while len(path) < distance:
            for modifier in searcharray:
                modPos = (target[0] + modifier[0], target[1] + modifier[1])
                modPosChar = gridcopy[modPos[0]][modPos[1]]
                if str(gridcopy[modPos[0]][modPos[1]]) == str(distance - len(path)):
                    path.append(modPos)
                    target = (modPos[0], modPos[1])
                    break
    nextStep = path[-1]

    return distance, nextStep


def calcDistances(grid, source, targets):
    distances = []
    nextSteps = []
    for target in targets:
        distance, nextStep = calcDistance(grid, source, target)
        distances.append(distance)
        nextSteps.append(nextStep)

    return distances, nextSteps


def moveUnit(grid, fromPos, toPos):
    currentUnit = grid[fromPos[0]][fromPos[1]]
    grid[toPos[0]][toPos[1]] = grid[fromPos[0]][fromPos[1]]
    grid[fromPos[0]][fromPos[1]] = "."


def part1(inputArray, baseHP="200", baseAP="3", output=True):

    grid = []
    for row in range(0, len(inputArray)):
        gridrow = []
        for column in range(0, len(inputArray[row])):
            point = inputArray[row][column]
            if point in "EG":
                # ex. E, 200hp, 3 attack power
                gridrow.append(point + "," + baseHP + "," + baseAP)
            else:
                gridrow.append(point)
        grid.append(gridrow)

    rounds = 0

    if output:
        printGrid(grid)

    while True:
        if output:
            print(
                "###### ROUND",
                rounds,
                "###### begins ",
                datetime.datetime.now().strftime("%d.%b %Y %H:%M:%S"),
            )
        # Each unit takes a turn
        for currentPos in getUnits(grid):
            x = currentPos[0]
            y = currentPos[1]
            currentUnit = grid[x][y]
            if currentUnit != ".":

                possibleTargets = getUnits(grid)
                # Remove current pos, and similar units
                possibleTargets.remove(currentPos)
                possibleTargets = [
                    x
                    for x in possibleTargets
                    if grid[x[0]][x[1]][0] in "EG".replace(currentUnit[0], "")
                ]

                if len(possibleTargets) == 0:
                    score = rounds * getHitpoints(grid)
                    if output:
                        print("Battle ended!", score)
                        printGrid(grid)
                    return score
                else:
                    # If anyone next to me is an enemy, attack
                    targets = []
                    for modifier in searcharray:
                        point = grid[currentPos[0] + modifier[0]][
                            currentPos[1] + modifier[1]
                        ]
                        if point[0] in "EG".replace(currentUnit[0], ""):
                            targets.append(
                                (
                                    currentPos[0] + modifier[0],
                                    currentPos[1] + modifier[1],
                                )
                            )
                    if len(targets) > 0:
                        attack(grid, source=currentPos, targets=targets)
                        # print("Attack",targets)

                    # Otherwise move
                    else:
                        # identify . near possible targets
                        for possibleTarget in possibleTargets:
                            for modifier in searcharray:
                                point = grid[possibleTarget[0] + modifier[0]][
                                    possibleTarget[1] + modifier[1]
                                ]
                                if point[0] in ".":
                                    targets.append(
                                        (
                                            possibleTarget[0] + modifier[0],
                                            possibleTarget[1] + modifier[1],
                                        )
                                    )

                        if len(targets) > 0:
                            # Get distances to each target
                            distances, nextSteps = calcDistances(
                                grid, currentPos, targets
                            )

                            # Move towards min distance (If a tie, normal order)
                            minDistance = 5000
                            minIndex = -1
                            for index in range(0, len(targets)):
                                if (
                                    distances[index] != -1
                                    and distances[index] < minDistance
                                ):
                                    minDistance = distances[index]
                                    minIndex = index

                            if minIndex != -1:
                                moveUnit(grid, (x, y), nextSteps[minIndex])
                                currentPos = nextSteps[minIndex]

                            # If anyone next to me is an enemy, attack again
                            targets = []
                            for modifier in searcharray:
                                point = grid[currentPos[0] + modifier[0]][
                                    currentPos[1] + modifier[1]
                                ]
                                if point[0] in "EG".replace(currentUnit[0], ""):
                                    targets.append(
                                        (
                                            currentPos[0] + modifier[0],
                                            currentPos[1] + modifier[1],
                                        )
                                    )

                            if len(targets) > 0:
                                attack(grid, source=currentPos, targets=targets)

                        else:
                            # print("Could not move, no possible targets")
                            pass

        if output:
            printGrid(grid)
        rounds += 1

    return -1


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":

    assert (
        part1(
            [
                "#######",
                "#G..#E#",
                "#E#E.E#",
                "#G.##.#",
                "#...#E#",
                "#...E.#",
                "#######",
            ],
            output=False,
        )
        == 36334
    )
    assert (
        part1(
            [
                "#######",
                "#E..EG#",
                "#.#G.E#",
                "#E.##E#",
                "#G..#.#",
                "#..E#.#",
                "#######",
            ],
            output=False,
        )
        == 39514
    )

    inputData = readFile("input")
    print(part1(inputData))  # 139.7s
