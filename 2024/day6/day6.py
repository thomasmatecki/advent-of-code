from collections import UserList


def turn_right(i_step, j_step):
    match i_step, j_step:
        case -1, 0:
            return 0, 1
        case 0, 1:
            return 1, 0
        case 1, 0:
            return 0, -1
        case 0, -1:
            return -1, 0
    raise ValueError("Invalid step")


class Grid(UserList[list]):
    @classmethod
    def from_file(cls, filename):
        with open(f"2024/day6/{filename}") as f:
            return cls(map(list, [line.strip() for line in f.readlines()]))

    def in_bounds(self, i, j):
        return 0 <= i < len(self) and 0 <= j < len(self[0])

    def find_start(self):
        for i, row in enumerate(self):
            for j, val in enumerate(row):
                if val == "^":
                    return i, j
        raise ValueError("No start found")

    def steps_from(self, i_step, j_step, i, j) -> tuple[bool, list]:
        steps = set()
        path = []
        while self.in_bounds(i, j):
            next_i = i + i_step
            next_j = j + j_step
            try:
                while self[next_i][next_j] in "#O":
                    i_step, j_step = turn_right(i_step, j_step)
                    next_i = i + i_step
                    next_j = j + j_step
            except IndexError:
                break
            i = next_i
            j = next_j

            step = (i, j, i_step, j_step)
            if step in steps:
                return True, path
            steps.add(step)
            path.append(step)

        return False, path

    def updated(self, i, j, v):
        i_row = self[i].copy()
        i_row[j] = v
        grid = Grid(self[:i] + [i_row] + self[i + 1 :])
        return grid

    def __str__(self) -> str:
        return "\n".join("".join(row) for row in self)


def part_one(filename):
    grid = Grid.from_file(filename)
    i_start, j_start = grid.find_start()
    cycle, path = grid.steps_from(-1, 0, i_start, j_start)
    assert not cycle
    positions = set((i, j) for i, j, _, _ in path)
    print(f"part one ({filename}): {len(positions)}")


def part_two(filename):
    grid = Grid.from_file(filename)
    cycles = 0
    i_start, j_start = grid.find_start()
    cycle, path = grid.steps_from(-1, 0, i_start, j_start)

    for i, j in set((i, j) for i, j, _, _ in path):
        cycle, _ = grid.updated(i, j, "O").steps_from(-1, 0, i_start, j_start)
        if cycle:
            cycles += 1

    print(f"part two({filename}): {cycles}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
