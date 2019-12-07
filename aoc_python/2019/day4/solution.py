def has_double(param):
    for char in param:
        if char + "" + char in param:
            return True
    return False


def has_double_part2(param):
    doubles = set()
    for char in param:
        char_combos = [char * x for x in range(2, len(param))]
        char_combos.reverse()
        for combo in char_combos:
            if combo in param:
                doubles.add(combo)
                break

    for double in doubles:
        if len(double) == 2:
            return True
    return False


def is_increasing(param):
    for i, char in enumerate(param):
        current = int(char)
        prev_pos = i - 1 if i > 0 else 0
        previous = int(param[prev_pos])
        if current < previous:
            return False
    return True


def part1(data):
    matches = []
    for i in range(int(data[0]), int(data[1])):
        containsdouble = has_double(str(i))
        increasing = is_increasing(str(i))
        if containsdouble and increasing:
            matches.append(matches)

    return len(matches)


def part2(data):
    matches = []
    for i in range(int(data[0]), int(data[1])):
        containsdouble = has_double_part2(str(i))
        increasing = is_increasing(str(i))
        if containsdouble and increasing:
            matches.append(matches)

    return len(matches)


def readFile(filename):
    return open(filename).read().strip().split("-")


if __name__ == "__main__":
    print(part1(readFile("input")))
    print(part2(readFile("input")))
