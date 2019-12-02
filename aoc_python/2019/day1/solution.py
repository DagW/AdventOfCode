def readFile(filename):
    return open(filename).read().strip().split("\n")


def part1(data):
    sum = 0
    for mass in data:
        sum += int(int(mass) / 3) - 2
    return sum


def part2(data):
    sum = 0

    for row in data:
        sum_fuel = 0

        fuel = int(row)
        while fuel > 0:
            fuel = part1([fuel])
            if fuel > 0:
                sum_fuel += fuel

        sum += sum_fuel
    return sum


if __name__ == "__main__":
    assert part1(["12"]) == 2
    assert part2(["12"]) == 2
    assert part2(["14"]) == 2
    assert part2(["1969"]) == 966
    assert part2(["100756"]) == 50346

    print(part1(readFile("input")))
    print(part2(readFile("input")))
