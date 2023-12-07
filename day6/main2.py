def find_shortest_hold_time(time, distance):
    low = 0
    high = time
    while low < high:
        mid = (low + high) // 2
        print(low, high, mid)
        if mid * (time - mid) > distance:
            high = mid
        else:
            low = mid + 1

    return low


def find_longest_hold_time(time, distance):
    low = 0
    high = time
    while low < high:
        mid = (low + high) // 2
        if mid * (time - mid) > distance:
            low = mid + 1
        else:
            high = mid

    return low


with open("data.txt") as f:
    data = f.readlines()
    time = int(data[0].split(":")[1].replace(" ", ""))
    distance = int(data[1].split(":")[1].replace(" ", ""))

    # since the number is so big now, lets binary search our way to shortest and longest hold time
    shortest = find_shortest_hold_time(time, distance)
    longest = find_longest_hold_time(time, distance)
    print(longest - shortest)
