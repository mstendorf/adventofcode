from itertools import cycle
from math import gcd

# I needed to lookup how to approach this, so the idea to use LCM is not mine.


def lcm(a):
    res = 1
    for x in a:
        res = (x * res) // gcd(x, res)
    return res


with open("data.txt") as f:
    lines = f.readlines()
    path = lines[0].strip()

    starters = []
    data = {}
    for line in lines[2:]:
        parts = line.split("=")
        val = {}
        val["L"] = parts[1].split(",")[0].strip().replace("(", "")
        val["R"] = parts[1].split(",")[1].strip().replace(")", "").replace("\n", "")
        key = parts[0].strip()
        data[key] = val
        if key[2] == "A":
            starters.append(key)

    walks = {}
    count = 0
    starter_length = len(starters)
    for c in cycle(path):
        count += 1
        temp_starters = []
        for i, key in enumerate(starters):
            if key.endswith("Z"):
                walks[i] = count - 1

            temp_starters.append(data[key][c])

        if len(starters) == len(walks):
            print(lcm(walks.values()))
            exit()

        starters = temp_starters
