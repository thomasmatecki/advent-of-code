const std = @import("std");
const aoc: type = @import("aoc");

const stdout = std.io.getStdOut().writer();

fn runIntcodes(noun: u32, verb: u32, xs: std.ArrayList(u32)) !u32 {
    var i: usize = 0;

    xs.items[1] = noun;
    xs.items[2] = verb;

    while (xs.items[i] != 99) : ({
        i += 4;
    }) {
        const a = &xs.items[i + 1];
        const b = &xs.items[i + 2];
        const c = &xs.items[i + 3];

        xs.items[c.*] = switch (xs.items[i]) {
            1 => xs.items[a.*] + xs.items[b.*],
            2 => xs.items[a.*] * xs.items[b.*],
            else => {
                return error.InvalidOp;
            },
        };
    }
    return xs.items[0];
}

pub fn part1() !void {
    var iterator = try aoc.inputIterator(u32, "day2/input.txt", ',');
    defer iterator.close();

    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const xs = try iterator.collect(arena.allocator());
    defer xs.deinit();

    const result = try runIntcodes(12, 2, @as(std.ArrayList(u32), xs));

    try stdout.print("Day 2, Part One: {d}\n", .{result});
}

pub fn part2() !void {
    // Brute force search over all possible noun/verb combinations
    var iterator = try aoc.inputIterator(u32, "day2/input.txt", ',');
    defer iterator.close();

    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    const allocator = arena.allocator();
    defer arena.deinit();

    const xs = try iterator.collect(allocator);
    defer xs.deinit();

    for (0..100) |verb| {
        for (0..100) |noun| {
            const ys = try xs.clone();
            const result = try runIntcodes(
                @intCast(noun),
                @intCast(verb),
                ys,
            );
            if (result == 19690720) {
                ys.deinit();
                try stdout.print("Day 2, Part Two: {d}\n", .{100 * noun + verb});
                break;
            }
            ys.deinit();
        }
    }
}

pub fn part2_incomplete() !void {
    // TODO: this is a better way, but not done.
    var iterator = try aoc.inputIterator(u32, "day2/input.txt", ',');
    defer iterator.close();

    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    const allocator = arena.allocator();
    defer arena.deinit();

    const xs = try iterator.collect(allocator);
    defer xs.deinit();

    var l: u32 = 0;
    var r: u32 = 99;
    var noun: u32 = 0;
    var verb: u32 = 0;

    while (l != r) {
        noun = (l + r) / 2;
        const ys = try xs.clone();
        const output = try runIntcodes(noun, 0, ys);
        if (output < 19690720) {
            l = noun + 1;
        } else {
            r = noun;
        }
    }

    try stdout.print("Day 1, Part Noun: {d}\n", .{noun});

    l = 0;
    r = 99;
    while (l != r) {
        verb = (l + r) / 2;
        const ys = try xs.clone();
        const output = try runIntcodes(0, verb, ys);
        if (output < 19690720) {
            l = verb + 1;
        } else {
            r = verb;
        }
    }

    try stdout.print("Day 2, Part Verb: {d}\n", .{verb});
}

pub fn main() !void {
    try part1();
    try part2();
}
