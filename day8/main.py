from itertools import cycle

with open("data.txt") as f:
    lines = f.readlines()
    path = lines[0].strip()

    data = {}
    for line in lines[2:]:
        parts = line.split("=")
        val = {}
        val["L"] = parts[1].split(",")[0].strip().replace("(", "")
        val["R"] = parts[1].split(",")[1].strip().replace(")", "").replace("\n", "")
        data[parts[0].strip()] = val

    count = 0
    current = data["AAA"]
    for c in cycle(path):
        count += 1
        if current[c] == "ZZZ":
            break
        current = data[current[c]]
    print(count)
