import math
import itertools


f = open("test.txt", "r")

lines = f.readlines()
points = []

for line in lines:
    x, y, z = map(int, line.strip().split(","))
    points.append((x, y, z))


# print(points)


def get_distances(points):
    distances = []
    pairs = itertools.combinations(points, 2)
    for pair in list(pairs):
        distance = math.dist(pair[0], pair[1])
        distances.append((pair[0], pair[1], distance))

    return distances


def part1(points):
    distances = get_distances(points)
    distances.sort(key=lambda x: x[2])

    max_connections = 10
    found_connections = 0
    circuits = []

    i = 0
    while found_connections < max_connections:
        p1, p2 = distances[i][0], distances[i][1]
        # check if we already have a dict for this point
        appended = False
        for circuit in circuits:
            if p1 in circuit and p2 in circuit:
                appended = True
                break
            elif p1 in circuit:
                print(f"{p1} found in circuit, appending {p2} - {circuit}")
                circuit.append(p2)
                appended = True
                found_connections += 1
                break
            elif p2 in circuit:
                print(f"{p2} found in circuit, appending {p1} - {circuit}")
                circuit.append(p1)
                found_connections += 1
                appended = True
                break
        if not appended:
            print(f"Creating new circuit with {p1} and {p2}")
            circuits.append([p1, p2])
            found_connections += 1
        i += 1

    for circuit in circuits:
        # print(f"Initial circuit: {circuit}")
        for point in circuit:
            for other_circuit in circuits:
                if circuit != other_circuit and point in other_circuit:
                    # merge circuits
                    for p in other_circuit:
                        if p not in circuit:
                            circuit.append(p)
                    circuits.remove(other_circuit)
                    break

    # sort circuits by length
    circuits.sort(key=lambda x: len(x), reverse=True)
    print(f"Lengths of circuits found: {len(circuits)}")
    for circuit in circuits:
        print(f"Circuit length: {len(circuit)} with points: {circuit}")

    print("Part 1:", (len(circuits[0]) * len(circuits[1]) * len(circuits[2])))


part1(points)
