from collections import UserList
from functools import reduce
from itertools import count
import operator
import re
from tokenize import Name
from typing import Counter, NamedTuple, SupportsIndex


INPUT_RE = r"p=(\d+),(\d+) v=(-?\d+),(-?\d+)\n"


class V(NamedTuple):
    x: int
    y: int

    def __mul__(self, value: int) -> "V":
        return V(self.x * value, self.y * value)


class P(NamedTuple):
    x: int
    y: int

    def __mod__(self, v) -> "P":
        x, y = v
        return P(self.x % x, self.y % y)

    def move(self, v) -> "P":
        return P(self.x + v.x, self.y + v.y)


def load(filename):
    with open(f"2024/day14/{filename}") as f:
        matches = re.findall(INPUT_RE, f.read())
        matches = (map(int, m) for m in matches)
        return [(P(px, py),  V(vx, vy)) for px, py, vx, vy in matches]


class Grid(UserList):
    def __init__(self, pvs, xb, yb):

        super().__init__(pvs)
        self.xb = xb
        self.yb = yb

    def tick(self, t):
        moved_pvs = ((
            (p.move(v * t)) % (self.xb, self.yb), v)
            for p, v in self
        )

        return Grid(
            moved_pvs,
            self.xb,
            self.yb
        )

    def iter_quadrants(self, include_bounds=False):
        m_x, m_y = self.xb//2, self.yb//2
        for p, _ in self:
            if p.y == m_y or p.x == m_x:
                if include_bounds:
                    yield -2
                else:
                    continue
            idx_y = 1 if p.y > m_y else 0
            idx_x = 1 if p.x > m_x else 0
            yield (idx_y * 2 + idx_x)

    def points(self):
        return Counter(p for p, _ in self)

    def __str__(self) -> str:
        points = set(self.points().keys())
        return "\n".join(
            "".join(
                "#" if P(x, y) in points else "."
                for x in range(self.xb)
            )
            for y in range(self.yb)
        )

    def seq_counts(self):
        grid_str = str(self)
        return Counter(bitcount(t) for t in zip(grid_str, grid_str[1:], grid_str[2:]))


def bitcount(t):
    v = 0
    for i, c in enumerate(t):
        if c == "#":
            v |= 1 << i
    return v


def part_one(filename, xb, yb):
    matches = load(filename)
    grid = Grid(matches, xb, yb)
    grid = grid.tick(100)

    c = Counter(grid.iter_quadrants())
    safety_factor = reduce(operator.mul, c.values())

    print(f"part one ({filename}): {safety_factor}")


def part_two(filename, xb, yb):
    matches = load(filename)
    grid = Grid(matches, xb, yb)

    for k in range(10000):
        grid = grid.tick(1)
        s = grid.seq_counts()
        if s.get(7, 0) > 100:
            print(f"outlier: {k}->" +
                  " ".join(f"{i}: {s[i]}" for i in range(7, 0, -1)))
            break


if __name__ == "__main__":
    part_one("test1.txt", 11, 7)
    part_one("input.txt", 101, 103)
    part_two("input.txt", 101, 103)
