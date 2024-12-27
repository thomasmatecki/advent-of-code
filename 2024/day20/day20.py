from collections import UserList, defaultdict
from dataclasses import dataclass, field
from email.policy import default
from functools import cached_property
from typing import Any


@dataclass
class Path:
    t: tuple[tuple[int, int]] = field(default_factory=tuple)
    d: dict[tuple[int, int], int] = field(default_factory=dict)

    def __contains__(self, item):
        return item in self.d

    def __len__(self):
        return len(self.t)

    def __iter__(self):
        return iter(self.t)

    def addend(self, item):
        t = (*self.t, item)
        d = {p: i for i, p in enumerate(t)}

        return Path(t, d)  # type: ignore

    @cached_property
    def end(self):
        return self.t[-1]

    def idx(self, i, j):
        return self.d[(i, j)]


class Racetrack(UserList):

    @classmethod
    def from_file(cls, filename):
        with open(f"2024/day20/{filename}", "r") as f:
            return cls([*map(str.strip, f.readlines())])

    @cached_property
    def start(self):
        for i, row in enumerate(self):
            for j, v in enumerate(row):
                if v == "S":
                    return i, j
        return -1, 1

    @cached_property
    def end(self):
        for i, row in enumerate(self):
            for j, v in enumerate(row):
                if v == "E":
                    return i, j
        return -1, 1

    def in_bounds(self, i: int, j: int):
        return 0 <= i < len(self) and 0 <= j < len(self[0])

    def neighbors(self, i: int, j: int):
        for di, dj in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            if self.in_bounds(i + di, j + dj):
                yield i + di, j + dj

    def dfs(self):
        p = Path().addend(self.start)

        while p.end != self.end:
            for n in self.neighbors(*p.end):
                i, j = n
                if self[i][j] in (".", "E") and n not in p:
                    p = p.addend(n)
                    break

        return p

    def two_cheat(self, path):
        c = defaultdict(int)
        for i, j in path:
            for di, dj in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
                i2, j2 = i + (2 * di), j + (2 * dj)
                if (i2, j2) in path and self[i + di][j + dj] == "#":
                    d = path.idx(i2, j2) - path.idx(i, j) - 2
                    if d > 0:
                        c[d] += 1
        return c

    def cheat_zone(self, p, l, path):
        i0, j0 = p
        for i in range(i0 - l, i0 + l + 1):
            k = l - abs(i0 - i)
            for j in range(j0 - k, j0 + k + 1):
                if self.in_bounds(i, j) and self[i][j] in (".", "E"):
                    yield i, j


def mdist(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def part_one(filename):

    track = Racetrack.from_file(filename)
    path = track.dfs()
    c = track.two_cheat(path)
    s = sum(c for s, c in c.items() if s >= 100)
    print(f"part one ({filename}): {s}")


def part_two(filename, m):

    track = Racetrack.from_file(filename)
    path = track.dfs()
    c = defaultdict(int)
    for p in path:
        for z in track.cheat_zone(p, 20, path):
            s = path.idx(*z) - path.idx(*p) - mdist(p, z)
            if s >= m:
                c[s] += 1

    s = sum(c for s, c in c.items())
    print(f"part two ({filename}): {s}")


if __name__ == "__main__":
    part_one("input.txt")
    part_two("test1.txt", 50)
    part_two("input.txt", 100)
