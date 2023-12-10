from collections import Counter
import functools

card_values = {
    "J": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "Q": 11,
    "K": 12,
    "A": 13,
    "pair": 14,
    "2pair": 15,
    "three": 16,
    "full": 17,
    "four": 18,
    "five": 19,
}


def five_of_a_kind(count):
    if ("J" in count and len(count) == 2) or (len(count) == 1 and "J" in count):
        return True

    return 5 in count.values()


def four_of_a_kind(count):
    js = count.pop("J", 0)
    missing = 4 - js
    if missing in count.values():
        return True

    if js > 0:
        count["J"] = js
    return 4 in count.values()


def full_house(count):
    js = count["J"] if "J" in count else 0
    if js == 1 and len(count) == 3 and 2 in count.values():
        return True
    return 3 in count.values() and 2 in count.values()


def three_of_a_kind(count):
    js = count["J"] if "J" in count else 0
    if js == 1 and len(count) == 4 or js == 2 and len(count) == 4:
        return True
    return 3 in count.values() and 2 not in count.values()


def two_pair(count):
    return len(count) == 3 and 2 in count.values()


def pair(count):
    js = count["J"] if "J" in count else 0
    if js == 1 and len(count) == 5:
        return True

    return len(count) == 4 and 2 in count.values()


def get_hand_value(hand):
    count = Counter(hand)

    if five_of_a_kind(count):
        return card_values["five"]

    if four_of_a_kind(count):
        return card_values["four"]

    if full_house(count):
        return card_values["full"]

    if three_of_a_kind(count):
        return card_values["three"]

    if two_pair(count):
        return card_values["2pair"]

    if pair(count):
        return card_values["pair"]

    return 1


def compare_cards(hand1, hand2):
    if hand1[2] > hand2[2]:
        return 1
    elif hand1[2] < hand2[2]:
        return -1
    else:
        for i in range(5):
            h1 = card_values[hand1[0][i]]
            h2 = card_values[hand2[0][i]]
            if h1 > h2:
                return 1
            elif h1 < h2:
                return -1
    return 0


with open("data.txt", "r") as f:
    data = f.readlines()

    cards = []
    for line in data:
        parts = line.split()
        hand = parts[0]
        bid = int(parts[1])
        value = get_hand_value(hand)
        cards.append((hand, bid, value))

    cards = sorted(cards, key=functools.cmp_to_key(compare_cards))
    for card in cards:
        if card[2] == 16 and "J" in card[0] and card[0].count("J") == 2:
            print(card)
    sum = 0
    for i, card in enumerate(cards):
        tsum = (i + 1) * card[1]
        sum += tsum
    print(sum)
