from collections import Counter
from itertools import count, tee
from operator import is_


def load_reports():
    with open("2024/day2/input.txt", "r") as file:
        for line in file.readlines():
            yield map(int, line.split(" "))


def is_safe(report):
    report_iter = iter(report)

    r0, r1 = tee(report_iter, 2)
    next(r1)
    ds = list(v1 - v0 for v0, v1 in zip(r0, r1))

    if all(0 < d <= 3 for d in map(abs, ds)) and len(set(d > 0 for d in ds)) == 1:
        return True

    return False


def part_one():
    safe_count = 0
    for report_iter in load_reports():
        if is_safe(report_iter):
            safe_count += 1

    print(f"part one: {safe_count}")


def part_two():
    safe_count = 0
    for report_iter in load_reports():
        report = list(report_iter)
        if is_safe(report):
            safe_count += 1
            continue
        for i in range(len(report)):
            if is_safe(report[:i] + report[i + 1 :]):
                safe_count += 1
                break

    print(f"part two: {safe_count}")


if __name__ == "__main__":
    part_one()
    part_two()
