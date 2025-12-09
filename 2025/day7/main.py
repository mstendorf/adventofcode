from functools import cache

f = open("input.txt", "r")


lines = f.readlines()

data = []
for line in lines:
    data.append(list(line.strip()))


def part1(lines):
    splits = 0
    for i in range(0, len(lines)):
        for j in range(0, len(lines[i])):
            if lines[i][j] == "S":
                lines[i + 1][j] = "|"
            if lines[i][j] == "^" and lines[i - 1][j] == "|":
                splits += 1
                lines[i][j - 1] = "|"
                lines[i][j + 1] = "|"
            if lines[i][j] == "." and lines[i - 1][j] == "|":
                lines[i][j] = "|"
    print("Part 1:", splits)


def part2(data):
    start = data[1].index("|")

    @cache
    def dfs(x, y):
        if y >= len(data):
            return 1

        if x < 0 or x >= len(data[0]):
            return 0

        if data[y][x] != "|":
            return 0

        if y + 1 < len(data):
            if data[y + 1][x] == "|":
                return dfs(x, y + 1)
            elif data[y + 1][x] == "^":
                return dfs(x - 1, y + 1) + dfs(x + 1, y + 1)
        else:
            # Reached the bottom
            return 1

        return 0

    print("Part 2:", dfs(start, 1))


part1(data)
part2(data)
