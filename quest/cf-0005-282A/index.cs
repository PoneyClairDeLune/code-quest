using System;

namespace Square {
	class Program {
		static void Main(string[] args) {
			short rounds = byte.Parse(Console.ReadLine());
			short x = 0;
			for (short round = 0; round < rounds; round ++) {
				string instruction = Console.ReadLine();
				if (instruction.IndexOf("++") > -1) {
					x ++;
				} else if (instruction.IndexOf("--") > -1) {
					x --;
				};
			};
			Console.WriteLine(x.ToString());
		}
	}
}
