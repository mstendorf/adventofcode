# from collections import defaultdict
#
# infile = "input"
# p1 = 0
# p2 = 0
# D = open(infile).read().strip()
#
# # E[x] is the set of pages that must come before x
# # ER[x] is the set of pages that must come after x
# E = defaultdict(set)
# ER = defaultdict(set)
# edges, queries = D.split("\n\n")
# for line in edges.split("\n"):
#     x, y = line.split("|")
#     x, y = int(x), int(y)
#     E[y].add(x)
#     ER[x].add(y)
#
# for query in queries.split("\n"):
#     vs = [int(x) for x in query.split(",")]
#     # assert len(vs) % 2 == 1
#     ok = True
#     for i, x in enumerate(vs):
#         for j, y in enumerate(vs):
#             if i < j and y in E[x]:
#                 ok = False
#     if ok:
#         p1 += vs[len(vs) // 2]
# print(p1)


from collections import defaultdict
from itertools import permutations

filename = "input"


def parse_input(filename):
    rules = defaultdict(set)
    pages = []
    with open(filename, "r") as f:
        e, q = f.read().strip().split("\n\n")
        for line in e.split("\n"):
            x, y = line.split("|")
            x, y = int(x), int(y)
            rules[y].add(x)

        for query in q.split("\n"):
            pages.append([int(x) for x in query.split(",")])

    return rules, pages


def validate(page, rules):
    for i, x in enumerate(page):
        for j, y in enumerate(page):
            if i < j and y in rules[x]:
                return False

    return True


def part1(rules, pages):
    count = 0
    for page in pages:
        if validate(page, rules):
            count += page[len(page) // 2]

    print(count)


def part2(rules, pages):
    count = 0

    for page in pages:
        if not validate(page, rules):
            perms = permutations(page)
            for perm in perms:
                print(perm)
                if validate(perm, rules):
                    count += perm[len(perm) // 2]
                    break

    print(count)


if __name__ == "__main__":
    before_map, pages = parse_input(filename)
    part1(before_map, pages)
    part2(before_map, pages)
