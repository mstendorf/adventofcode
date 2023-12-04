with open("./data.txt") as f:
    lines = f.read().splitlines()
    numbers = []
    for line in lines:
        line_numbers = []
        for c in line:
            if c.isdigit():
                line_numbers.append(c)

        numbers.append(int(line_numbers[0] + line_numbers[-1]))

print(sum(numbers))
