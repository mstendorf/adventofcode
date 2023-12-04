with open("data.txt") as file:
    possible_games = []
    for line in file.readlines():
        print("-", line)
        splits = line.split(":")
        game_id = int(splits[0].split(" ")[1])
        rounds = splits[1].split(";")

        game_valid = True

        for round in rounds:
            colors = {}
            colors["red"] = 12
            colors["green"] = 13
            colors["blue"] = 14
            dices = round.split(",")
            for dice in dices:
                parts = dice.strip().split(" ")
                amount = int(parts[0])
                color = parts[1]
                colors[color] -= amount
                if colors[color] < 0:
                    game_valid = False
                    break
            if not game_valid:
                break
        if game_valid:
            print("Game", game_id, "is valid")
            possible_games.append(game_id)
        else:
            print("Game", game_id, "is invalid")
print(sum(possible_games))
