from collections import UserList
from typing import NamedTuple


class Map(UserList):
    def __init__(self, data):
        super().__init__(data)
        for i, row in enumerate(self):
            for j, cell in enumerate(row):
                if cell == "@":
                    self.robot = (i, j)
                    return

    def handle_dblock(self, i, j, di, dj, check=False):
        match di, dj:
            case 0, _:  # horizontal
                _j1, _j2 = j + dj, j + 2 * dj
                ok = self.move(i, _j2, di, dj, check=check)

                if not check and ok:
                    self[i][_j2] = self[i][_j1]
                    self[i][_j1] = self[i][j]
                    self[i][j] = "."
                return ok

            case _, 0:  # vertical
                v = self[i + di][j]
                j_ = j + 1 if v == "[" else j - 1

                ok = check or self.move(i + di, j_, di, dj, check=True)
                ok = ok and self.move(i + di, j, di, dj, check=check)
                ok = ok and self.move(i + di, j_, di, dj, check=check)

                if not check and ok:
                    self[i + di][j] = self[i][j]
                    self[i][j] = "."
                return ok
            case _, _:
                raise ValueError("Invalid value")

    def move(self, i, j, di, dj, check=False):
        _i, _j = i + di, j + dj
        match self[_i][_j]:
            case ".":
                if not check:
                    self[_i][_j] = self[i][j]
                    self[i][j] = "."
                return True
            case "[" | "]":
                return self.handle_dblock(i, j, di, dj, check=check)
            case "#":
                return False
            case "O":
                ok = self.move(_i, _j, di, dj, check=check)
                ok = ok and self.move(i, j, di, dj, check=check)
                return ok
            case _:
                raise ValueError("Invalid value")

    def move_robot(self, i, j):
        ok = self.move(*self.robot, i, j)
        if ok:
            self.robot = (self.robot[0] + i, self.robot[1] + j)
        return ok

    def __str__(self):
        return "\n".join("".join(row) for row in self.data)

    def gps_coordinates(self):
        for i, row in enumerate(self):
            for j, cell in enumerate(row):
                if cell in ("O", "["):
                    yield 100 * i + j

    def run(self, moves, verbose=False):
        if verbose:
            print(self)

        for move in moves:
            success = self.move_robot(*move)
            if verbose:
                print(f"\nMove {move}: " + ("success" if success else "failed"))
                print(self)

        return sum(self.gps_coordinates())


class Move(NamedTuple):
    i: int
    j: int

    @classmethod
    def from_str(cls, s):
        match s:
            case "^":
                return cls(-1, 0)
            case "v":
                return cls(1, 0)
            case "<":
                return cls(0, -1)
            case ">":
                return cls(0, 1)
            case _:
                raise ValueError("Invalid move")

    def __str__(self):
        match self:
            case Move(-1, 0):
                return "^"
            case Move(1, 0):
                return "v"
            case Move(0, -1):
                return "<"
            case Move(0, 1):
                return ">"
            case _:
                raise ValueError("Invalid move")


def load(filename, subst={}):
    with open(f"2024/day15/{filename}") as f:
        grid, movement_cs = f.read().split("\n\n")

    moves = [Move.from_str(m) for m in "".join(movement_cs.split())]
    warehouse = Map(map(list, grid.splitlines()))

    return warehouse, moves


def part_one(filename):

    warehouse, moves = load(filename)

    warehouse.run(moves)
    gps_sum = sum(warehouse.gps_coordinates())

    print(f"part one ({filename}): {gps_sum}")


def part_two(filename):

    warehouse, moves = load(filename, {"#": "##", "O": "[]", ".": "..", "@": "@."})

    warehouse.run(moves)
    gps_sum = sum(warehouse.gps_coordinates())

    print(f"part two ({filename}): {gps_sum}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("input.txt")

    part_two("test2.txt")
    part_two("input.txt")
