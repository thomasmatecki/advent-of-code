from collections import UserList, defaultdict, deque
from errno import EEXIST
from itertools import zip_longest
from typing import NamedTuple


class C(NamedTuple):
    file_id: int | None
    start: int
    end: int

    @property
    def length(self) -> int:
        return self.end - self.start

    def __repr__(self) -> str:
        return f"C(id={self.file_id}, length={self.length}, start={self.start}, end={self.end})"


def part_one(filename):
    with open(f"2024/day9/{filename}") as f:
        disk_map = f.read().strip()
        chunks = list()
        cursor = 0

        for file_id, (file_size, free_space) in enumerate(
            zip_longest(disk_map[::2], disk_map[1::2], fillvalue=0)
        ):
            file_end = cursor + int(file_size)
            chunks.append(C(file_id, cursor, file_end))
            cursor = file_end
            free_length = int(free_space)
            if free_length > 0:
                free_end = cursor + free_length
                chunks.append(C(None, cursor, free_end))
                cursor = free_end

        r_idx = len(chunks) - 1

        for i, chunk in enumerate(chunks):
            if i >= r_idx:
                break

            if chunk.file_id is None:
                empty = chunk
                while chunks[r_idx].file_id is None:
                    r_idx -= 1

                move = chunks.pop(r_idx)
                free_diff = empty.length - move.length
                if free_diff > 0:
                    chunks[i] = C(move.file_id, empty.start, empty.start + move.length)
                    chunks.insert(i + 1, C(None, empty.start + move.length, empty.end))
                elif free_diff < 0:
                    chunks[i] = C(move.file_id, empty.start, empty.end)
                    chunks.insert(
                        r_idx, C(move.file_id, move.start, move.start - free_diff)
                    )
                else:
                    chunks[i] = C(move.file_id, chunk.start, chunk.start + move.length)
                    r_idx -= 1

        checksum = 0
        for chunk in chunks:
            if chunk.file_id is None:
                continue
            for i in range(chunk.start, chunk.end):
                checksum += i * chunk.file_id

        print(f"part one ({filename}): {checksum}")


def part_two(filename):

    with open(f"2024/day9/{filename}") as f:
        disk_map = f.read().strip()
        chunks = list()
        free_spaces = list()
        offsets = defaultdict(int)
        cursor = 0

        for file_id, (file_size, free_space) in enumerate(
            zip_longest(disk_map[::2], disk_map[1::2], fillvalue=0)
        ):
            file_end = cursor + int(file_size)
            chunks.append(C(file_id, cursor, file_end))
            cursor = file_end
            free_length = int(free_space)
            if free_length > 0:
                free_end = cursor + free_length
                free_spaces.append(C(None, cursor, free_end))
                cursor = free_end

        for i, chunk in reversed(list(enumerate(chunks))):
            start = offsets[chunk.length]
            for j_, free_space in enumerate(free_spaces[start:]):
                j = j_ + start
                if chunk.start < free_space.start:
                    break

                if free_space.length < chunk.length:
                    # Keep track of the lowest index where chunk.length won't
                    # fit, start the nexto chunk.length search from there
                    offsets[chunk.length] = j
                    continue

                if free_space.length > chunk.length:
                    chunks[i] = C(
                        chunk.file_id, free_space.start, free_space.start + chunk.length
                    )
                    free_spaces[j] = C(
                        None, free_space.start + chunk.length, free_space.end
                    )

                else:
                    chunks[i] = C(
                        chunk.file_id, free_space.start, free_space.start + chunk.length
                    )
                    free_spaces.pop(j)

                break

        checksum = 0
        for chunk in chunks:
            if chunk.file_id is None:
                continue
            for i in range(chunk.start, chunk.end):
                checksum += i * chunk.file_id

        print(f"part two ({filename}): {checksum}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("test2.txt")
    part_one("input.txt")
    part_two("test2.txt")
    part_two("input.txt")
