# this is inspired by other solutions, initially i solved it by brute force wich was way too slow
with open("data.txt", "r") as f:
    data = f.read()
    seed_range, *blocks = data.split("\n\n")

    seed_range = list(map(int, seed_range.split(":")[1].split()))

    seeds = []

    for i in range(0, len(seed_range), 2):
        seeds.append((seed_range[i], seed_range[i] + seed_range[i + 1]))

    print(seeds)
    print("seeds done")
    for block in blocks:
        ranges = []
        for line in block.splitlines()[1:]:
            ranges.append(list(map(int, line.split())))

        next_identifiers = []
        while len(seeds) > 0:
            start, end = seeds.pop()
            for dst, src, length in ranges:
                # calculate the part of the seed range that overlaps with the current range
                overlap_start = max(start, src)
                overlap_end = min(end, src + length)
                if overlap_start < overlap_end:
                    next_identifiers.append(
                        (overlap_start - src + dst, overlap_end - src + dst)
                    )
                    # the seed range might be in the middle of the current range, so we need to add the parts before and after the seed range
                    # we simply do that by adding the ranges before and after the seed range to the seeds list
                    # and iterate over them in a later iteration
                    if overlap_start > start:
                        seeds.append((start, overlap_start))
                    if overlap_end < end:
                        seeds.append((overlap_end, end))
                    break
        # map seeds to the new values for the next block
        seeds = next_identifiers

print(min(seeds))
