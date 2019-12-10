import math


def test_asteroid(map, from_y, from_x):
    asteroids = {}
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if column == '#' and (x, y) != (from_x, from_y):
                # Get relative position of other asteroid
                rel_x = x - from_x
                rel_y = y - from_y
                angle = math.degrees(math.atan2(rel_y, rel_x))
                if angle in asteroids:
                    pass
                    """
                    Below not needed to solve the problem..
                    # if we have a collission! Best manhattan dist :-)
                    other = asteroids[angle]
                    if (abs(from_y - y) + abs(from_x - x)) < (abs(from_y - other[0]) + abs(from_x - other[1])):
                        asteroids[angle] = (y, x)
                    """
                else:
                    asteroids[angle] = (y, x)
    return len(asteroids.keys())


def part1(map):
    # print(map)
    # print()
    max_visibles = 0
    for y, row in enumerate(map.split("\n")):
        for x, column in enumerate(row):
            if column == '#':
                visibles = test_asteroid(map, y, x)
                max_visibles = max(max_visibles, visibles)
                # print(y, x, "visibles", visibles)
                # print(visibles, end='')
            # else:
            # print(".", end='')
        # print()
    return max_visibles


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
