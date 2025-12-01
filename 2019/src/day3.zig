const std = @import("std");
const aoc: type = @import("aoc");

const stdout = std.io.getStdOut().writer();

pub fn parseInstruction(instruction: []const u8) !struct {
    u8,
    i32,
} {
    if (instruction.len < 2) {
        return error.InvalidInstruction;
    }
    const direction = instruction[0];
    const distance = try std.fmt.parseInt(i32, instruction[1..], 10);
    return .{
        direction,
        distance,
    };
}

const MoveIterator = struct {
    step: i32 = 0,
    direction: u8,
    distance: i32,
    point: *Point,

    pub fn next(self: *MoveIterator) ?Point {
        if (self.step >= self.distance) {
            return null;
        }

        self.point.step(self.direction);

        self.step += 1;

        return self.point.*;
    }
};

const Point = struct {
    x: i32 = 0,
    y: i32 = 0,

    fn step(self: *Point, direction: u8) void {
        switch (direction) {
            'U' => self.y += 1,
            'D' => self.y -= 1,
            'R' => self.x += 1,
            'L' => self.x -= 1,
            else => {},
        }
    }

    fn moveIterator(self: *Point, direction: u8, distance: i32) MoveIterator {
        return MoveIterator{
            .direction = direction,
            .distance = distance,
            .point = self,
        };
    }
};

const Path = struct {
    buf: ?[]u8 = undefined,
    split: std.mem.SplitIterator(u8, .sequence),
    moves: ?MoveIterator = null,
    position: Point = Point{},

    pub fn init(buf: ?[]u8) Path {
        const split = std.mem.splitSequence(u8, buf.?, ",");

        return Path{
            .buf = buf,
            .split = split,
        };
    }

    pub fn moreMoves(self: *Path) !bool {
        if (self.split.next()) |instruction| {
            const direction, const distance = try parseInstruction(instruction);
            self.moves = self.position.moveIterator(direction, distance);
            return true;
        } else return false;
    }

    pub fn next(self: *Path) !?Point {
        if (self.moves == null and !try self.moreMoves()) {
            return null;
        }

        if (self.moves.?.next()) |point| {
            return point;
        } else {
            self.moves = null;
            return try self.next();
        }
    }
};

const max_u32 = std.math.maxInt(u32);

pub fn part1() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    const allocator = arena.allocator();

    var file = try std.fs.cwd().openFile("day3/input.txt", .{});
    var buffered_reader = std.io.bufferedReader(file.reader());
    const line_reader = buffered_reader.reader();

    var min: u32 = max_u32;

    const buf0 = try line_reader.readUntilDelimiterOrEofAlloc(allocator, '\n', 8 * 256);
    var path = Path.init(buf0);

    var ps = std.hash_map.AutoHashMap(Point, void).init(allocator);
    defer ps.deinit();

    while (try path.next()) |point| {
        try ps.put(point, {});
    }

    const buf1 = try line_reader.readUntilDelimiterOrEofAlloc(allocator, '\n', 8 * 256);
    defer allocator.free(buf1.?);

    var path1 = Path.init(buf1);

    while (try path1.next()) |point| {
        if (ps.contains(point)) {
            const manhattan = @abs(point.x) + @abs(point.y);
            min = @min(min, manhattan);
        }
    }

    try stdout.print("Day 1, Part One: {d}\n", .{min});
}

pub fn main() !void {
    try part1();
    //    try part2();
}
