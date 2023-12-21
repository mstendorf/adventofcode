ITERATIONS = 1000000000
# ITERATIONS = 1000


def print_data(data):
    print()
    for line in data:
        print("".join(line))


def roll_north(data):
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char == "O":
                data[y][x] = "."
                i = y - 1
                while i >= 0 and data[i][x] == ".":
                    i -= 1
                data[i + 1][x] = "O"


def roll_south(data):
    data.reverse()
    roll_north(data)
    data.reverse()


def roll_west(data):
    for x in range(len(data[0])):
        for y in range(len(data)):
            if data[y][x] == "O":
                data[y][x] = "."
                i = x - 1
                while i >= 0 and data[y][i] == ".":
                    i -= 1
                data[y][i + 1] = "O"


def roll_east(data):
    for x in range(len(data[0]) - 1, -1, -1):
        for y in range(len(data)):
            if data[y][x] == "O":
                data[y][x] = "."
                i = x + 1
                while i < len(data[0]) and data[y][i] == ".":
                    i += 1
                data[y][i - 1] = "O"


def run_cycle(data):
    for direction in ["north", "west", "south", "east"]:
        match direction:
            case "north":
                roll_north(data)
            case "south":
                roll_south(data)
            case "west":
                roll_west(data)
            case "east":
                roll_east(data)


def get_boulder_positions(data):
    boulders = []
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char == "O":
                boulders.append((x, y))
    return tuple(boulders)


def calculate_beam_load(data):
    load = 0
    length = len(data[0])
    for i in range(length):
        boulders = data[i].count("O")
        boulder_load = boulders * (length - i)
        load += boulder_load

    return load


with open("data.txt", "r") as f:
    data = [[x for x in line.strip()] for line in f.readlines()]

    platform_states = {}
    cycles = 0
    while True:
        boulder_positions = get_boulder_positions(data)
        if boulder_positions in platform_states:
            cycle_length = cycles - platform_states[boulder_positions]
            # when we detect a cycle we skip ahead to last multiple of cycle_length before ITERATIONS
            print(f"{ITERATIONS} - ({ITERATIONS} - {cycles}) % {cycle_length}")
            cycles = ITERATIONS - (ITERATIONS - cycles) % cycle_length
            platform_states = {}
        platform_states[boulder_positions] = cycles
        run_cycle(data)

        cycles += 1
        if cycles == ITERATIONS:
            break
    print(calculate_beam_load(data))
