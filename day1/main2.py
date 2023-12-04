import re

digits = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
with open("./data.txt") as f:
    lines = f.read().splitlines()
    numbers = []
    for line in lines:
        line_numbers = {}
        for i, c in enumerate(line):
            if c.isdigit():
                line_numbers[i] = c
        for digit in digits:
            if digit in line:
                indexes = [m.start() for m in re.finditer(digit, line)]
                cur_num = str(digits.index(digit) + 1)
                for index in indexes:
                    line_numbers[index] = cur_num

        indexes = sorted(list(line_numbers.keys()))
        numbers.append(int(line_numbers[indexes[0]] + line_numbers[indexes[-1]]))

print(sum(numbers))
