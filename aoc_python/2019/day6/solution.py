def get_indirect_orbits(orbits, source):
    if len(orbits.get(source, [])) > 0:
        indirect_orbits = 0
        for orbiter in orbits[source]:
            indirect_orbits += get_indirect_orbits(orbits, orbiter)
        return len(orbits[source]) + indirect_orbits
    else:
        return 0


def part1(data):
    orbits = {"COM": []}

    for row in data:
        source, orbiter = row.split(")")
        if source not in orbits:
            orbits[source] = []
        orbits[source].append(orbiter)

    direct_orbits = indirect_orbits = 0
    for source, orbiters in orbits.items():
        direct_orbits += len(orbiters)
        indirect_orbits += get_indirect_orbits(orbits, source)

    indirect_orbits -= direct_orbits

    return direct_orbits, indirect_orbits


def readFile(filename):
    return open(filename).read().strip().split("\n")


def find_transfer_path(orbits, current, target, path=[]):
    # Next time do bfs
    path = path + [current]

    if current == target:
        return path

    possible_nexts = [x for x in orbits[current] if x not in path]
    for orbiter in possible_nexts:
        if orbiter not in path:
            found = find_transfer_path(orbits, orbiter, target, path)
            if found:
                return found

    return None


def part2(data):
    orbits = {"COM": []}

    for row in data:
        source, orbiter = row.split(")")
        if source not in orbits:
            orbits[source] = []
        orbits[source].append(orbiter)
        if orbiter not in orbits:
            orbits[orbiter] = []
        orbits[orbiter].append(source)

    shortest_path = find_transfer_path(orbits, "YOU", "SAN")
    print("Shortest path", shortest_path)
    shortest_path = len(shortest_path) - 3
    return shortest_path


if __name__ == "__main__":
    assert (part1(["COM)B", "B)C", "C)D", "D)E", "E)F", "B)G",
                   "G)H", "D)I", "E)J", "J)K", "K)L", ]) == (11, 31))

    print(part1(readFile("input")))

    assert (part2(["COM)B", "B)C", "C)D", "D)E", "E)F", "B)G",
                   "G)H", "D)I", "E)J", "J)K", "K)L",
                   "K)YOU", "I)SAN", ]) == 4)

    print(part2(readFile("input")))
