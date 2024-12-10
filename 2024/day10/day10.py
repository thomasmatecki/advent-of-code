from collections import UserList


class Grid(UserList):

    @classmethod
    def _int(cls, v):
        try:
            return int(v)
        except ValueError:
            return -1

    @classmethod
    def from_file(cls, filename: str) -> "Grid":
        with open(f"2024/day10/{filename}") as f:
            return cls(
                map(list, (map(cls._int, line)
                    for line in map(str.strip, f.readlines())))
            )

    def trailheads(self):
        for i, row in enumerate(self):
            for j, cell in enumerate(row):
                if cell == 0:
                    yield i, j

    def neighbors(self, i, j):
        for i0, j0 in [(i-1, j), (i+1, j), (i, j-1), (i, j+1)]:
            if 0 <= i0 < len(self) and 0 <= j0 < len(self[0]):
                yield i0, j0

    def value(self, i, j):
        return self[i][j]

    def traverse(self, collect, *p):
        paths = [p]
        while paths:
            path = paths.pop()
            v = self.value(*path[-1])
            for neighbor in self.neighbors(*path[-1]):
                u = self.value(*neighbor)
                if u != v + 1:
                    continue
                if u == 9:
                    collect(neighbor)
                else:
                    new_path = path + (neighbor,)
                    paths.append(new_path)

    def score(self, trailhead):
        trailends = set()
        self.traverse(trailends.add, trailhead)
        return len(trailends)

    def rating(self, trailhead):
        trailends = list()
        self.traverse(trailends.append, trailhead)
        return len(trailends)


def part_one(filename):
    grid = Grid.from_file(filename)
    total_score = sum(grid.score(trailhead)for trailhead in grid.trailheads())
    print(f"part one ({filename}): {total_score}")


def part_two(filename):
    grid = Grid.from_file(filename)
    total_score = sum(grid.rating(trailhead)for trailhead in grid.trailheads())
    print(f"part two ({filename}): {total_score}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("test3.txt")
    part_one("test4.txt")
    part_one("input.txt")
    part_two("input.txt")
