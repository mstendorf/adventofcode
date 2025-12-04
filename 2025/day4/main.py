f = open("input.txt", "r")

lines = f.readlines()

grid = []

for line in lines:
    grid.append(list(line.strip()))


def check_neighbors(grid, i, j):
    directions = [(-1, 0), (-1, -1), (-1, 1), (1, 0), (1, 1), (1, -1), (0, -1), (0, 1)]
    count = 0
    for direction in directions:
        ni, nj = i + direction[0], j + direction[1]
        if 0 <= ni < len(grid) and 0 <= nj < len(grid[0]):
            if grid[ni][nj] == "@":
                count += 1
    return count < 4


def part1(grid):
    sum = 0
    for i in range(0, len(grid)):
        for j in range(0, len(grid[i])):
            cur = grid[i][j]
            if cur == "@":
                if check_neighbors(grid, i, j):
                    sum += 1
    print("Part 1:", sum)


def part2(grid):
    can_remove = True
    sum = 0
    while can_remove:
        cur_sum = 0
        for i in range(0, len(grid)):
            for j in range(0, len(grid[i])):
                cur = grid[i][j]
                if cur == "@":
                    if check_neighbors(grid, i, j):
                        cur_sum += 1
                        grid[i][j] = "."
        if cur_sum == 0:
            can_remove = False
        sum += cur_sum
    print("Part 2:", sum)


# print_grid(grid)
part1(grid)
part2(grid)
