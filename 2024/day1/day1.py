from collections import Counter


def load_lists():
    with open("2024/day1/input.txt", "r") as file:
        return (
            (map(int, r))
            for r in zip(*[row.split("   ", 2) for row in file.read().split("\n")])
        )


def part_one():
    r0, r1 = map(sorted, load_lists())
    print(
        f"part one: {sum(max(v0, v1) - min(v0, v1) for v0, v1 in zip(r0, r1))}")


def part_two():
    r0, r1 = load_lists()
    r1 = Counter((map(int, r1)))
    print(f"part two: {sum(v * r1[v] for v in r0)}")


if __name__ == "__main__":
    part_one()
    part_two()
