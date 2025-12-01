const std = @import("std");
const aoc: type = @import("aoc");

fn calcFuel(i: i32) i32 {
    return @divFloor(i, 3) - 2;
}

pub fn part1() !void {
    std.debug.print("Day 1, Part One: Not implemented yet\n", .{});
}
