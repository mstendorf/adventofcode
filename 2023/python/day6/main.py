from math import prod

with open("data.txt") as f:
    data = f.readlines()
    times = list(map(int, data[0].split(":")[1].split()))
    distances = list(map(int, data[1].split(":")[1].split()))

    round_win_options = []
    for time, distance in zip(times, distances):
        print(f"Race: {time} {distance}")
        speed = 0
        wins = 0
        for i in range(time + 1):
            traveled = (time - i) * i
            if traveled > distance:
                print(f"win: {i}  {traveled} {distance}")
                wins += 1

        round_win_options.append(wins)


print(prod(round_win_options))
