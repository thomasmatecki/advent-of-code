const std = @import("std");
const math = @import("math");

const stdout = std.io.getStdOut().writer();

const Range = struct {
    start: u64,
    end: u64,
};

fn compareRanges(_: void, a: Range, b: Range) bool {
    var as: i64 = @intCast(a.start);
    var bs: i64 = @intCast(b.start);
    const cmp = as - bs;
    if (cmp != 0) {
        return cmp < 0;
    }

    as = @intCast(a.end);
    bs = @intCast(b.end);

    return (as - bs) < 0;
}

fn compareInts(_: void, a: u64, b: u64) bool {
    const lhs = [1]u64{a};
    const rhs = [1]u64{b};
    return std.mem.order(
        u64,
        &lhs,
        &rhs,
    ) == std.math.Order.lt;
}

fn findMinRange(val: u64, ranges: []const Range) ?Range {
    if (ranges.len == 0) {
        return null;
    }

    var low: usize = 0;
    var high: usize = ranges.len - 1;

    while (low < high) {
        const mid = (low + high) / 2;
        const mid_range = ranges[mid];
        if (mid_range.start <= val and val <= mid_range.end) {
            high = mid;
            low = mid;
        } else if (val < mid_range.start) {
            high = mid - 1;
        } else if (val > mid_range.end) {
            low = mid + 1;
        }
    }

    const min_range = ranges[high];

    if (val < min_range.start) {
        return null;
    }

    if (val > min_range.end) {
        return null;
    }

    return min_range;
}

const Database = struct {
    ranges: std.ArrayList(Range),
    ids: std.ArrayList(u64),

    pub fn load(allocator: std.mem.Allocator, path: []const u8) !Database {
        var file = try std.fs.cwd().openFile(path, .{});
        var reader = file.reader();

        var d = Database{
            .ranges = std.ArrayList(Range).init(allocator),
            .ids = std.ArrayList(u64).init(allocator),
        };

        var buf: [256 * 3]u8 = undefined;

        while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            if (line.len == 0) {
                break;
            }

            var split = std.mem.splitScalar(u8, line, '-');

            const start_str = split.next() orelse break;
            const end_str = split.next() orelse return error.InvalidInput;

            const r = Range{
                .start = try std.fmt.parseInt(u64, start_str, 10),
                .end = try std.fmt.parseInt(u64, end_str, 10),
            };

            try d.ranges.append(r);
        }

        std.mem.sort(Range, d.ranges.items, {}, compareRanges);

        while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            try d.ids.append(try std.fmt.parseInt(u64, line, 10));
        }

        std.mem.sort(u64, d.ids.items, {}, compareInts);

        return d;
    }

    pub fn scanRanges(self: *Database, val: u64) ?Range {
        for (self.ranges.items) |r| {
            if (r.start <= val and val <= r.end) {
                return r;
            }
        }
        return null;
    }

    pub fn deinit(self: *Database) void {
        self.ranges.deinit();
        self.ids.deinit();
    }
};

pub fn part1(allocator: std.mem.Allocator) !void {
    var db = try Database.load(allocator, "day5/input.txt");
    defer db.deinit();

    var total: u32 = 0;
    var rangeIdx = 0;

    for (db.ids.items) |id| {
        if (db.scanRanges(id) != null) {
            total += 1;
        }
    }

    try stdout.print(
        "Day 4, Part One: {d}\n",
        .{total},
    );
}

pub fn part2(allocator: std.mem.Allocator) !void {
    var db = try Database.load(allocator, "day5/input.txt");
    defer db.deinit();

    var cursor: u64 = 0;
    var total: u64 = 0;

    for (db.ranges.items) |r| {
        if (cursor <= r.start) {
            total += r.end - r.start + 1;
            cursor = r.end + 1;
        } else if (cursor < r.end) {
            total += r.end - cursor + 1;
            cursor = r.end + 1;
        }
    }

    try stdout.print(
        "Day 4, Part Two: {d}\n",
        .{total},
    );
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    try part1(allocator);
    try part2(allocator);
}
