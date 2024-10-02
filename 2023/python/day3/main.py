def validate_number(line_number, start, end, lines):
    line_above = max(0, line_number - 1)
    line_below = min(len(lines) - 1, line_number + 1)

    left_index = max(0, start - 1)
    right_index = min(len(lines[line_number]) - 1, end)

    if left_index >= 0:
        left_char = lines[line_number][left_index]
        if left_char != "." and not left_char.isdigit():
            return True

    if right_index <= len(lines[line_number]) - 1:
        right_char = lines[line_number][right_index]
        if right_char != "." and right_char != "\n" and not right_char.isdigit():
            return True

    if line_above != 0:
        if validate_adjecent_line(lines[line_above], left_index, right_index):
            return True

    if line_below != len(lines) - 1:
        if validate_adjecent_line(lines[line_below], left_index, right_index):
            return True

    return False


def validate_adjecent_line(line, start, end):
    for i in range(start, end + 1):
        if line[i] != "." and line[i] != "\n" and not line[i].isdigit():
            return True

    return False


with open("data.txt", "r") as f:
    lines = f.readlines()
    sum = 0
    for i, line in enumerate(lines):
        print(i, line)
        num_start = -1
        length = 0
        for j, c in enumerate(line):
            if c.isdigit():
                if num_start == -1:
                    num_start = j
                length += 1
            else:
                if num_start != -1:
                    if validate_number(i, num_start, num_start + length, lines):
                        cur_num = int(line[num_start : num_start + length])
                        sum += cur_num

                    num_start = -1
                    length = 0

    print(sum)
