import math
import sys


def test_asteroid(map, from_y, from_x, part2=False):
    asteroids = {}
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if column == '#' and (x, y) != (from_x, from_y):
                # Get relative position of other asteroid
                rel_x = x - from_x
                rel_y = y - from_y
                angle = math.degrees(math.atan2(rel_y, rel_x))
                if part2:
                    angle += 90
                    # set 0 degrees north
                    if angle < 0:
                        # Start at 0, rotate clockwise
                        angle = 360 + angle
                if angle in asteroids:
                    if part2:
                        # if we have a collission! Best manhattan dist wins
                        other = asteroids[angle]
                        if (abs(from_y - y) + abs(from_x - x)) < (abs(from_y - other[0]) + abs(from_x - other[1])):
                            asteroids[angle] = (y, x)
                    else:
                        pass
                else:
                    asteroids[angle] = (y, x)
    if part2:
        retarr = []
        # print(asteroids)
        for key in sorted(asteroids.keys()):
            # print(asteroids[key], key)
            retarr.append(asteroids[key])
        return retarr
    else:
        return asteroids.values()


def part1(map):
    # print(map)
    # print()
    max_visibles = 0
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if column == '#':
                visibles = len(test_asteroid(map, y, x))
                max_visibles = max(max_visibles, visibles)
                # print(y, x, "visibles", visibles)
                # print(visibles, end='')
            # else:
            # print(".", end='')
        # print()
    return max_visibles


def part2(map):
    # print(map)
    # print("----")
    max_visibles = 0
    position = (-1, -1)
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if column == '#':
                visibles = len(test_asteroid(map, y, x))
                if visibles > max_visibles:
                    max_visibles = visibles
                    position = (y, x)

    """print("Best position", position[1], position[0])
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if x == position[1] and y == position[0]:
                column = 'X'
            print(column, end='')
        print()
    print("----")"""

    # Now destroy some asteroids
    destroyed_asteroids = test_asteroid(map, position[0], position[1], True)
    return destroyed_asteroids[199][1] * 100 + destroyed_asteroids[199][0]


def readFile(filename):
    return open(filename).read()


if __name__ == "__main__":
    assert part1(""".#..#
.....
#####
....#
...##""") == 8
    assert part1("""......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####""") == 33
    assert part1("""#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.""") == 35
    assert part1(""".#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..""") == 41
    assert part1(""".#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##""") == 210

    print(part1(readFile("input")))

    assert part2(""".#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##""") == 802
    print(part2(readFile("input")))
