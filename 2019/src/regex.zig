//! A simple Zig wrapper around a C regex library.

const std = @import("std");

const re = @cImport(@cInclude("regez.h"));

const RegexErrors = error{
    CompileFailed,
    ExecFailed,
};

/// A compiled regular expression.
const Regex = struct {
    regex: ?*re.struct_regex_wrapper,
    //    nsub: usize = 0,

    pub fn free(self: *Regex) void {
        re.regex_free(self.regex);
    }

    pub fn match(
        self: *const Regex,
        input: []const u8,
    ) ![]re.regmatch_t {
        var matches: [1]re.regmatch_t = undefined;

        const exec_rc = re.regex_exec(
            self.regex,
            input.ptr,
            &matches,
            0,
        );

        if (exec_rc == 0) {
            return matches[0..1];
        } else {
            return RegexErrors.ExecFailed;
        }
    }
};

pub fn regex(
    pattern: []const u8,
) !Regex {
    const regex_ptr = re.regex_compile(
        pattern.ptr,
        re.REG_EXTENDED,
    );

    if (regex_ptr == null) {
        return RegexErrors.CompileFailed;
    }

    return Regex{
        .regex = regex_ptr,
    };
}
