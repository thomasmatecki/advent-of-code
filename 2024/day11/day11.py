
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
        _stones = []
#        stones = chain.from_iterable(map(change, stones))
        for s in stones:
            _stones.extend(change(s))
        stones = _stones
    return stones


@cache
def depth(stone, i):
    l = 0
    for j in range(i):
        adv = breadth(stone, depth=15)
        stone, *tail = adv
        k = i-j - 1
        if k >= 1:
            r += iter_depth(tail, k)

    l += len(adv)

    return l


def iter_depth(stones, i):
    return sum(depth(stone, i) for stone in stones)


if __name__ == "__main__":
    with open("2024/day11/input.txt") as f:
        stones = [*map(int, f.read().split())]

    result = iter_depth(stones, 5)
    print(result)
