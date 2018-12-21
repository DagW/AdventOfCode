
def rotate(arr, index):
    return arr[index:] + arr[:index]

def part1(inputArray):
    # https://oeis.org/A006257
    players = inputArray[0]
    lastMarbleWorth = inputArray[1]
    result = [0]
    position = 0
    scores = [0]*players
    #print(scores)
    #print(0, result)
    for marble in range(1,lastMarbleWorth+1):
        player = marble%players
        if marble % 23 == 0:
            print("Special")
            scores[player] += marble
            #remove 7 marbles to the left
            scores[player] += result.pop((position-6) % len(result))
        else:
            position = (position+2)%len(result)
            result.insert(position+1, marble)
        #print(player, result)
        
    print(scores)
    return 5


if __name__ == "__main__":
    part1([10,1618])
    """assert part1([10,1618]) == 8317
    assert part1([13,7999]) == 146373
    assert part1([17,1104]) == 2764
    assert part1([21,6111]) == 54718
    assert part1([30,5807]) == 37305

    inputData = readFile("input")
    print(part1([491, 71058]))"""
