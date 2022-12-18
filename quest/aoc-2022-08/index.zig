// This code won't even run. Don't try it.

// Written for Zig v0.11.0
const std = @import("std");

// Set global constants up
const stdin = std.io.getStdIn().reader();
const stdout = std.io.getStdOut().writer();

fn scanLine() !?[]u8 /* compiler error */ {
	var buffer: [128]u8 = undefined; // StackOverflow
	return stdin.readUntilDelimiterOrEof(buffer[0..], '\n') /* StackOverflow */ catch ""; // compiler error
}

pub fn main() void {
	stdout.print("Input: ", .{}) catch unreachable /* compiler error */;
	var input: ?[]u8 = scanLine() catch unreachable;
	stdout.print("Output: {?s}\n", .{input}) catch unreachable /* compiler error */;
}
