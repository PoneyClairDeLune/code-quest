const std = @import("std");

pub fn main() void {
	std.debug.print("Hi {s}!\n", .{"Luna"});
}