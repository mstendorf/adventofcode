def find_vertical_reflection(pattern):
    for i in range(len(pattern[0]) - 1):
        if pattern[0][i] == pattern[0][i + 1]:
            valid = True
            for j in range(len(pattern)):
                if pattern[j][i] != pattern[j][i + 1]:
                    valid = False

            if valid:
                return i + 1
    return 0


def find_horizontal_reflection(pattern):
    for i in range(len(pattern) - 1):
        if pattern[i] == pattern[i + 1]:
            valid = True
            for k in range(len(pattern)):
                if pattern[i][k] != pattern[i + 1][k]:
                    valid = False

            if valid:
                return i + 1


def find_reflection(pattern):
    vertical_reflection = find_vertical_reflection(pattern)
    if vertical_reflection != 0:
        return vertical_reflection

    return find_horizontal_reflection(pattern)


with open("datasample.txt", "r") as f:
    current_pattern = []
    for line in f:
        print(line)
        if line == "\n":
            print("New pattern")
            ref_point = find_reflection(current_pattern)
            print(ref_point)
            current_pattern = []
        else:
            current_pattern.append(line.strip())

    ref_point = find_reflection(current_pattern)
    print(ref_point)
