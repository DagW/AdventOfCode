import numpy as np

def computePowerlevel(x,y,serialNumber):
    rackId = (x)+10
    powerLevel = (y)*rackId
    powerLevel += serialNumber
    powerLevel *= rackId
    powerLevel /= 100
    powerLevel = int(powerLevel)
    powerLevel %= 10
    powerLevel -= 5
    return powerLevel

def part1(serialNumber):

    grid = np.zeros((300,300))

    for x in range(len(grid)):
        for y in range(len(grid)):
            grid[x][y] = computePowerlevel(x+1,y+1, serialNumber)

    # Sliding window sum mask
    sumgrid = np.zeros((300,300))
    for x in range(len(sumgrid)):
        for y in range(len(sumgrid)):
            sumgrid[x][y] = np.sum(grid[x:x+3, y:y+3])

    x,y = np.unravel_index(sumgrid.argmax(), sumgrid.shape)

    return x+1, y+1 # sumgrid[x,y]

def part2(serialNumber):

    grid = np.zeros((300,300))

    for x in range(len(grid)):
        for y in range(len(grid)):
            grid[x][y] = computePowerlevel(x+1,y+1, serialNumber)

    # Sliding window sum mask
    sizeVals = {}
    # Running all 300 takes 381.1s
    # but after 15 it really just goes down
    for i in range(1,15): 
        print("Checking", i, end=" ")
        sumgrid = np.zeros((300,300))
        for x in range(len(sumgrid)):
            for y in range(len(sumgrid)):
                sumgrid[x][y] = np.sum(grid[x:x+i, y:y+i])

        x,y = np.unravel_index(sumgrid.argmax(), sumgrid.shape)
        sizeVals[int(sumgrid[x,y])] = str(x)+","+str(y)+","+str(sumgrid[x,y])
        print(x,y,sumgrid[x,y])

    print(x,y,sizeVals)

    return max(sizeVals, key=sizeVals.get)


if __name__ == "__main__":
    assert computePowerlevel(3,5,8) == 4
    assert computePowerlevel(122,79,57) == -5
    assert computePowerlevel(217,196,39) == 0
    assert computePowerlevel(101,153,71) == 4

    assert part1(8) == (21, 186)
    assert part1(18) == (33, 45)
    assert part1(42) == (21, 61)

    print(part1(7315)) # (21, 72, 30.0)
    print(part2(7315)) # 242,13,9
