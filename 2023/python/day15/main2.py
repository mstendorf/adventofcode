def to_hash(data):
    cv = 0
    for c in data:
        cv += ord(c)
        cv = (cv * 17) % 256
    return cv


with open("data.txt") as f:
    boxes = {}
    for i in range(256):
        boxes[i] = {}

    for line in f:
        parts = line.strip().split(",")
        for elem in parts:
            box = 0
            label = ""
            focal_length = -1
            if "=" in elem:
                label, focal_length = elem.split("=")
                box = to_hash(label)
                boxes[box][label] = int(focal_length)
            elif "-" in elem:
                label = elem[:-1]
                box = to_hash(label)
                if label in boxes[box]:
                    del boxes[box][label]

    sum = 0
    for box, lenses in boxes.items():
        slot = 1
        for label, focal_length in lenses.items():
            sum += (box + 1) * slot * focal_length
            slot += 1
    print(sum)
