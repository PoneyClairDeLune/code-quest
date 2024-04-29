using System;

namespace ArrayProps {
	class Program {
		static void Main(string[] args) {
			int[] dummy = new int[10];
			Console.Write("Among the array:");
			int max = 0;
			int at = 0;
			int sum = 0;
			for (int i = 0; i < dummy.Length; i ++) {
				int e = dummy[i];
				Console.Write(" {0}", e);
				if (max < e) {
					max = e;
					at = i;
				};
				sum += e;
			};
			Console.Write("\n");
			Console.WriteLine("There exists a maximum of {0} at #{1}", max, (at + 1));
			Console.WriteLine("And a mean of {0}", ((float) sum) / 10);
		}
	}
}
