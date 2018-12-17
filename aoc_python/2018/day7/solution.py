import re
import numpy as np
import pandas as pd
import string
import math
from collections import Counter


def readFile(filename):
    return open(filename).read().strip().split("\n")


def part1(inputArray):
    graph = {}
    for row in inputArray:
        matches = re.compile(".*([A-Z]).*([A-Z]).*").match(row)
        dependency = matches.group(1)
        toPoint = matches.group(2)
        if dependency not in graph:
            graph[dependency] = []
        if toPoint in graph:
            graph[toPoint].append(dependency)
        else:
            graph[toPoint] = [dependency]

    visited = []
    tovisit = []

    for key, value in graph.items():
        if graph[key] == []:
            tovisit.append(key)
        # Sort each list a-z
        graph[key] = "".join(sorted(value))

    # visit all nodes in Dependency-tree
    # Start with items without deps
    while len(tovisit) > 0:
        currentPoint = tovisit.pop(0)

        # Remove all deps we have visited already
        for char in list(visited):
            graph[currentPoint] = graph[currentPoint].replace(char, "")

        # If we cant visit it yet, place it at the end of the list
        if graph[currentPoint] != "":
            tovisit.append(currentPoint)

        # Visit!
        visited.append(currentPoint)

        # What opens up now?
        for key in sorted(graph.keys()):

            dependencies = graph[key]
            # If all deps in visited
            if set(dependencies).issubset(visited):
                # print(key, "can be visited now",dependencies)
                tovisit.append(key)

        # Make sure we visit a-z
        tovisit = sorted(tovisit)
        tovisit = [x for x in tovisit if x not in visited]
        print("Current:", currentPoint, "tovisit:", tovisit)

    print("Visited", visited)

    return "".join(visited)


if __name__ == "__main__":
    assert (
        part1(
            [
                "Step C must be finished before step A can begin.",
                "Step C must be finished before step F can begin.",
                "Step A must be finished before step B can begin.",
                "Step A must be finished before step D can begin.",
                "Step B must be finished before step E can begin.",
                "Step D must be finished before step E can begin.",
                "Step F must be finished before step E can begin.",
            ]
        )
        == "CABDFE"
    )

    inputData = readFile("input")
    print(part1(inputData))
