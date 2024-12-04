from collections import UserList


def mapm(iterable, *funcs):
    for f in funcs:
        iterable = map(f, iterable)
    return iterable


def vrev(v):
    return mapm(v, reversed, str().join)


def in_bounds(rows, i, j):
    return i >= 0 and j >= 0 and i < len(rows) and j < len(rows[0])


def diags(rows, k_start, i_step, j_start, j_step):
    diag_strs = []
    for k in range(k_start, len(rows)):
        j = j_start
        i = k
        diag = []
        while in_bounds(rows, i, j):
            diag.append(rows[i][j])
            i += i_step
            j += j_step
        diag_strs.append(str().join(diag))

    return diag_strs


class SearchStrings(UserList[str]):

    def include(self, v):
        self.extend(v)
        self.extend(vrev(v))

    def count(self, string):
        return sum(s.count(string) for s in self)


def part_one(filename):

    search_strs = SearchStrings()

    with open(f"2024/day4/{filename}", "r") as file:
        rows = list(map(str.strip, file.readlines()))
        search_strs.include(rows)

        cols = list(map(str().join, zip(*rows)))
        search_strs.include(cols)

        ur_diags = diags(rows, 0, -1, 0, 1)
        search_strs.include(ur_diags)

        dl_diags = diags(rows, 1, 1, len(rows[0]) - 1, -1)
        search_strs.include(dl_diags)

        ul_diags = diags(rows, 0, -1, len(rows[0]) - 1, -1)
        search_strs.include(ul_diags)

        dr_diags = diags(rows, 1, 1, 0, 1)
        search_strs.include(dr_diags)

        print(f"part_one ({filename}): {search_strs.count('XMAS')}")


# Possible X-MAS matches
X_MATCHES = {"MSMS", "MMSS", "SSMM", "SMSM"}


def is_x_mas(rows, i, j):
    try:
        corners = [
            rows[i - 1][j - 1],
            rows[i - 1][j + 1],
            rows[i + 1][j - 1],
            rows[i + 1][j + 1],
        ]
        vs = str().join(corners)
        return vs in X_MATCHES
    except IndexError:
        pass
    return False


def part_two(filename):
    count = 0
    with open(f"2024/day4/{filename}", "r") as file:
        rows = list(mapm(file.readlines(), str.strip, list))

        for i in range(1, len(rows) - 1):
            for j in range(1, len(rows[0]) - 1):
                if rows[i][j] == "A":
                    if is_x_mas(rows, i, j):
                        count += 1

    print(f"part_two ({filename}): {count}")


if __name__ == "__main__":
    part_one("test1.txt")
    part_one("input.txt")
    part_two("test1.txt")
    part_two("input.txt")
