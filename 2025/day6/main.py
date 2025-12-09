import math


f = open("input.txt", "r")

data = []


for line in f:
    data.append(line.strip().split())


def part1(data):
    # flip the array
    data = list(map(list, zip(*data)))

    res = 0

    for row in data:
        action = row[-1]
        parsed_row = list(map(int, row[:-1]))
        match action:
            case "+":
                res += sum(parsed_row)
            case "*":
                res += math.prod(parsed_row)
            case default:
                print("Unknown action:", action)
    print("Part 1:", res)


def part2():
    f = open("input.txt", "r")
    lines = f.readlines()
    data = list(map(list, zip(*lines)))
    element = ""
    numbers = []
    action = ""
    res = 0
    for i in range(0, len(data)):
        operator = data[i][-1].strip()
        if operator != "":
            action = operator
        for j in range(0, len(data[i]) - 1):
            element += data[i][j].strip()

        if len(element) == 0:
            parsed_numbers = list(map(int, numbers))
            match action:
                case "+":
                    res += sum(parsed_numbers)
                case "*":
                    res += math.prod(parsed_numbers)
            element = ""
            action = ""
            numbers = []

        if len(element) > 0:
            numbers.append(int(element))
            element = ""

    print("Part 2:", res)

    # for i in length(lines[0]):
    #     for j in range(len(lines)):
    #


part1(data)
part2()
