f = open("input.txt", "r")

lines = f.readlines()

ranges = []
items = []

ranges_part = True

for line in lines:
    if line.strip() == "":
        ranges_part = False
        continue

    if ranges_part:
        parts = line.strip().split("-")
        low, high = map(int, parts)
        ranges.append((low, high))
    else:
        num = int(line.strip())
        items.append(num)


def part1(ranges, items):
    sum = 0
    for item in items:
        for r in ranges:
            if r[0] <= item <= r[1]:
                sum += 1
                break

    print("Part 1:", sum)


def part2(ranges):
    merged = []
    for r in sorted(ranges):
        if not merged or merged[-1][1] < r[0] - 1:
            merged.append(r)
        else:
            merged[-1] = (merged[-1][0], max(merged[-1][1], r[1]))

    sum = 0
    for r in merged:
        sum += r[1] - r[0] + 1

    print("Part 2:", sum)


part1(ranges, items)
part2(ranges)
