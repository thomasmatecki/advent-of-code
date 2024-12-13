from math import gcd
import re
from typing import NamedTuple

INPUT_RE = r"Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X\=(\d+), Y\=(\d+)\n"


def lcm(a, b):
    return a * b // gcd(a, b)


class Machine(NamedTuple):
    x_a: int
    y_a: int
    x_b: int
    y_b: int
    px: int
    py: int
    offset: int = 0

    def solve(self):
        """
        x:  x_aA + x_bB = px
        y:  y_aA + y_bB = py
        """
        lcm_a = lcm(self.x_a, self.y_a)

        mx = lcm_a // self.x_a
        my = lcm_a // self.y_a

        py = self.py + self.offset
        px = self.px + self.offset

        b, r_b = divmod((my * py - mx * px),
                        (self.y_b * my - self.x_b * mx))

        a, r_a = divmod(px - self.x_b * b,  self.x_a)

        if r_a != 0 or r_b != 0:
            raise ValueError("No solution")

        return a, b


def load(filename, offset=0):
    with open(f"2024/day13/{filename}") as f:
        matches = re.findall(INPUT_RE, f.read())
    return [Machine(*map(int, match), offset=offset) for match in matches]


def solve(machines):
    for m in machines:
        try:
            yield m.solve()
        except ValueError:
            pass


def part_one(filename):
    tokens = sum(3 * a + b for a, b in solve(load(filename)))
    print(f"part one ({filename}): {tokens}")


def part_two(filename):
    machines = load(filename, offset=10000000000000)
    tokens = sum(3 * a + b for a, b in solve(machines))
    print(f"part two ({filename}): {tokens}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")

    part_two("test1.txt")
    part_two("input.txt")
