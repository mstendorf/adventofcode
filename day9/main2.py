def extrapolate(numbers):
    number_set = set(numbers)

    if len(number_set) == 1 and number_set.pop() == 0:
        return 0

    next = []
    for i in range(len(numbers) - 1):
        next.append(numbers[i + 1] - numbers[i])

    return numbers[0] - extrapolate(next)


with open("data.txt") as f:
    data = f.readlines()

    results = 0
    for line in data:
        numbers = list(map(int, line.split()))
        results += extrapolate(numbers)
    print(results)
