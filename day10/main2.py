dirs = {
    "U": (-1, 0),
    "D": (1, 0),
    "L": (0, -1),
    "R": (0, 1),
}
pipes = {
    "|": {"U": "U", "D": "D"},
    "-": {"L": "L", "R": "R"},
    "L": {"D": "R", "L": "U"},
    "J": {"D": "L", "R": "U"},
    "7": {"U": "L", "R": "D"},
    "F": {"U": "R", "L": "D"},
    "S": {"U": "U", "D": "D", "L": "L", "R": "R"},
}


def find_start(data):
    for i, line in enumerate(data):
        idx = line.index("S") if "S" in line else -1
        if idx != -1:
            return (i, idx)
    return (-1, -1)  # just to make the linter shut up, we should always have an S


def find_starting_direction(start):
    directions = []
    for key, direction in dirs.items():
        next_pipe = data[start[0] + direction[0]][start[1] + direction[1]]
        if next_pipe in pipes and key in pipes[next_pipe].keys():
            return key

    return directions.pop()


def count_insides_in_line(line):
    count = 0
    for i, char in enumerate(line):
        if char == ".":
            if ray_trace(line[i:]):
                line[i] = "I"
                count += 1
    return count


def ray_trace(line):
    intersections = 0
    for char in line:
        if char in ("|", "F", "7", "S"):
            intersections += 1
    return intersections % 2 == 1


def surf(start, direction):
    last_pipe = None
    count = 0
    while last_pipe != "S":
        count += 1
        start = (start[0] + dirs[direction][0], start[1] + dirs[direction][1])
        last_pipe = data[start[0]][start[1]]
        data_out[start[0]][start[1]] = last_pipe
        direction = pipes[last_pipe][direction]


with open("data.txt", "r") as f:
    data = [[*x] for x in f.readlines()]
    data_out = [["." for _ in range(len(data[0]))] for _ in range(len(data))]
    start = find_start(data)
    direction = find_starting_direction(start)
    surf(start, direction)
    sum = 0
    for line in data_out:
        sum += count_insides_in_line(line)
        # print("".join(line))
    print(sum)
