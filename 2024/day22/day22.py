
from collections import defaultdict
from functools import reduce
from itertools import islice, pairwise
from operator import xor
import secrets


mix = xor


def prune(n0):
    return n0 % 16777216


def next_secret(number):
    m_64 = number * 64
    number = mix(number, m_64)
    number = prune(number)

    d_32 = number // 32
    number = mix(number, d_32)
    number = prune(number)

    m_2048 = number * 2048
    number = mix(number, m_2048)
    number = prune(number)

    return number


def part_one(filename):
    s = 0
    with open(f"2024/day22/{filename}") as f:
        for line in f.readlines():
            number = int(line.strip())
            for _ in range(2000):
                number = next_secret(number)
            s += number
    print(f"part one ({filename}): {s}")


def secret_seq(start, count):
    secret = start
    for _ in range(count):
        yield secret
        secret = next_secret(secret)


def part_two(filename):
    counts = defaultdict(int)
    with open(f"2024/day22/{filename}") as f:
        for line in f.readlines():
            ks = set()
            number = int(line.strip())
            secrets = secret_seq(number, 2000)
            prices = (a % 10 for a in secrets)
            p_diffs = ((b, b-a) for a, b in pairwise(prices))
            ps, seq_k = zip(*tuple(islice(p_diffs, 4)))

            counts[seq_k] += ps[3]

            for p, d in p_diffs:
                seq_k = seq_k[1:] + (d,)
                if seq_k in ks:
                    continue
                counts[seq_k] += p
                ks.add(seq_k)

    m = max((b, a) for a, b in counts.items())

    print(f"part two ({filename}): {m}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")

    part_two("test3.txt")
    part_two("input.txt")
