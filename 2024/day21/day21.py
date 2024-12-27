
from calendar import c
from collections import defaultdict
from email.policy import default
from functools import cache, reduce
from itertools import chain, permutations, product

DIRECTIONAL = (
    (None, "^", "A"),
    ("<", "v", ">"),
)

NUMERIC = (
    ("7", "8", "9"),
    ("4", "5", "6"),
    ("1", "2", "3"),
    (None, "0", "A"),
)


def diff(p1, p2):
    return p2[0] - p1[0], p2[1] - p1[1]


def collate(paths):
    return map(chain.from_iterable, product(*paths))


@cache
def is_path_safe(buttons, path, p):

    for c in path:
        match c:
            case "A":
                pass
            case "^":
                p = p[0] - 1, p[1]
            case "v":
                p = p[0] + 1, p[1]
            case "<":
                p = p[0], p[1] - 1
            case ">":
                p = p[0], p[1] + 1
            case _:
                raise ValueError(f"Invalid character {c}")

        v = buttons[p[0]][p[1]]

        if v is None:
            return False

    return True


@cache
def path_set(buttons, to_p, from_p):
    i, j = diff(from_p, to_p)
    s = []

    if j > 0:
        s.append(">" * j)
    if i > 0:
        s.append("v" * i)
    if i < 0:
        s.append("^" * abs(i))
    if j < 0:
        s.append("<" * abs(j))

    paths = (
        "".join(_s)+"A" for _s in permutations(s, len(s)))

    paths = frozenset(_s for _s in paths if is_path_safe(buttons, _s, from_p))

    return paths


class Keypad:
    def __init__(self, buttons, start) -> None:
        self.p = start
        self.buttons = buttons
        self.coords = {v: (i, j) for i, row in enumerate(buttons)
                       for j, v in enumerate(row)}

    def get_coords(self, *vs):
        return (self.coords[v] for v in vs)

    def do(self, v):
        p = self.coords[v]
        v_paths = path_set(self.buttons, p, self.p)
        self.p = p
        return v_paths

    def backward(self, v):
        return (self.do(c) for c in v)

    def expand_paths(self,  paths, *_):
        return chain.from_iterable(
            collate(self.backward(path)) for path in paths)

    @cache
    def dfs(self, vs: str, depth: int):
        if depth == 0:
            return len(vs)

        plen = 0
        from_v = "A"
        for to_v in vs:
            from_p = self.coords[from_v]
            to_p = self.coords[to_v]
            paths = path_set(self.buttons, to_p, from_p)
            plen += min(self.dfs(path, depth-1) for path in paths)
            from_v = to_v

        return plen


def part_one(filename):
    with open(f"2024/day21/{filename}") as f:
        lines = map(str.strip, f.readlines())

    n = Keypad(NUMERIC, start=(3, 2))
    d = Keypad(DIRECTIONAL, start=(0, 2))

    s = 0
    for line in lines:

        c0 = collate(n.backward(line))
        path_s2 = reduce(d.expand_paths, range(2), c0)

        sum_min_lens = min(map(len, map(list, path_s2)))

        print(line, sum_min_lens)

        s += int(line[:3]) * sum_min_lens

    print(f"part one: ({filename})", s)


def part_two(filename):

    n = Keypad(NUMERIC, start=(3, 2))
    d = Keypad(DIRECTIONAL, start=(0, 2))

    s = 0

    with open(f"2024/day21/{filename}") as f:
        lines = map(str.strip, f.readlines())

    for line in lines:
        l = 0
        from_v = "A"
        for to_v in line:
            l += min(
                d.dfs(path, 25)
                for path
                in path_set(
                    n.buttons,
                    *n.get_coords(from_v, to_v)))

            from_v = to_v

        s += int(line[:3]) * l

    print(f"part two: ({filename})", s)


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
