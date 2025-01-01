
from collections import defaultdict
from email.policy import default
import re


CONNTECTION_RE = r"(\w+)-(\w+)\n"


def part_one(filename):
    with open(f"2024/day23/{filename}") as f:
        matches = re.findall(CONNTECTION_RE, f.read())

    computer_idx = defaultdict(set)

    for a, b in matches:
        computer_idx[a].update([b])
        computer_idx[b].update([a])

    parties = set()

    for k0, s0 in computer_idx.items():
        if k0.startswith("t"):
            for k1 in s0:
                for k2 in computer_idx[k1]:
                    if k2 in s0:
                        parties.add(frozenset([k0, k1, k2]))

    print(f"part one ({filename}): {len(parties)}")


def part_two(filename):
    with open(f"2024/day23/{filename}") as f:
        matches = re.findall(CONNTECTION_RE, f.read())

    N = defaultdict(set)
    V = set()
    parties = set()

    for a, b in matches:
        V.update([a, b])
        N[a].add(b)
        N[b].add(a)

    def bron_kerbosch(R, P, X):
        if not P and not X:
            parties.add(frozenset(R))
        while P:
            v = P.pop()
            bron_kerbosch(
                R | {v},
                P & N[v],
                X & N[v]
            )
            X.add(v)

    bron_kerbosch(set(), V, set())

    _, largest = max((len(p), p) for p in parties)

    password = ",".join(sorted(largest))

    print(f"part two ({filename}): {password}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")

    part_two("test1.txt")
    part_two("input.txt")
