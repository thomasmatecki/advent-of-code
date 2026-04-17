const std = @import("std");

const stdout = std.io.getStdOut().writer();

const Cell = enum { empty, roll };

const DELTAS = [_]i8{ -1, 0, 1 };

fn Grid(comptime width: usize, comptime height: usize) type {
    return struct {
        cells: [width * height]Cell,

        pub fn init() @This() {
            return .{
                .cells = [_]Cell{.empty} ** (width * height),
            };
        }

        pub fn load(path: []const u8) !@This() {
            var grid = init();

            var file = try std.fs.cwd().openFile(path, .{});
            var reader = file.reader();

            var i: usize = 0;

            while (true) {
                const byte = reader.readByte() catch |err| switch (err) {
                    error.EndOfStream => break,
                    else => return err,
                };

                const cell = switch (byte) {
                    '@' => Cell.roll,
                    '.' => Cell.empty,
                    else => continue,
                };

                grid.index(i).* = cell;
                i += 1;
            }

            return grid;
        }

        pub fn set(
            self: *@This(),
            x: usize,
            y: usize,
            val: Cell,
        ) void {
            const idx = y * width + x;
            self.cells[idx] = val;
        }

        pub fn get(
            self: *const @This(),
            x: usize,
            y: usize,
        ) Cell {
            const idx = y * width + x;
            return self.cells[idx];
        }

        pub fn index(self: *@This(), idx: usize) *Cell {
            return &self.cells[idx];
        }

        pub fn neighbors(self: *@This(), x: usize, y: usize) [8]?Cell {
            var ret = [_]?Cell{null} ** 8;
            var idx: usize = 0;

            for (DELTAS) |dx| {
                for (DELTAS) |dy| {
                    if (dx == 0 and dy == 0) continue;

                    const _x: i16 = @intCast(x);
                    const nx: i16 = _x + dx;

                    const _y: i16 = @intCast(y);
                    const ny: i16 = _y + dy;

                    if (nx < 0 or ny < 0 or
                        nx >= @as(i16, width) or ny >= @as(i16, height))
                    {
                        ret[idx] = null;
                        idx += 1;
                        continue;
                    }

                    ret[idx] = self.get(
                        @intCast(nx),
                        @intCast(ny),
                    );

                    idx += 1;
                }
            }

            return ret;
        }

        pub fn canRemove(
            self: *@This(),
            x: usize,
            y: usize,
        ) bool {
            if (self.get(x, y) != Cell.roll) return false;

            const n = self.neighbors(x, y);
            var rcount: u4 = 0;
            for (n) |neighbor| {
                if (neighbor == Cell.roll) {
                    rcount += 1;
                }
            }
            return rcount < 4;
        }
    };
}

pub fn part1(
    input_path: []const u8,
    comptime width: usize,
    comptime height: usize,
) !void {
    var grid = try Grid(width, height).load(input_path);

    var count: u32 = 0;

    for (0..height) |y| {
        for (0..width) |x| {
            if (grid.canRemove(x, y)) {
                count += 1;
            }
        }
    }

    try stdout.print(
        "Day 4, Part One ({s}): {d}\n",
        .{ input_path, count },
    );
}

pub fn part2(
    allocator: std.mem.Allocator,
    input_path: []const u8,
    comptime width: usize,
    comptime height: usize,
) !void {
    var grid = try Grid(width, height).load(input_path);

    var total: u32 = 0;

    while (true) {
        var removed = std.ArrayList(struct { x: usize, y: usize }).init(allocator);
        defer removed.deinit();

        for (0..height) |y| {
            for (0..width) |x| {
                if (grid.canRemove(x, y)) {
                    total += 1;
                    try removed.append(.{
                        .x = x,
                        .y = y,
                    });
                }
            }
        }

        if (removed.items.len == 0) break;

        for (removed.items) |coord| {
            grid.set(coord.x, coord.y, Cell.empty);
        }
    }

    try stdout.print(
        "Day 4, Part One ({s}): {d}\n",
        .{ input_path, total },
    );
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const allocator = gpa.allocator();

    try part1("day4/test1.txt", 10, 10);
    try part1("day4/input.txt", 139, 139);
    try part2(allocator, "day4/test1.txt", 10, 10);
    try part2(allocator, "day4/input.txt", 139, 139);
}
