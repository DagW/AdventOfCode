def print_map(wire_positions):
    # print(wire_positions)
    y_min = x_min = y_max = x_max = 0
    for pos in wire_positions[0]:
        x = pos[0]
        y = pos[1]
        if x > x_max:
            x_max = x
        if x < x_min:
            x_min = x
        if y > y_max:
            y_max = y
        if y < y_min:
            y_min = y

    for pos in wire_positions[1]:
        x = pos[0]
        y = pos[1]
        if x > x_max:
            x_max = x
        if x < x_min:
            x_min = x
        if y > y_max:
            y_max = y
        if y < y_min:
            y_min = y

    map = []
    for y in range(y_min, y_max + 3):
        column = []
        for x in range(x_min, x_max + 3):
            column.append(".")
        map.append(column)
    for pos in wire_positions[0]:
        map[pos[1] + 1 + abs(y_min)][pos[0] + 1 + abs(x_min)] = "x"
    for pos in wire_positions[1]:
        if map[pos[1] + 1 + abs(y_min)][pos[0] + 1 + abs(x_min)] == "x":
            map[pos[1] + 1 + abs(y_min)][pos[0] + 1 + abs(x_min)] = "X"
        else:
            map[pos[1] + 1 + abs(y_min)][pos[0] + 1 + abs(x_min)] = "y"

    map[1 + abs(y_min)][1 + abs(x_min)] = "o"

    map.reverse()
    for y in map:
        for x in y:
            print(x, end='')
        print()


def parse_positions(data):
    wire_positions = []
    for wire in data:
        positions = [
            [0, 0]
        ]
        for wire_point in wire.split(","):
            current = positions[-1].copy()
            if wire_point[0] == 'U':
                for i in range(1, int(wire_point[1:]) + 1):
                    positions.append([current[0], current[1] + i])
            elif wire_point[0] == 'D':
                for i in range(1, int(wire_point[1:]) + 1):
                    positions.append([current[0], current[1] - i])
            elif wire_point[0] == 'L':
                for i in range(1, int(wire_point[1:]) + 1):
                    positions.append([current[0] - i, current[1]])
            elif wire_point[0] == 'R':
                for i in range(1, int(wire_point[1:]) + 1):
                    positions.append([current[0] + i, current[1]])
        wire_positions.append(positions)
    return wire_positions


def part1(data):
    wire_positions = parse_positions(data)

    # print_map(wire_positions)

    min_distance = float('+inf')
    for pos in wire_positions[1]:
        if pos != [0, 0] and pos in wire_positions[0]:
            manhattan_dist = abs(pos[0]) + abs(pos[1])
            # print("Collission at", pos, "dist", manhattan_dist)
            if manhattan_dist < min_distance:
                min_distance = manhattan_dist

    print(int(min_distance))
    return int(min_distance)


def part2(data):
    wire_positions = parse_positions(data)

    # print_map(wire_positions)

    intersections = []
    for pos in wire_positions[0]:
        if pos != [0, 0] and pos in wire_positions[1]:
            intersections.append(pos)

    lowest_steps = float('+inf')
    for intersection in intersections:
        steps = wire_positions[0].index(intersection) + wire_positions[1].index(intersection)
        if steps < lowest_steps:
            lowest_steps = steps

    print(lowest_steps)
    return lowest_steps


def readFile(filename):
    return open(filename).read().strip().split("\n")


if __name__ == "__main__":
    assert part1(["R8,U5,L5,D3", "U7,R6,D4,L4"]) == 6
    assert part1(["R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"]) == 159
    # print(part1(readFile("input")))
    assert part2(["R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"]) == 610
    assert part2(["R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"]) == 410
    print(part2(readFile("input")))
