from heapq import heappop, heappush
from itertools import chain, combinations
from math import e, sqrt
import re
from typing import NamedTuple


class Pos(NamedTuple):
    x: int
    y: int
    h: float

    def __gt__(self, value) -> bool:
        return self.h > value.h

    @property
    def xy(self):
        return self.x, self.y


class Space(NamedTuple):
    mx: int
    my: int
    corrupted: set[tuple[int, int]]

    def corrupt(self, x, y):
        self.corrupted.add((x, y))

    def search(self):
        s = Pos(0, 0, self.dist(0, 0))
        qs = [s]
        ps = {}
        cs = {s: 0}

        while qs:

            p = heappop(qs)

            if p.xy == (self.mx - 1, self.my - 1):
                break

            neighbors = [*self.neighbours(*p.xy)]
            for n in neighbors:
                c = cs[p] + 1
                if n not in cs or cs[n] > c:
                    cs[n] = c
                    ps[n.xy] = p
                    heappush(qs, n)
        return p, ps

    def path(self):
        p, ps = self.search()

        path = []
        while p:
            path.append(p)
            p = ps.get(p.xy)

        return path

    def dist(self, x, y):
        return sqrt((x - self.mx) ** 2 + abs(y - self.my) ** 2)

    def in_bounds(self, x, y):
        return 0 <= x < self.mx and 0 <= y < self.my and (x, y) not in self.corrupted

    def neighbours(self, _x, _y):
        return (
            Pos(x, y, self.dist(x, y))
            for x, y in [
                (_x + 1, _y),
                (_x - 1, _y),
                (_x, _y + 1),
                (_x, _y - 1),
            ]
            if self.in_bounds(x, y)
        )


RE_POS = r"(\d+),(\d+)\n"


def part_one(filename, mx, my, c):
    with open(f"2024/day18/{filename}") as f:
        matches = re.findall(RE_POS, f.read())

    corrupted = set((int(x), int(y)) for x, y in matches[:c])

    s = Space(mx, my, corrupted)
    path = s.path()

    print(f"part one ({filename}): {len(path)-1}")


def part_two(filename, mx, my):
    with open(f"2024/day18/{filename}") as f:
        matches = re.findall(RE_POS, f.read())

    corrupted = [(int(x), int(y)) for x, y in matches]
    s = Space(mx, my, set())

    for i, c in enumerate(corrupted):
        print(i, c)
        s.corrupt(*c)
        _, ps = s.search()
        if not (70, 70) in ps:
            print(f"part two ({filename}): {c}")
            break


if __name__ == "__main__":
    # part_one("test1.txt", 7, 7, 12)
    part_one("input.txt", 71, 71, 1028)

    # part_two("test1.txt", 7, 7)
    part_two("input.txt", 71, 71)
