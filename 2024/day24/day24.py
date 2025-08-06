from collections import defaultdict
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
        r_idx = defaultdict(list)
        initials_str, gates_str = f.read().split("\n\n")

        for name, initial in re.findall(INITIALS_RE, initials_str):
            idx[name] = Initial(bool(int(initial)))

        for cls in [OR, AND, XOR]:
            gate_re = GATE_RE_TEMPLATE.format(cls.__name__)
            for l, r, name in re.findall(gate_re, gates_str):
                idx[name] = cls(idx, l, r)
                r_idx[l].append(name)
                r_idx[r].append(name)

        return idx, r_idx


def part_one(filename):
    idx, _ = load(filename)
    o = number_for_prefix(idx, "z")
    print(f"Part one ({filename}): {o}")


def swap(a, b, idx):
    t = idx[a]
    idx[a] = idx[b]
    idx[b] = t


def factors(idx, gate) -> set[str]:
    q = [gate]
    f = set(q)

    while q:
        g = q.pop()
        match idx[g]:
            case Initial(_):
                pass
            case XOR(_, l, r) | OR(_, l, r) | AND(_, l, r):
                n = (l, r)
                q.extend(n)
                f.update(n)

    return f


def multiples(r_idx, gate) -> set[str]:
    q = [gate]
    m = set()
    while q:
        g = q.pop()
        n = [ng for ng in r_idx[g] if r_idx[ng]]
        q.extend(n)
        m.update(n)
    return m


def part_two(filename):
    idx, r_idx = load(filename)

    swaps = [
        ("z09", "gwh"),
        # ("wgb", "wbw"),
        # ("z21", "rcb"),
        # ("jct", "z39")
    ]

    for a, b in swaps:
        swap(a, b, idx)

    correct = set()
    correct.difference()

    n_x = number_for_prefix(idx, "x")
    n_y = number_for_prefix(idx, "y")
    n_z = number_for_prefix(idx, "z")
    fs = dict()
    ds = dict()
    for i, c in enumerate(reversed(bin(n_z ^ n_x + n_y))):
        f = factors(idx, f"z{i:02}")
        if c == "1":
            fs[i] = f
            ds[i] = f.difference(correct)
            n0 = multiples(r_idx, f"x{i:02}") | multiples(r_idx, f"y{i:02}")
        else:
            correct.update(f)

    pass

#    print(f"part two: {','.join(sorted(k for t in swaps for k in t))}")


if __name__ == "__main__":
    #    part_one("test1.txt")
    #    part_one("test2.txt")
    #    part_one("input.txt")

    part_two("input.txt")
