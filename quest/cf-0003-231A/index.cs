using System;

namespace Team {
	class Program {
		static byte CountResults(string text) {
			byte count = 0;
			for (int i = 0; i < text.Length; i ++) {
				if (text[i] == '1') {
					count ++;
				};
			};
			return count;
		}
		static void Main(string[] args) {
			ushort rounds = ushort.Parse(Console.ReadLine());
			ushort validRuns = 0;
			for (ushort round = 0; round < rounds; round ++) {
				if (CountResults(Console.ReadLine()) > 1) {
					validRuns ++;
				};
			};
			Console.WriteLine(validRuns.ToString());
		}
	}
}
