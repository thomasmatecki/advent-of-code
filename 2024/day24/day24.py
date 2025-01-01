
from collections import defaultdict
from itertools import pairwise
import re
from typing import NamedTuple, OrderedDict


class XOR(NamedTuple):
    idx: dict
    l: str
    r: str

    @property
    def val(self):
        return self.idx[self.l].val ^ self.idx[self.r].val

    @property
    def ln(self):
        return self.idx[self.l]

    @property
    def rn(self):
        return self.idx[self.r]

    def __repr__(self) -> str:
        return f"XOR({self.l}, {self.r})"


class OR(NamedTuple):
    idx: dict
    l: str
    r: str

    @property
    def val(self):
        return self.idx[self.l].val | self.idx[self.r].val

    @property
    def ln(self):
        return self.idx[self.l]

    @property
    def rn(self):
        return self.idx[self.r]

    def __repr__(self) -> str:
        return f"OR({self.l}, {self.r})"


class AND(NamedTuple):
    idx: dict
    l: str
    r: str

    @property
    def val(self):
        return self.idx[self.l].val & self.idx[self.r].val

    @property
    def ln(self):
        return self.idx[self.l]

    @property
    def rn(self):
        return self.idx[self.r]

    def __str__(self) -> str:
        return f"AND({self.l}, {self.r})"


class Initial(NamedTuple):
    val: bool


INITIALS_RE = r"([\d|\w]+): (\d)"
GATE_RE_TEMPLATE = r"([\d|\w]+) {} ([\d|\w]+) -> ([\d|\w]+)\n"


def gates_for_prefix(idx, c):
    return OrderedDict(sorted((k, v) for k, v in idx.items() if k.startswith(c)))


def bits_for_prefix(idx, c):
    ks = gates_for_prefix(idx, c).values()
    return [int(v.val) for v in ks]


def number_for_prefix(idx, c):
    n = 0
    vs = bits_for_prefix(idx, c)
    for i, v in enumerate(vs):
        n |= (v << i)
    return n


def load(filename):

    with open(f"2024/day24/{filename}") as f:
        idx = dict()
        initials_str, gates_str = f.read().split("\n\n")

        for name, initial in re.findall(INITIALS_RE, initials_str):
            idx[name] = Initial(bool(int(initial)))

        for cls in [OR, AND, XOR]:
            gate_re = GATE_RE_TEMPLATE.format(cls.__name__)
            for l, r, name in re.findall(gate_re, gates_str):
                idx[name] = cls(idx, l, r)

        return idx


def part_one(filename):
    idx = load(filename)
    o = number_for_prefix(idx, "z")
    print(f"Part one ({filename}): {o}")


def swap(a, b, idx):
    t = idx[a]
    idx[a] = idx[b]
    idx[b] = t


def part_two(filename):
    idx = load(filename)

    swaps = [("z09", "gwh"),
             ("wgb", "wbw"),
             ("z21", "rcb"),
             ("jct", "z39")]

    for a, b in swaps:
        swap(a, b, idx)

    n_x = number_for_prefix(idx, "x")
    n_y = number_for_prefix(idx, "y")
    n_z = number_for_prefix(idx, "z")
    for i, c in enumerate(reversed(bin(n_z ^ n_x + n_y))):
        if c == "1":
            raise ValueError(f"Bit {i} is wrong")

    print(f"part two: {','.join(sorted(k for t in swaps for k in t))}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("input.txt")

    part_two("input.txt")
