filename = "input"
input = open(filename, "r").read().split("\n")[:-1]


def count_xmas(s):
    return s.count("XMAS") + s.count("SAMX")


def columns(input):
    return [
        "".join([input[i][j] for i in range(len(input))]) for j in range(len(input[0]))
    ]


def diagonals(input):
    size = len(input)
    out = []
    for x in range(size):
        out.append("".join(input[i][x + i] for i in range(size - x)))
        out.append("".join(input[size - 1 - i][x + i] for i in range(size - x)))
    for y in range(1, size - 1):
        out.append("".join(input[y + i][i] for i in range(size - y)))
        out.append("".join(input[y - i][i] for i in range(y + 1)))

    return out


def print_grid(input):
    for i in input:
        for j in i:
            print(j, end="")
        print()


def part1(input):
    count = 0
    for i in input:
        count += count_xmas(i)
    for i in columns(input):
        count += count_xmas(i)
    for i in diagonals(input):
        count += count_xmas(i)
    print(count)


def part2(input):
    count = 0
    for x in range(1, len(input) - 1):
        for y in range(1, len(input[0]) - 1):
            if input[x][y] == "A":
                diag1 = input[x - 1][y - 1] + input[x + 1][y + 1]
                diag2 = input[x - 1][y + 1] + input[x + 1][y - 1]
                if diag1 in ("SM", "MS") and diag2 in ("SM", "MS"):
                    count += 1

    print(count)


part1(input)
part2(input)
