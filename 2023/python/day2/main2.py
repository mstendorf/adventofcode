with open("data.txt") as file:
    powers = []
    for line in file.readlines():
        splits = line.split(":")
        rounds = splits[1].split(";")

        cubes = {"red": 0, "green": 0, "blue": 0}
        for round in rounds:
            dices = round.split(",")
            for dice in dices:
                parts = dice.strip().split(" ")
                amount = int(parts[0])
                color = parts[1]
                cubes[color] = max(cubes[color], amount)

        power = 1
        for _, value in cubes.items():
            power = value * power
        powers.append(power)
print(sum(powers))
