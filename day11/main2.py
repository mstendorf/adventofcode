EXPANSION_RATE = 1_000_000 - 1


def expand_vertical(data):
    empty_lines = 0
    for y, line in enumerate(data):
        if len(set(line)) > 1:
            data[y] = [
                (x, y + empty_lines * EXPANSION_RATE) if n != "." else "."
                for x, n in enumerate(line)
            ]
        else:
            empty_lines += 1


def expand_horizontal(data):
    empty_lines = 0
    for x in range(len(data[0])):
        line = [data[y][x] for y in range(len(data))]
        if len(set(line)) > 1:
            for y in range(len(data)):
                data[y][x] = (
                    (data[y][x][0] + empty_lines * EXPANSION_RATE, data[y][x][1])
                    if data[y][x] != "."
                    else "."
                )
        else:
            empty_lines += 1


def find_all_galaxies(data):
    indexes = []
    for line in data:
        for c in line:
            if c != ".":
                indexes.append(c)
    return indexes


def find_distance_between_points(point1, point2):
    return abs(point1[0] - point2[0]) + abs(point1[1] - point2[1])


def find_distance_sum(idx_list):
    sum = 0
    if len(idx_list) == 1:
        return 0

    current = idx_list[0]
    for point in idx_list[1:]:
        distance = find_distance_between_points(current, point)
        sum += distance
    sum += find_distance_sum(idx_list[1:])
    return sum


with open("data.txt", "r") as f:
    data = [
        [(x, y) if c == "#" else "." for x, c in enumerate(line.strip())]
        for y, line in enumerate(f)
    ]

    expand_vertical(data)
    expand_horizontal(data)
    galaxies = find_all_galaxies(data)
    print(find_distance_sum(galaxies))
