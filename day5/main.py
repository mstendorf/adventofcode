with open("data.txt", "r") as f:
    data = f.read()
    seeds, *blocks = data.split("\n\n")

    seeds = list(map(int, seeds.split(":")[1].split()))

    for block in blocks:
        ranges = []
        for line in block.splitlines()[1:]:
            ranges.append(list(map(int, line.split())))

        next_identifiers = []
        for seed in seeds:
            for dst, src, length in ranges:
                if src <= seed < src + length:
                    # calculate the offset in this range
                    next_identifiers.append(seed - src + dst)
                    break
            else:
                # if the seed wasn't in any range, it doesn't change
                next_identifiers.append(seed)

        # map seeds to the new values for the next block
        seeds = next_identifiers

print(min(seeds))
