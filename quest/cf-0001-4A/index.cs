using System;

namespace Watermelon {
	class Program {
		static void Main(string[] args) {
			ushort weight = ushort.Parse(Console.ReadLine());
			Console.WriteLine((weight > 2 && (weight & 1) == 0) ? "YES" : "NO");
		}
	}
}
