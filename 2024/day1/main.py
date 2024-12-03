lines = [tuple(map(int, x.split())) for x in open("input").readlines()]

left = []
right = []
for line in open("input").readlines():
    l = line.split()
    l = list(map(int, l))
    left.append(l[0])
    right.append(l[1])

left.sort()
right.sort()


def part1():
    sum = 0
    for i in range(len(left)):
        sum += abs(left[i] - right[i])
    print(sum)


def part2():
    sum = 0
    count = 0
    current = 0
    for i in left:
        current = i
        for j in right:
            if i == j:
                count += 1
            if i < j:
                sum += count * current
                break

        count = 0
    print(sum)


part1()
part2()
