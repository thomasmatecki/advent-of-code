from collections import deque
from operator import add, mul


def concat(a, b):
    return int(str(a) + str(b))


def iter_eqs(eqs, ops):
    for result, operands in eqs:
        q = deque([operands])
        while q:
            operands = q.pop()
            if len(operands) > 1:
                for oper in ops:
                    o0 = operands[0]
                    o1 = operands[1]
                    q.append([oper(o0, o1)] + operands[2:])
            elif operands[0] == result:
                yield result
                break


def load_eqs(filename):
    with open(f"2024/day7/{filename}") as f:
        lines = f.readlines()
        lines = [line.split(":") for line in lines]
        eqs = [(int(line[0]), list(map(int, line[1].split()))) for line in lines]
    return eqs


def part_one(filename):
    total_result = sum(iter_eqs(load_eqs(filename), (add, mul)))
    print(f"part one ({filename}): {total_result}")


def part_two(filename):
    eqs = load_eqs(filename)
    total_result = sum(iter_eqs(eqs, (add, mul, concat)))

    print(f"part two ({filename}): {total_result}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
