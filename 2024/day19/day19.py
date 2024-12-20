from functools import cache
from itertools import groupby


def load(filename):
    with open(f"2024/day19/{filename}") as f:
        patterns, _, *designs = f.read().splitlines()
        patterns = [s.strip() for s in patterns.split(",")]

    by_len = {l: set([*ps])
              for l, ps in groupby(sorted(patterns, key=len), len)}
    return by_len, designs


def part_one(filename):
    by_len, designs = load(filename)

    def check(design):
        if design == "":
            return True
        return any(design[:l] in ps and check(design[l:]) for l, ps in by_len.items())

    print(sum(check(design) for design in designs))


def part_two(filename):

    by_len, designs = load(filename)

    @cache
    def check(design):
        if design == "":
            return 1

        return sum(check(design[l:]) for l, ps in by_len.items() if design[:l] in ps)

    print(sum(check(design) for design in designs))


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
