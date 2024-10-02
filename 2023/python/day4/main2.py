def parse_line(line):
    a, b = line.split(":")[1].split("|")
    winners = [x for x in a.strip().split(" ") if x != ""]
    scratchcard = [x for x in b.strip().split(" ") if x != ""]
    return winners, scratchcard


with open("data.txt") as f:
    data = f.read().splitlines()
    copies = {}
    sum = 0
    for i, line in enumerate(data):
        sum += 1 + copies.get(i, 0)
        winners, scratchcard = parse_line(line)
        card_sum = 0
        for number in scratchcard:
            if number in winners:
                card_sum += 1
                if i + card_sum not in copies:
                    copies[i + card_sum] = 0
                copies[i + card_sum] += 1 + copies.get(i, 0)

print(sum)
