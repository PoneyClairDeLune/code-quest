using System;

namespace Square {
	class Program {
		static void Main(string[] args) {
			string[] kn = Console.ReadLine().Split(' ');
			byte n = byte.Parse(kn[0]);
			byte k = byte.Parse(kn[1]);
			string[] scores = Console.ReadLine().Split(' ');
			byte criteria = byte.Parse(scores[k]);
			byte advanceCount = 0;
			foreach (string scoreText in scores) {
				byte score = byte.Parse(scoreText);
				if (score >= criteria && score > 0) {
					advanceCount ++;
				};
			};
			Console.WriteLine(advanceCount);
		}
	}
}
