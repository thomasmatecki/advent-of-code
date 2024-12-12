from functools import cache
from itertools import chain


def change(n):
    str_n = str(n)
    length = len(str_n)
    if n == 0:
        return (1,)
    elif length % 2 == 0:
        half = length // 2
        return (int(str_n[:half]), int(str_n[half:]))
    else:
        return (n * 2024, )


@cache
def breadth(stone, depth):
    stones = [stone]
    for _ in range(depth):
        stones = chain.from_iterable(change(s) for s in stones)
    return tuple(stones)


@cache
def depth(stone, i):
    l = 0
    for j in range(i, 0, -1):
        adv = breadth(stone, depth=3)
        if (k := j-1) >= 1:
            stone, *tail = adv
            l += iter_depth(tail, k)

    l += len(adv)
    return l


def iter_depth(stones, i):
    return sum(depth(stone, i) for stone in stones)


if __name__ == "__main__":
    with open("2024/day11/input.txt") as f:
        stones = [*map(int, f.read().split())]

    result = iter_depth(stones, 25)
    print(result)
