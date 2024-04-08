using System;

namespace Square {
	class Program {
		static void Main(string[] args) {
			string[] subs = Console.ReadLine().Split(" ");
			double n = double.Parse(subs[0]);
			double m = double.Parse(subs[1]);
			ulong a = ulong.Parse(subs[2]);
			double r = Math.Ceiling(n / a) * Math.Ceiling(m / a);
			Console.WriteLine(r.ToString("."));
		}
	}
}
