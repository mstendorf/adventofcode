def expand_universe(data):
    expand_universe_horizontal(data)
    expand_universe_vertical(data)


def expand_universe_horizontal(data):
    data_tmp = data.copy()
    diff = 0

    for i, c in enumerate(data_tmp[0]):
        if c == ".":
            expand = True
            for line in data_tmp:
                if line[i] != ".":
                    expand = False
                    break

            if expand:
                for j, line in enumerate(data):
                    data[j] = line[: i + diff] + "." + line[i + diff :]
                diff += 1


def expand_universe_vertical(data):
    length = len(data)
    i = 0
    while i < length:
        line = data[i]
        if set(line) == {"."}:
            data.insert(i + 1, line)
            i += 1
        i += 1


def find_all_pound(data):
    indexes = []
    for i, line in enumerate(data):
        for j, c in enumerate(line):
            if c == "#":
                indexes.append((i, j))
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
        # print(f"Distance between {idx_list[0]} and {point} is {distance}")
        sum += distance
    sum += find_distance_sum(idx_list[1:])
    return sum


with open("datasample.txt", "r") as f:
    data = f.read().split()
    expand_universe(data)
    indexes = find_all_pound(data)

    print(f"{len(data[0])} {len(data)}")
    print(find_distance_sum(indexes))
