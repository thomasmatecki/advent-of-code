const std = @import("std");

fn calcFuel(i: i32) i32 {
    return @divFloor(i, 3) - 2;
}

const Iterator = struct {
    file: std.fs.File,
    file_reader: std.fs.File.Reader,
    buffered_reader: std.io.BufferedReader(4096, std.fs.File.Reader),
    line_reader: std.io.AnyReader,
    buf: [32]u8 = undefined,
    fn next(self: *Iterator) !?i32 {
        if (try self.line_reader.readUntilDelimiterOrEof(&self.buf, '\n')) |line| {
            return try std.fmt.parseInt(i32, line, 10);
        } else {
            return null;
        }
    }
    fn close(self: *Iterator) void {
        self.file.close();
    }
};

fn inputIterator(path: []const u8) !Iterator {
    var file = try std.fs.cwd().openFile(path, .{});
    const file_reader = file.reader();
    var buffered_reader = std.io.bufferedReader(file_reader);

    return Iterator{
        .file = file,
        .file_reader = file_reader,
        .buffered_reader = buffered_reader,
        .line_reader = buffered_reader.reader().any(),
    };
}

pub fn part1() !void {
    var iterator = try inputIterator("day1/input.txt");
    defer iterator.close();

    var s: i32 = 0;
    while (try iterator.next()) |i| {
        s += calcFuel(i);
    }

    try std.io.getStdOut().writer().print("Day 1, Part One: {d}\n", .{s});
}

pub fn part2() !void {
    var iterator = try inputIterator("day1/input.txt");
    defer iterator.close();
    var s: i32 = 0;

    while (try iterator.next()) |i| {
        var j = calcFuel(i);
        while (j > 0) {
            s += j;
            j = calcFuel(j);
        }
    }

    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 1, Part Two: {d}\n", .{s});
}

pub fn main() !void {
    try part1();
    try part2();
}
