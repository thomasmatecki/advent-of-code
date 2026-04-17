const std = @import("std");
const stdout = std.io.getStdOut().writer();

fn part1() !void {
    var file = try std.fs.cwd().openFile("day2/input.txt", .{});

    var reader = file.reader();
    var buf: [256 * 3]u8 = undefined;
    const line = try reader.readUntilDelimiterOrEof(&buf, '\n') orelse {
        return error.EmptyInput;
    };
    var total: u64 = 0;

    var num_buffer: [32]u8 = undefined;

    var split = std.mem.splitScalar(u8, line, ',');

    while (split.next()) |range| {
        var bounds = std.mem.splitScalar(u8, range, '-');

        if (bounds.next()) |start_str| {
            const start = try std.fmt.parseInt(u64, start_str, 10);

            const lead_str = start_str[0 .. (start_str.len + 1) / 2];
            const lead = try std.fmt.parseInt(u64, lead_str, 10);
            const mult = std.math.pow(u64, 10, (lead_str.len));

            const step = mult + 1;

            var id = if (@mod(start_str.len, 2) == 0)
                lead * step
            else
                (std.math.pow(u64, 10, (start_str.len)) +
                    std.math.pow(u64, 10, (start_str.len / 2)));

            const end_str = bounds.next() orelse return error.InvalidRange;

            const end = try std.fmt.parseInt(u64, end_str, 10);

            while (id <= end) : (id += step) {
                const num_slice = try std.fmt.bufPrint(&num_buffer, "{}", .{id});
                if (@mod(num_slice.len, 2) == 0 and id >= start) {
                    total += id;
                }
            }
        }
    }

    try stdout.print("Day 2, Part One: {d}\n", .{total});
}

fn part2() !void {
    var file = try std.fs.cwd().openFile("day2/input.txt", .{});

    var reader = file.reader();
    var buf: [256 * 3]u8 = undefined;
    const line = try reader.readUntilDelimiterOrEof(&buf, '\n') orelse {
        return error.EmptyInput;
    };

    var split = std.mem.splitScalar(u8, line, ',');

    var invalid_sum: u64 = 0;

    while (split.next()) |range| {
        var bounds = std.mem.splitScalar(u8, range, '-');

        const start_str = bounds.next() orelse return error.InvalidRange;
        const start = try std.fmt.parseInt(u64, start_str, 10);
        const end_str = bounds.next() orelse return error.InvalidRange;
        const end = try std.fmt.parseInt(u64, end_str, 10);

        for (start..end + 1) |
            value,
        | {
            var num_buffer: [32]u8 = undefined;
            var num_slice = try std.fmt.bufPrint(&num_buffer, "{}", .{value});
            for ([_]u64{ 2, 3, 5, 7, 11 }) |m| {
                if (m > num_slice.len) {
                    break;
                }

                if (@mod(num_slice.len, m) == 0) {
                    const p = @divExact(num_slice.len, m);
                    var fr = p;
                    var to = p * 2;

                    while (fr < num_slice.len and
                        std.mem.eql(
                            u8,
                            num_slice[fr - p .. fr],
                            num_slice[fr..to],
                        ))
                    {
                        fr += p;
                        to += p;
                    }
                    if (fr == num_slice.len) {
                        invalid_sum += value;
                        break;
                    }
                }
            }
        }
    }

    try stdout.print("Day 2, Part Two: {d}\n", .{invalid_sum});
}

pub fn main() !void {
    try part1();
    try part2();
}
