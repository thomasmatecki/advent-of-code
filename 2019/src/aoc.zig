const std = @import("std");

fn Iterator(comptime T: type) type {
    return struct {
        file: std.fs.File,
        file_reader: std.fs.File.Reader,
        buffered_reader: std.io.BufferedReader(4096, std.fs.File.Reader),
        line_reader: std.io.AnyReader,
        buf: [32]u8 = undefined,
        delimiter: u8 = undefined,

        pub fn next(self: *Iterator(T)) !?T {
            if (try self.line_reader.readUntilDelimiterOrEof(&self.buf, self.delimiter)) |line| {
                return try std.fmt.parseInt(T, line, 10);
            } else {
                return null;
            }
        }
        pub fn close(self: *Iterator(T)) void {
            self.file.close();
        }
        pub fn take(self: *Iterator(T), comptime i: T) ![i]?T {
            var result: [i]?T = undefined;
            var idx: T = 0;
            while (idx < i) : (idx += 1) {
                result[idx] = try self.next();
                if (result[idx] == null) break;
            }
            return result;
        }
        pub fn collect(
            self: *Iterator(T),
            allocator: std.mem.Allocator,
        ) !std.ArrayList(T) {
            var list = std.ArrayList(T).init(allocator);
            while (try self.next()) |i| {
                try list.append(i);
            }
            return list;
        }
    };
}

pub fn inputIterator(
    comptime T: type,
    path: []const u8,
    delimiter: u8,
) !Iterator(T) {
    var file = try std.fs.cwd().openFile(path, .{});
    const file_reader = file.reader();
    var buffered_reader = std.io.bufferedReader(file_reader);

    return Iterator(T){
        .file = file,
        .file_reader = file_reader,
        .buffered_reader = buffered_reader,
        .line_reader = buffered_reader.reader().any(),
        .delimiter = delimiter,
    };
}
