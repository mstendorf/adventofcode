from math import prod

# we will waste time finding the same numbers multiple times, and the set is a dirty hack to avoid that.
dirs = [(0, -1), (-1, -1), (-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1)]


def expand_to_full_number(line, index):
    left_index = index
    right_index = index
    while left_index >= 0 and line[left_index].isdigit():
        left_index -= 1
    while right_index < len(line) and line[right_index].isdigit():
        right_index += 1
    return int(line[left_index + 1 : right_index])


def get_adjacent_gear_sum(lines, line_index, char_index) -> int:
    # this works for this input but would break easily if there were 2 equal numbers in different positions.
    nums = set()
    for x, y in dirs:
        if line_index + x < 0 or line_index + x >= len(lines):
            continue
        if char_index + y < 0 or char_index + y >= len(lines[line_index + x]):
            continue
        if lines[line_index + x][char_index + y].isdigit():
            nums.add(expand_to_full_number(lines[line_index + x], char_index + y))

    return prod(nums) if len(nums) == 2 else 0


with open("data.txt", "r") as f:
    lines = f.readlines()
    sum = 0
    for i, line in enumerate(lines):
        print(i, line)
        num_start = -1
        length = 0
        for j, c in enumerate(line):
            if c == "*":
                sum += get_adjacent_gear_sum(lines, i, j)

    print(sum)
