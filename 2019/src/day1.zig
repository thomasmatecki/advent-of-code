const std = @import("std");
const aoc: type = @import("aoc");

fn calcFuel(i: i32) i32 {
    return @divFloor(i, 3) - 2;
}

pub fn part1() !void {
    var iterator = try aoc.inputIterator(i32, "day1/input.txt", '\n');
    defer iterator.close();

    var s: i32 = 0;
    while (try iterator.next()) |i| {
        s += calcFuel(i);
    }

    try std.io.getStdOut().writer().print("Day 1, Part One: {d}\n", .{s});
}

pub fn part2() !void {
    var iterator = try aoc.inputIterator(i32, "day1/input.txt", '\n');
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
