using System;

namespace HollowSquare {
	class Program {
		static string UsedChar(long length, byte line, byte col) {
			long crit = length - 1;
			if (
				line == 0 ||
				col == 0 ||
				line == crit ||
				col == crit
			) {
				return "*";
			} else {
				return " ";
			};
		}
		static void Main(string[] args) {
			long length = 0;
			do {
				Console.Write("Length of the square's side (1~20): ");
				length = long.Parse(Console.ReadLine());
			} while (length < 1 || length > 20);
			for (byte line = 0; line < length; line ++) {
				for (byte col = 0; col < length; col ++) {
					Console.Write(UsedChar(length, line, col));
				};
				Console.Write("\n");
			};
		}
	}
}
