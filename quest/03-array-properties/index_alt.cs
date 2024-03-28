using System;

namespace ArrayProps {
	class Program {
		static void TableFill(int[] targetArray) {
			for (int i = 0; i < targetArray.Length; i ++) {
				do {
					Console.Write("Provide figure #{0} (0~100): ", (i + 1));
					targetArray[i] = int.Parse(Console.ReadLine());
				} while (targetArray[i] < 1 || targetArray[i] > 100);
			};
		}
		static void TablePrint(int[] targetArray) {
			Console.Write("Among the array:");
			foreach (int e in targetArray) {
				Console.Write(" {0}", e);
			};
			Console.Write("\n");
		}
		static void PrintMax(int[] targetArray) {
			int max = 0;
			int at = 0;
			for (int i = 0; i < targetArray.Length; i ++) {
				if (targetArray[i] > max) {
					max = targetArray[i];
					at = i;
				};
			};
			Console.WriteLine("There exists a maximum of {0} at #{1}", max, (at + 1));
		}
		static void PrintMean(int[] targetArray) {
			int sum = 0;
			foreach (int e in targetArray) {
				sum += e;
			};
			Console.WriteLine("And a mean of {0}", ((float) sum) / 10);
		}
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
