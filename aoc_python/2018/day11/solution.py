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

    return x+1, y+1, sumgrid[x,y]


if __name__ == "__main__":
    # Tests
    assert computePowerlevel(3,5,8) == 4
    assert computePowerlevel(122,79,57) == -5
    assert computePowerlevel(217,196,39) == 0
    assert computePowerlevel(101,153,71) == 4

    assert part1(8) == (21, 186, 29)
    assert part1(18) == (33, 45, 29)
    assert part1(42) == (21, 61, 30)

    print(part1(7315)) # (21, 72, 30.0)
