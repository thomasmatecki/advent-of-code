from collections import UserList
from enum import Enum
import heapq
from typing import NamedTuple


from functools import wraps
import time


def timeit(func):
    @wraps(func)
    def timeit_wrapper(*args, **kwargs):
        start_time = time.perf_counter()
        result = func(*args, **kwargs)
        end_time = time.perf_counter()
        total_time = end_time - start_time
        print(
            f"Function {func.__name__}{args} {kwargs if kwargs else ''} Took {total_time: .4f} seconds"
        )
        return result

    return timeit_wrapper


class Dir(Enum):
    N = 0
    E = 1
    S = 2
    W = 3

    def __gt__(self, other):
        return self.value > other.value


class Node(NamedTuple):
    i: int
    j: int
    d: Dir

    def cost(self, other):
        return (
            abs(self.i - other.i)
            + abs(self.j - other.j)
            + abs((self.d.value - other.d.value) * 1000)
        )

    @property
    def ij(self):
        return self.i, self.j


class Minimum(NamedTuple):
    cost: int
    prevs: list[Node]


class Maze(UserList):

    @classmethod
    def from_file(cls, filename):
        with open(f"2024/day16/{filename}") as f:
            data = f.read().splitlines()
        return cls(data)

    def start(self):
        for i, row in enumerate(self):
            if "S" in row:
                start = Node(i, row.index("S"), Dir.E)
                break
        return start

    def ends(self):
        for i, row in enumerate(self):
            if "E" in row:
                j = row.index("E")
                for d in Dir:
                    yield Node(i, j, d)

    def in_bounds(self, node):
        i, j, _ = node
        return 0 <= i < len(self) and 0 <= j < len(self[0]) and self[i][j] != "#"

    def neigbors(self, n):
        i, j, d = n

        match d:
            case Dir.N:
                ms = (
                    (1, Node(i - 1, j, Dir.N)),
                    (1000, Node(i, j, Dir.E)),
                    (1000, Node(i, j, Dir.W)),
                )
            case Dir.E:
                ms = (
                    (1, Node(i, j + 1, Dir.E)),
                    (1000, Node(i, j, Dir.N)),
                    (1000, Node(i, j, Dir.S)),
                )
            case Dir.S:
                ms = (
                    (1, Node(i + 1, j, Dir.S)),
                    (1000, Node(i, j, Dir.E)),
                    (1000, Node(i, j, Dir.W)),
                )
            case Dir.W:
                ms = (
                    (1, Node(i, j - 1, Dir.W)),
                    (1000, Node(i, j, Dir.N)),
                    (1000, Node(i, j, Dir.S)),
                )

        for c, m in ms:
            if self.in_bounds(m):
                yield c, m

    def _graph(self):
        g = {}
        for i, row in enumerate(self):
            for j, c in enumerate(row):
                if c != "#":
                    for d in Dir:
                        n = Node(i, j, d)
                        g[n] = {n_: c for c, n_ in self.neigbors(n)}
        return g

    def djikstra(self, start):
        g = self._graph()

        mins = {}
        q = []
        heapq.heapify(q)

        mins[start] = Minimum(0, [])
        heapq.heappush(q, (0, start))
        while q:
            c, v = heapq.heappop(q)

            if mins[v].cost < c:
                continue

            for n, c_ in g[v].items():
                _c = c + c_
                if n not in mins or mins[n].cost > _c:
                    mins[n] = Minimum(_c, [v])
                    heapq.heappush(q, (_c, n))
                elif mins[n].cost == _c:
                    mins[n].prevs.append(v)

        return mins


@timeit
def part_one(filename):
    maze = Maze.from_file(filename)
    mins = maze.djikstra(maze.start())

    score = min(mins[end].cost for end in maze.ends())

    print(f"part one: {score}")


def backtrace(mins: dict[Node, Minimum], ends: list[Minimum]):
    q = [(p,) for e in ends for p in e.prevs]
    tiles = set()
    while q:
        path = q.pop()
        last = path[-1]
        tiles.add(last.ij)
        _, prevs = mins[last]
        for p in prevs:
            p0 = path + (p,)
            q.append(p0)  # type: ignore

    return len(tiles) + 1


@timeit
def part_two(filename):
    maze = Maze.from_file(filename)
    mins = maze.djikstra(maze.start())

    end_mins = [mins[end] for end in maze.ends()]
    min_cost = min(end_mins).cost
    mins_ends = [e for e in end_mins if e.cost == min_cost]

    tiles = backtrace(mins, mins_ends)

    print(f"part two: {tiles}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("input.txt")

    part_two("test1.txt")
    part_two("test2.txt")
    part_two("input.txt")
