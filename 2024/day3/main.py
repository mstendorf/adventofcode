import re

input = open("input").read()


def part1():
    regex = re.compile(r"mul\((\d+),(\d+)\)")
    sum = 0
    for match in regex.finditer(input):
        a, b = map(int, match.groups())

        sum += a * b

    print(sum)


def part2():
    # my initial attempt at using a regex did not work, so I had to clean the input manually
    do_parts = input.split("do()")
    input_cleaned = ""
    for part in do_parts:
        parts = part.split("don't()")
        input_cleaned += parts[0]

    regex = re.compile(r"mul\((\d+),(\d+)\)")
    sum = 0
    for match in regex.finditer(input_cleaned):
        a, b = map(int, match.groups())

        sum += a * b

    print(sum)


part1()
part2()
