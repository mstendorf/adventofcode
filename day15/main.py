def to_hash(data):
    cv = 0
    for c in data:
        cv += ord(c)
        cv = (cv * 17) % 256
    return cv


with open("data.txt") as f:
    sum = 0
    for line in f:
        parts = line.strip().split(",")
        for elem in parts:
            sum += to_hash(elem)
    print(sum)
