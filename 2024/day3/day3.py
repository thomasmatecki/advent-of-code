import enum
import re
from dataclasses import dataclass

MULTIPLY_EXPRESSION = r"mul\((\d+),(\d+)\)"


class MatchType(enum.Enum):
    MULTIPLY = MULTIPLY_EXPRESSION
    DO = r"do\(\)"
    DONT = r"don't\(\)"


@dataclass
class Match:
    match_type: MatchType
    re_match: re.Match


def part_one():
    with open("2024/day3/input.txt", "r") as file:
        input = file.read()
        r = sum(int(x) * int(y) for x, y in re.findall(MULTIPLY_EXPRESSION, input))
        print(f"part one: {r}")


def exec_match(matches):
    multiply_enabled = True
    for m in matches:
        match m:
            case Match(MatchType.MULTIPLY, re_match):
                if multiply_enabled:
                    yield int(re_match.group(1)) * int(re_match.group(2))
            case Match(MatchType.DO, _):
                multiply_enabled = True
            case Match(MatchType.DONT, _):
                multiply_enabled = False
            case _:
                raise ValueError("Unknown match type")


def part_two(filename):
    with open(f"2024/day3/{filename}", "r") as file:
        input = file.read()
        matches = sorted(
            (Match(t, m) for t in MatchType for m in re.finditer(t.value, input)),
            key=lambda x: x.re_match.start(),
        )
    total = sum(exec_match(matches))
    print(f"part two ({filename}): {total}")


if __name__ == "__main__":
    part_one()
    part_two("input.txt")
    part_two("test1.txt")
