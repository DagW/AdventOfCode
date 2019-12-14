from day13.intcomp import IntcodeComputer, Halted


def part1(program):
    c = IntcodeComputer(program)
    try:
        outputs = c.run([])
    except Halted as e:
        outputs = e.last_output
    map = [[]]
    for pos in range(len(outputs) // 3):
        base_pos = pos * 3
        x = outputs[base_pos]
        y = outputs[base_pos + 1]
        id = outputs[base_pos + 2]
        if len(map) <= y:
            map.append([])
        map[y].append(id)

    numblocks = 0
    for y in map:
        for x in y:
            character = ' '
            if x == 1:
                character = "X"
            elif x == 2:
                character = "□"
                numblocks += 1
            elif x == 3:
                character = "-"
            elif x == 4:
                character = "o"
            print(character, end='')
        print()
    return numblocks


def createMap(outputs):
    map = ""
    for pos in range(len(outputs) // 3):
        base_pos = pos * 3
        x = outputs[base_pos]
        y = outputs[base_pos + 1]
        id = outputs[base_pos + 2]
        if x == -1:
            print("Score=", id)
        else:
            if y >= len(map.split("\n")):
                map += "\n"
            map += str(id)

    print(
        map.replace("0", " ").replace("1", "X")
            .replace("2", "□")
            .replace("3", "-")
            .replace("4", "o")
    )
    return map


def updateGame(map, outputs, print_map=False):
    ball_x = -1
    paddle_x = -1
    for pos in range(len(outputs) // 3):
        base_pos = pos * 3
        x = outputs[base_pos]
        y = outputs[base_pos + 1]
        id = outputs[base_pos + 2]
        if x == -1:
            print("Score", id)
            pass
        else:
            if id == 3:
                paddle_x = x
            elif id == 4:
                ball_x = x

            # original_char = map.split("\n")[y][x]
            # print("Replacing '{}'".format(original_char), "with", id)
            rows = map.split("\n")
            rows[y] = rows[y][:x] + str(id) + rows[y][x + 1:]  # [x] = id
            map = "\n".join(rows)

    if print_map:
        print(
            map.replace("0", " ").replace("1", "X")
                .replace("2", "□")
                .replace("3", "-")
                .replace("4", "o")
        )

    return ball_x, paddle_x, map


def part2(program):
    program[0] = 2
    c = IntcodeComputer(program)

    outputs = c.run([])
    map = createMap(outputs)
    map = map.replace("4", "0").replace("3", "0")
    c.clear_outputs()

    direction = -1
    # for i in range(0, 50):
    #   time.sleep(0.3)
    while True:
        try:
            # print(chr(27) + "[2J")  # Clear screen
            # print("-- Round --")
            outputs = c.run([direction])
            ball_x, paddle_x, map = updateGame(map, outputs, print_map=True)
            c.clear_outputs()
            if ball_x > paddle_x:
                direction = 1
            elif ball_x < paddle_x:
                direction = -1
            else:
                direction = 0
        except Halted as e:
            print("Game ended", e)
            break
    return ""


def readFile(filename):
    current_array = open(filename).read().strip().split(",")
    return [int(numeric_string) for numeric_string in current_array]


if __name__ == "__main__":
    print("Part1")
    program = readFile("input")
    print(part1(program.copy()))

    print("Part2")
    print(part2(program))
