const std = @import("std");

const stdout = std.io.getStdOut().writer();

pub fn parseRotation(rotation: []const u8) !struct {
    u8,
    i16,
} {
    if (rotation.len < 2) {
        return error.InvalidRotation;
    }
    const direction = rotation[0];
    const clicks = try std.fmt.parseInt(i16, rotation[1..], 10);
    return .{
        direction,
        clicks,
    };
}

pub fn part1() !void {
    var file = try std.fs.cwd().openFile("day1/input.txt", .{});
    var buffered_reader = std.io.bufferedReader(file.reader());
    const line_reader = buffered_reader.reader();

    var buf: [256]u8 = undefined;
    var position: i16 = 50;
    var zero_count: u32 = 0;

    while (try line_reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const rotation, const clicks = try parseRotation(line);

        switch (rotation) {
            'L' => position -= clicks,
            'R' => position += clicks,
            else => return error.InvalidRotation,
        }

        position = @mod(position, 100);

        if (position == 0) {
            zero_count += 1;
        }
    }

    try stdout.print("Day 1, Part One: {d}\n", .{zero_count});
}

pub fn part2() !void {
    var file = try std.fs.cwd().openFile("day1/input.txt", .{});
    var buffered_reader = std.io.bufferedReader(file.reader());
    const line_reader = buffered_reader.reader();

    var buf: [256]u8 = undefined;
    var position: i16 = 50;
    var zero_count: i32 = 0;

    while (try line_reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const direction, const clicks = try parseRotation(line);
        const mod_clicks = @mod(clicks, 100);
        var zero_clicks: i32 = @divFloor(clicks, 100);

        const new_position: i16 = switch (direction) {
            'L' => position - mod_clicks,
            'R' => position + mod_clicks,
            else => return error.InvalidRotation,
        };

        if (position != 0 and (new_position < 0 or new_position > 100)) {
            zero_clicks += 1;
        }

        position = @mod(new_position, 100);

        if (position == 0) {
            zero_clicks += 1;
        }

        zero_count += zero_clicks;
    }

    try stdout.print("Day 1, Part Two: {d}\n", .{zero_count});
}

pub fn main() !void {
    try part1();
    try part2();
}
