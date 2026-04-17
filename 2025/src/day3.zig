const std = @import("std");
const stdout = std.io.getStdOut().writer();

const Max = struct { index: usize, value: u8 };

fn part1() !void {
    var file = try std.fs.cwd().openFile("day3/input.txt", .{});

    var reader = file.reader();
    var buf: [256 * 3]u8 = undefined;
    var total: u64 = 0;

    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        if (line.len == 0) {
            return error.EmptyInput;
        }

        var lmax: Max = .{
            .index = 0,
            .value = 0,
        };

        for (line[0 .. line.len - 1], 0..) |value, i| {
            if (value > lmax.value) {
                lmax = .{
                    .index = i,
                    .value = value,
                };
            }
        }

        var rmax: Max = .{
            .index = 0,
            .value = 0,
        };

        for (line[lmax.index + 1 ..], 0..) |value, i| {
            if (value > rmax.value) {
                rmax = .{
                    .index = i,
                    .value = value,
                };
            }
        }
        var num_buffer: [2]u8 = undefined;

        const num_slice = try std.fmt.bufPrint(&num_buffer, "{d}{d}", .{
            try std.fmt.parseInt(u8, &[_]u8{lmax.value}, 10),
            try std.fmt.parseInt(u8, &[_]u8{rmax.value}, 10),
        });

        total += try std.fmt.parseInt(u64, num_slice, 10);
    }

    try stdout.print("Day 3, Part One: {d}\n", .{total});
}

fn part2() !void {
    var file = try std.fs.cwd().openFile("day3/input.txt", .{});

    var reader = file.reader();
    var buf: [256 * 3]u8 = undefined;
    var total: u64 = 0;

    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var max_values: [12]Max = @splat(Max{
            .index = 0,
            .value = 0,
        });

        var fr: usize = 0;

        for (&max_values, 0..12) |*mv, i| {
            const to = line.len - (12 - i) + 1;

            for (line[fr..to], fr..) |value, j| {
                if (value > mv.value) {
                    mv.* = .{
                        .index = j,
                        .value = value,
                    };
                }
            }
            fr = mv.index + 1;
        }

        var num_buffer: [12]u8 = undefined;

        for (&max_values, 0..12) |*mv, i| {
            num_buffer[i] = mv.value;
        }

        const n = try std.fmt.parseInt(u64, &num_buffer, 10);
        total += n;
    }
    try stdout.print("Day 3, Part Two: {d}\n", .{total});
}

pub fn main() !void {
    try part1();
    try part2();
}
