from collections import UserList, deque
from itertools import groupby
from operator import itemgetter


class Grid(UserList):

    @classmethod
    def from_file(cls, filename):
        with open(f"2024/day12/{filename}") as file:
            return cls(map(str.strip, file.read().splitlines()))

    def neighbors(self, i, j):
        for i0, j0 in [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]:
            yield i0, j0

    def in_bounds(self, i, j):
        return 0 <= i < len(self) and 0 <= j < len(self[0])

    def explore(self, i, j, visited):
        v = self[i][j]
        region = dict()
        q = {(i, j)}

        while q:
            i, j = q.pop()
            region[(i, j)] = 0
            for i0, j0 in self.neighbors(i, j):
                if not self.in_bounds(i0, j0):
                    region[(i, j)] += 1
                elif (i0, j0) in visited:
                    region[(i, j)] += 1
                elif (i0, j0) in region:
                    continue
                elif self[i0][j0] == v:
                    q.add((i0, j0))
                else:
                    region[(i, j)] += 1

        return region

    def regions(self):

        visited = set()
        regions = []

        for i, row in enumerate(self):
            for j, _ in enumerate(row):
                if (i, j) in visited:
                    continue
                region = self.explore(i, j, visited)
                visited.update(region.keys())
                regions.append(region)

        return regions

    def _disjoints(self, xs):
        xs = sorted(xs)
        x = [_x for _x, _ in xs]
        return 1 + sum(1 for n in range(1, len(x)) if x[n] - x[n - 1] > 1)

    def _collate_edges(self, edges, group, get_v):
        return sum(
            self._disjoints(map(get_v, vs))
            for g, vs in groupby(sorted(edges, key=group), group)
            if g[0] != g[1]
        )

    def count_sides(self, region):
        edges = []
        for i0, j0 in region:
            edges.extend(
                ((i0, i1), (j0, j1))
                for i1, j1 in self.neighbors(i0, j0)
                if (i1, j1) not in region
            )

        get_i = itemgetter(0)
        get_j = itemgetter(1)

        return self._collate_edges(edges, get_i, get_j) + self._collate_edges(
            edges, get_j, get_i
        )


def part_one(filename):
    grid = Grid.from_file(filename)
    total_price = sum(len(r) * sum(r.values()) for r in grid.regions())
    print(f"part one ({filename}): {total_price}")


def part_two(filename):

    grid = Grid.from_file(filename)
    total_price = sum(len(r) * grid.count_sides(r) for r in grid.regions())

    print(f"part two ({filename}): {total_price}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("input.txt")

    part_two("test2.txt")
    part_two("test3.txt")
    part_two("input.txt")
