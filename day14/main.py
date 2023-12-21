def print_data(data):
    for line in data:
        print("".join(line))


def roll_north(data, point):
    x, y = point
    data[x][y] = "."
    while x > 0 and data[x - 1][y] == ".":
        x -= 1

    data[x][y] = "O"


def tilt_lever_north(data):
    for i in range(len(data)):
        for j in range(len(data[i])):
            if data[i][j] == "O":
                roll_north(data, (i, j))


def calculate_beam_load(data):
    load = 0
    length = len(data[0])
    for i in range(length):
        boulders = data[i].count("O")
        load += boulders * (length - i)

    return load


with open("data.txt", "r") as f:
    data = [[x for x in line.strip()] for line in f.readlines()]
    tilt_lever_north(data)
    print(calculate_beam_load(data))
