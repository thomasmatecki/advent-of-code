const std = @import("std");

pub const day1 = @import("day1");
pub const day2 = @import("day2");
pub const day3 = @import("day3");

const re = @import("regex");

//const re = @cImport(@cInclude("regez.h"));

pub fn main() !void {

    //const input: [*:0]const u8 = "xxbcabac";
    var abc_expr = try re.regex("a(bc)");
    abc_expr.free();

    //const pattern: [*:0]const u8 = "hello[[:space:]]+([[:alpha:]]+)";
    //const ret = c.re_compile(regex, pattern, c.REG_EXTENDED);
    //if (ret != 0) {
    //var errbuf: [128]u8 = undefined;
    //_ = c.re_error(ret, regex, &errbuf, errbuf.len);
    //std.debug.print("Regex compile failed: {s}\n", .{std.mem.sliceTo(&errbuf, 0)});
    //return;
    //}
    //defer c.re_free(regex);

    //const input: [*:0]const u8 = "hello Zig";
    //var matches: [5]c.regmatch_t = undefined;

    //const exec_ret = c.re_exec(regex, input, matches.len, &matches, 0);
    //if (exec_ret == 0) {
    //std.debug.print("Matched!\n", .{});
    //for (matches, 0..) |m, i| {
    //if (m.rm_so == -1) break;
    //const start = @intCast(usize, m.rm_so);
    //const end = @intCast(usize, m.rm_eo);
    //const text = input[start..end];
    //std.debug.print("  group[{d}]: {s}\n", .{i, text});
    //}
    //} else {
    //std.debug.print("No match.\n", .{});
    //}
    //}

    //    try day1.part1();
    //    try day1.part2();
    //    try day2.part1();
    //    try day2.part2();
    //    try day3.part1();
}
