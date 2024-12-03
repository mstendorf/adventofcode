file = open("input", "r")
map = [[x[:1] for x in line] for line in file.readlines()]


def count_energized(grid, start):
    # (row, col, movement row, movement col)
    queue = [start]
    seen = set()

    while queue:
        x, y, dx, dy = queue.pop(0)
        x += dx
        y += dy

        if x < 0 or x >= len(grid) or y < 0 or y >= len(grid[0]):
            continue

        new_pos = grid[x][y]

        if (
            new_pos == "."
            or (new_pos == "-" and dy != 0)
            or (new_pos == "|" and dx != 0)
        ):
            queue.append((x, y, dx, dy))
            seen.add((x, y, dx, dy))

        elif new_pos == "\\":
            dx, dy = dy, dx
            if (x, y, dx, dy) not in seen:
                queue.append((x, y, dx, dy))
                seen.add((x, y, dx, dy))

        elif new_pos == "/":
            dx, dy = -dy, -dx
            if (x, y, dx, dy) not in seen:
                queue.append((x, y, dx, dy))
                seen.add((x, y, dx, dy))

        else:
            for dr, dc in [(1, 0), (-1, 0)] if new_pos == "|" else [(0, 1), (0, -1)]:
                if (x, y, dr, dc) not in seen:
                    queue.append((x, y, dr, dc))
                    seen.add((x, y, dr, dc))

    visited = {(row, col) for (row, col, _, _) in seen}

    return len(visited)


# row -1 because its like you start outside the grid
solution1 = count_energized(map, (0, -1, 0, 1))
print("Solution 1", solution1 - 1)  # -1 because we start outside the grid
