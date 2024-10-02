import functools


@functools.cache
def calculate_arrangements(record, groups):
    # if we do not have any more groups we could be at a valid endpoint or not.
    if not groups:
        # if there is not more damaged springs we are at a valid endpoint
        if "#" not in record:
            return 1
        else:
            return 0

    # we still have groups but no where to put them
    if not record:
        return 0

    # Look at the next element in the record and gruop
    next_char = record[0]
    next_group = groups[0]

    arrangements = 0
    match next_char:
        case ".":
            # dots does not matter so just skip over them.
            arrangements += calculate_arrangements(record[1:], groups)
        case "#":
            # if first char is a pound we need to fit the entire group here
            cur_group = record[:next_group]
            cur_group = cur_group.replace("?", "#")

            if cur_group != "#" * next_group:
                # if the group does not match the group size we can not fit it here
                return 0

            # if the rest of the record is just the last group we are at a valid endpoint
            if len(record) == next_group:
                if len(groups) == 1:
                    # we have a valid solution
                    return 1
                else:
                    # there is not more groups but still damaged springs so this path doesnt work
                    return 0

            # if we get to here we have to have a possible dot following the group
            if record[next_group] in "?.":
                # since we know it has to be a dot we can skip if for this path even if it is a ?
                return calculate_arrangements(record[next_group + 1 :], groups[1:])

            # there is no possible path from here
            return 0
        case "?":
            # if we have a ? we can try both paths
            arrangements += calculate_arrangements("." + record[1:], groups)
            arrangements += calculate_arrangements("#" + record[1:], groups)

    return arrangements


with open("data.txt", "r") as f:
    sum = 0
    for line in f:
        record, groups = line.split(" ")
        groups = tuple(map(int, groups.split(",")))

        record = "?".join([record for _ in range(5)])
        groups = groups * 5

        sum += calculate_arrangements(record, groups)

    print(sum)
