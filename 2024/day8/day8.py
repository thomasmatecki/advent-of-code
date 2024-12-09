
from collections import UserList, defaultdict
from itertools import combinations
from typing import NamedTuple


class T(NamedTuple):
    i: int
    j: int

    def __sub__(self, other):
        return T(self.i - other.i, self.j - other.j)

    def __add__(self, other):
        return T(self.i + other.i, self.j + other.j)

    def __neg__(self):
        return T(-self.i, -self.j)


class Grid(UserList):
    @classmethod
    def from_file(cls, filename):
        with open(f"2024/day8/{filename}") as f:
            lines = map(str.strip, f.readlines())
            return cls(lines)

    def in_bounds(self, v):
        return 0 <= v.i < len(self) and 0 <= v.j < len(self[0])

    def antennas(self):
        antennas = defaultdict(list)
        for i, row in enumerate(self):
            for j, col in enumerate(row):
                if col != '.':
                    antennas[col].append(T(i, j))
        return antennas

    def antinodes(self, gen_points):
        antinodes = set()
        for positions in self.antennas().values():
            for p0, p1 in combinations(positions, 2):
                antinodes.update(gen_points(p0, p1))
        return antinodes


def part_one(filename):
    grid = Grid.from_file(filename)

    def gen_points(p0, p1):
        v = p0 - p1
        for a in [p1 - v, p0 + v]:
            if grid.in_bounds(a):
                yield a

    antinodes = grid.antinodes(gen_points)

    print(f"part one: {len(antinodes)}")


def part_two(filename):
    grid = Grid.from_file(filename)

    def gen_points(p0, p1):
        for v in [p0 - p1, -p0 + p1]:
            a = p0
            while grid.in_bounds(a):
                yield a
                a += v

    antinodes = grid.antinodes(gen_points)

    print(f"part two: {len(antinodes)}")


if __name__ == '__main__':
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
