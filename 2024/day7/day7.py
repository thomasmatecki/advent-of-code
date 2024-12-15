from collections import deque
from itertools import product, zip_longest
from operator import add, mul

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


@timeit
def part_one(filename):
    total_result = sum(iter_eqs(load_eqs(filename), (add, mul)))
    print(f"part one ({filename}): {total_result}")


@timeit
def part_two(filename):
    eqs = load_eqs(filename)
    total_result = sum(iter_eqs(eqs, (add, mul, concat)))

    print(f"part two ({filename}): {total_result}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
