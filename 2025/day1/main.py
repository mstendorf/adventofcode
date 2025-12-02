f = open("input.txt", "r")
start = 50


def part1(start):
    count = 0
    for line in f:
        direction = line[0]
        value = int(line[1:])

        if direction == "L":
            start = abs((start - value) % 100)
        else:
            start = abs((start + value) % 100)

        if start == 0:
            count += 1

    print(count)


def part2(start):
    count = 0
    for line in f:
        direction = line[0]
        value = int(line[1:])

        if direction == "L":
            for i in range(1, value + 1):
                tmp = abs((start - i) % 100)
                if tmp == 0:
                    print("passing 0 - left value:", i)
                    count += 1
            start = tmp
        else:
            for i in range(1, value + 1):
                tmp = abs((start + i) % 100)
                if tmp == 0:
                    print("passing 0 - right value:", i)
                    count += 1
            start = tmp

    print(count)


# part1(start)
part2(start)
