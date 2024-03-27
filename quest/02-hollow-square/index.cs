using System;

namespace HollowSquare {
	class Program {
		static string UsedChar(long length, byte line, byte col) {
			int crit = line * col;
			if (
				crit == 0 ||
				crit % (length - 1) == 0
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
