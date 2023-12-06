def parse_line(line):
    a, b = line.split(":")[1].split("|")
    winners = [x for x in a.strip().split(" ") if x != ""]
    scratchcard = [x for x in b.strip().split(" ") if x != ""]
    return winners, scratchcard


with open("data.txt") as f:
    data = f.read().splitlines()
    sum = 0
    for line in data:
        ticket_sum = 0
        winners, scratchcard = parse_line(line)
        print(winners, scratchcard)
        for number in scratchcard:
            if number in winners:
                print("Found a winner: ", number)
                ticket_sum = 1 if ticket_sum == 0 else ticket_sum * 2
                print("Ticket sum: ", ticket_sum)
        sum += ticket_sum
        print("--------------------")
print(sum)
