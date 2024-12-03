levels = [list(map(int, x.split())) for x in open("input").readlines()]


def is_safe(level):
    direction = 0
    for i in range(len(level) - 1):
        diff = level[i + 1] - level[i]
        if diff > 0:
            if direction == -1:
                return False
            direction = 1
        elif diff < 0:
            if direction == 1:
                return False
            direction = -1

        if not (0 < abs(diff) <= 3):
            return False

    return True


def try_remove(level):
    for i in range(len(level)):
        temp = level.copy()
        temp.pop(i)
        if is_safe(temp):
            return True

    return False


def part1():
    count = 0
    for level in levels:
        if is_safe(level):
            count += 1

    print(count)


def part2():
    count = 0
    for level in levels:
        if is_safe(level):
            count += 1
        else:
            if try_remove(level):
                count += 1

    print(count)


if __name__ == "__main__":
    part1()
    part2()
