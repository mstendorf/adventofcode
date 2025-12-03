f = open("input.txt", "r")
lines = f.readlines()


def part1(lines):
    sum = 0
    for line in lines:
        line = line.strip()
        largest = 0
        for i in range(0, len(line) - 1):
            for j in range(i + 1, len(line)):
                cur = int(line[i] + line[j])

                if cur > largest:
                    largest = cur
        sum += largest

    print("Pa4rt 1:", sum)


def part2(lines):
    # though process being that we are taking the biggest digit possible every time
    # as long as we have enough digits left to complete the 12 digit number
    sum = 0
    for line in lines:
        line = line.strip()

        result = []
        start = 0
        n = len(line)

        for i in range(12):
            needed = 12 - i
            end = n - needed

            biggest_digit = max(line[start : end + 1])

            pos = line.index(biggest_digit, start, end + 1)
            result.append(biggest_digit)
            start = pos + 1

        largest = "".join(result)
        sum += int(largest)
    print("Part 2:", sum)


part1(lines)
part2(lines)
