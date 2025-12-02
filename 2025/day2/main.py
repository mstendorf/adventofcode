f = open("input.txt", "r")

lines = f.readlines()
parts = lines[0].strip().split(",")
ranges = []
for part in parts:
    a, b = part.split("-")
    ranges.append((int(a), int(b)))


def part1(ranges):
    sum = 0
    for r in ranges:
        for i in range(r[0], r[1] + 1):
            tmp = str(i)
            if len(tmp) % 2 != 0:
                continue
            mid = len(tmp) // 2
            left = tmp[:mid]
            right = tmp[mid:]
            if left == right:
                sum += i

    print("Part 1:", sum)


def part2(ranges):
    sum = 0
    for r in ranges:
        for i in range(r[0], r[1] + 1):
            tmp = str(i)

            if tmp in (tmp + tmp)[1:-1]:
                sum += i

    print("Part 2:", sum)


part1(ranges)
part2(ranges)
