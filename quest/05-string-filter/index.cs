using System;
using System.Text;

namespace StringFilter {
	class Program {
		static void Main(string[] args) {
			string suffix = "er";
			string[] strings = new string[5];
			for (int i = 0; i < strings.Length; i ++) {
				Console.Write("Write a string: ");
				strings[i] = Console.ReadLine();
			};
			Console.Write("\n");
			foreach (string item in strings) {
				if (item.EndsWith(suffix)) {
					Console.WriteLine("Suffix \"{0}\" is found in string: {1}", suffix, item);
				};
			};
		}
	}
}
