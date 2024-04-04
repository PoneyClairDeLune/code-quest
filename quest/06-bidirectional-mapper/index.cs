using System;
using System.IO;

namespace BidirectionalMapper {
	class Program {
		static void Main(string[] args) {
			Console.Write("Search for word: ");
			string searchTerm = Console.ReadLine();
			bool foundYet = false;
			StreamReader dictStream = new StreamReader("./dict.txt");
			while (dictStream.Peek() > -1) {
				string dictLine = dictStream.ReadLine();
				int commaIdx = dictLine.IndexOf(",");
				if (commaIdx < 0) {
					continue;
				};
				string ponishWord = dictLine.Substring(0, commaIdx);
				string prenchWord = dictLine.Substring(commaIdx + 1);
				if (ponishWord.Equals(searchTerm)) {
					foundYet = true;
					Console.WriteLine("[Ponish] {0} [Prench] {1}", ponishWord, prenchWord);
				} else if (prenchWord.Equals(searchTerm)) {
					foundYet = true;
					Console.WriteLine("[Prench] {0} [Ponish] {1}", prenchWord, ponishWord);
				};
			};
			if (!foundYet) {
				Console.WriteLine("Word \"{0}\" could not be found.", searchTerm);
			};
		}
	}
}
