from functools import cache
from itertools import chain


@cache
def change(n):
    str_n = str(n)
    length = len(str_n)
    if n == 0:
        return (1,)
    elif length % 2 == 0:
        half = length // 2
        return (int(str_n[:half]), int(str_n[half:]))
    else:
        return (n * 2024,)


@cache
def breadth(stone, depth):
    stones = [stone]
    for _ in range(depth):
        stones = chain.from_iterable(map(change, stones))
    return tuple(stones)


@cache
def depth(stone, i):
    l = 0
    for j in range(i, 0, -1):
        adv = breadth(stone, depth=5)
        if (k := j - 1) >= 1:
            stone, tail = adv[0], adv[1:]
            l += iter_depth(tail, k)

    l += sum(1 for _ in adv)
    return l


@cache
def iter_depth(stones, i):
    return sum(depth(stone, i) for stone in stones)


def load(filename):
    with open(f"2024/day11/{filename}") as f:
        return tuple(map(int, f.read().split()))


def part_one():
    stones = load("input.txt")
    result = iter_depth(stones, 4)
    print(result)


def part_two():
    stones = load("input.txt")
    result = iter_depth(stones, 15)
    print(result)


if __name__ == "__main__":
    part_one()
    part_two()
