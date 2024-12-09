from collections import defaultdict


def load_input(filename):
    with open(f"2024/day5/{filename}") as f:
        lines = map(str.strip, f.readlines())
        successors = defaultdict(set)
        for line in lines:
            if line == "":
                break
            before, after = line.split("|")
            successors[before].add(after)

        updates = []
        for line in lines:
            updates.append(line.split(","))
    return successors, updates


def wrong_order(update, successors):
    """Find the first pair of elements in the update list that are in the wrong
    order, return the index of the element that should be moved and the index it
    should be moved to."""
    prevs = dict()
    for idx, v in enumerate(update):
        fail = successors[v].intersection(prevs)
        if fail:
            splice = min(prevs[f] for f in fail)
            return idx, splice
        prevs[v] = idx
    return None


def middle(u):
    return int(u[len(u)//2])


def part_one(filename):
    successors, updates = load_input(filename)
    result = sum(middle(u)for u in updates if not wrong_order(u, successors))
    print(f"part one ({filename}): {result}")


def reordered(updates, successors):
    for u in updates:
        reorderings = 0
        while (r := wrong_order(u, successors)) is not None:
            idx, splice = r
            u.insert(splice, u.pop(idx))
            reorderings += 1
        if reorderings > 0:
            yield middle(u)


def part_two(filename):
    successors, updates = load_input(filename)
    result = sum(reordered(updates, successors))
    print(f"part two ({filename}): {result}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
