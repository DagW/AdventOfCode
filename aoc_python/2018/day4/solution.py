import re
import numpy as np
import pandas as pd


def part1(inputArray):

    # Sort array
    inputArray.sort()

    data = {
        'date':[],
        'id':[]
        }
    for minute in range(0, 60):
        data[minute] = []

    guard = None
    date = ''
    minutes = np.zeros(60)
    for row in inputArray:

        matches = re.compile("\[[0-9]+-([0-9]+-[0-9]+) [0-9]+:([0-9]+)\] (.*)").match(row)
        date = matches.group(1)
        minute = int(matches.group(2))
        message = matches.group(3)

        if 'begins shift' in message:
            if guard != None:
                appendRow(data, date, guard, minutes)
            guard = re.compile(".*#(\d+).*").match(message).group(1)
            minutes = np.zeros(60)
        elif message == "falls asleep":
            minutes[minute:] = 1
        elif message == "wakes up":
            minutes[minute:] = 0

    if guard != None:
        appendRow(data, date, guard, minutes)

    df = pd.DataFrame(data=data)
    # print(df.head())

    resultGuard = -1
    resultTotalMinutes = -1
    resultMinute = -1
    for guard in df.id.unique():
        filt = df.loc[df['id'] == guard]
        totsum = 0
        max = 0
        maxMinute = -1
        for minute in range(0, 60):
            colSum = sum(filt[minute])
            totsum += colSum
            if colSum > max:
                max = colSum
                maxMinute = minute

        # print(guard, totsum, maxMinute)
        if totsum > resultTotalMinutes:
            resultGuard = guard
            resultTotalMinutes = totsum
            resultMinute = maxMinute

    print(resultGuard, resultMinute)
    return int(resultGuard) * int(resultMinute)

def part2(inputArray):

    # Sort array
    inputArray.sort()

    data = {
        'date':[],
        'id':[]
        }
    for minute in range(0, 60):
        data[minute] = []

    guard = None
    date = ''
    minutes = np.zeros(60)
    for row in inputArray:

        matches = re.compile("\[[0-9]+-([0-9]+-[0-9]+) [0-9]+:([0-9]+)\] (.*)").match(row)
        date = matches.group(1)
        minute = int(matches.group(2))
        message = matches.group(3)

        if 'begins shift' in message:
            if guard != None:
                appendRow(data, date, guard, minutes)
            guard = re.compile(".*#(\d+).*").match(message).group(1)
            minutes = np.zeros(60)
        elif message == "falls asleep":
            minutes[minute:] = 1
        elif message == "wakes up":
            minutes[minute:] = 0

    if guard != None:
        appendRow(data, date, guard, minutes)

    df = pd.DataFrame(data=data)
    # print(df.head())

    resultGuard = -1
    resultTotalMinutes = -1
    resultMinute = -1
    for guard in df.id.unique():
        filt = df.loc[df['id'] == guard]
        totsum = 0
        max = 0
        maxMinute = -1
        for minute in range(0, 60):
            colSum = sum(filt[minute])
            totsum += colSum
            if colSum > max:
                max = colSum
                maxMinute = minute
        print(guard, max, maxMinute)

        # print(guard, totsum, maxMinute)
        if max > resultTotalMinutes:
            resultGuard = guard
            resultTotalMinutes = max
            resultMinute = maxMinute

    return int(resultGuard) * int(resultMinute)

def appendRow(data, date, guard, minutes):
    data['date'].append(date)
    data['id'].append(guard)
    for minute in range(0, 60):
        data[minute].append(minutes[minute])

def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1([
        "[1518-11-01 00:00] Guard #10 begins shift",
        "[1518-11-01 00:05] falls asleep",
        "[1518-11-01 00:25] wakes up",
        "[1518-11-01 00:30] falls asleep",
        "[1518-11-01 00:55] wakes up",
        "[1518-11-01 23:58] Guard #99 begins shift",
        "[1518-11-02 00:40] falls asleep",
        "[1518-11-02 00:50] wakes up",
        "[1518-11-03 00:05] Guard #10 begins shift",
        "[1518-11-03 00:24] falls asleep",
        "[1518-11-03 00:29] wakes up",
        "[1518-11-04 00:02] Guard #99 begins shift",
        "[1518-11-04 00:36] falls asleep",
        "[1518-11-04 00:46] wakes up",
        "[1518-11-05 00:03] Guard #99 begins shift",
        "[1518-11-05 00:45] falls asleep",
        "[1518-11-05 00:55] wakes up"
    ]) == 240
    assert part2([
        "[1518-11-01 00:00] Guard #10 begins shift",
        "[1518-11-01 00:05] falls asleep",
        "[1518-11-01 00:25] wakes up",
        "[1518-11-01 00:30] falls asleep",
        "[1518-11-01 00:55] wakes up",
        "[1518-11-01 23:58] Guard #99 begins shift",
        "[1518-11-02 00:40] falls asleep",
        "[1518-11-02 00:50] wakes up",
        "[1518-11-03 00:05] Guard #10 begins shift",
        "[1518-11-03 00:24] falls asleep",
        "[1518-11-03 00:29] wakes up",
        "[1518-11-04 00:02] Guard #99 begins shift",
        "[1518-11-04 00:36] falls asleep",
        "[1518-11-04 00:46] wakes up",
        "[1518-11-05 00:03] Guard #99 begins shift",
        "[1518-11-05 00:45] falls asleep",
        "[1518-11-05 00:55] wakes up"
    ]) == 4455

    inputData = readFile("input")
    print(part1(inputData))
    print(part2(inputData))
