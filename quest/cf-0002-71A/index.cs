using System;

namespace TooLongWords {
	class Program {
		static void Main(string[] args) {
			ushort rounds = ushort.Parse(Console.ReadLine());
			for (ushort round = 0; round < rounds; round ++) {
				string word = Console.ReadLine();
				if (word.Length > 10) {
					Console.WriteLine("{0}{1}{2}", word[0], (word.Length - 2).ToString(), word[word.Length - 1]);
				} else {
					Console.WriteLine(word);
				};
			};
		}
	}
}
