import numpy as np


def part1(data, width, height):
    data = [int(x) for x in data]

    num_layers = len(data) // (width * height)
    # print("datapoints", len(data), "Num layers", num_layers, "Layer size", (width * height))

    layers = []
    min_count_zeros = float('+inf')
    min_count_zeros_index = 0
    for layer_index in range(num_layers):
        layer = data[(width * height) * layer_index:(width * height) * (layer_index + 1)]
        num_zeros = layer.count(0)
        # print("Layer", layer_index, "count", num_zeros, layer)

        layers.append(layer)
        if num_zeros < min_count_zeros:
            min_count_zeros = num_zeros
            min_count_zeros_index = layer_index

    result_layer = layers[min_count_zeros_index]
    return result_layer.count(1) * result_layer.count(2)


def part2(data, width, height):
    data = [int(x) for x in data]

    num_layers = len(data) // (width * height)
    print("datapoints", len(data), "Num layers", num_layers, "Layer size", (width * height))

    layers = []
    for layer_index in range(num_layers):
        layer = data[(width * height) * layer_index:(width * height) * (layer_index + 1)]
        layer = np.array(layer).reshape((height, width))
        layers.append(layer)

    image = []
    for y in range(width):
        for x in range(height):
            pos_color = 2
            for layer in layers:
                color = layer[x][y]
                if color != 2:
                    pos_color = color
                    break
            image.append(pos_color)

    image = np.array(image).reshape(width, height)
    for y in range(len(image[0])):
        for x in range(len(image)):
            print("X" if image[x][y] == 1 else " ", end=' ')
        print()


def readFile(filename):
    return open(filename).read().strip().split("\n")[0]


if __name__ == "__main__":
    assert part1("123456789012", 3, 2) == 1
    assert part1(readFile("input"), 25, 6) == 1716
    print("Part2")
    part2("0222112222120000", 2, 2)
    part2("2222111100000000", 2, 2)
    part2("2222000000000000", 2, 2)
    part2(readFile("input"), 25, 6)
