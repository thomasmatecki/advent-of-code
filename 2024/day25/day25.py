from itertools import product


def part_one(filename):
    with open(f"2024/day25/{filename}") as f:
        gs = f.read().split("\n\n")
        specs = [l.split("\n") for l in map(str.strip, gs)]

    locks = []
    keys = []

    for spec in specs:
        cols = zip(*spec)
        if all(c == "#" for c in spec[0]):
            locks.append(tuple(sum(c == "#" for c in col)-1 for col in cols))
        else:
            keys.append(tuple(sum(c == "#" for c in col)-1 for col in cols))

    count = 0

    for lock, key in product(locks, keys):
        if all(l + k <= 5 for l, k in zip(lock, key)):
            count += 1

    print(f"part one {filename}: {count}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
