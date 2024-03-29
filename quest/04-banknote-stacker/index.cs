using System;
using System.Text;

namespace BanknoteStacker {
	class Program {
		static string PadStart(string text, int length, string token) {
			if (token.Length > 1) {
				throw new ArgumentException("Token must be a single character long");
			};
			StringBuilder textBuffer = new StringBuilder("", length);
			int leftFill = length - text.Length;
			while (textBuffer.Length < leftFill) {
				textBuffer.Append(token);
			};
			textBuffer.Append(text);
			return textBuffer.ToString();
		}
		static int[] faceValues = {500, 200, 100, 50, 20, 10, 5, 2, 1};
		static void countBitStack(int bits, int[] sum) {
			Console.Write(PadStart(bits.ToString(), 8, " "));
			for (int i = 0; i < faceValues.Length; i ++) {
				short count = 0;
				while (bits >= faceValues[i]) {
					count ++;
					bits -= faceValues[i];
				};
				sum[i] += count;
				Console.Write(PadStart(count.ToString(), 5, " "));
			};
			Console.Write("\n");
		}
		static void Main(string[] args) {
			int[] salary = new int[5];
			for (int i = 0; i < salary.Length; i ++) {
				Console.Write("Salary of employee #{0}: ", i + 1);
				salary[i] = int.Parse(Console.ReadLine());
			};
			Console.WriteLine("\n  Salary  500  200  100   50   20   10    5    2    1");
			int[] faceSum = new int[faceValues.Length];
			for (int i = 0; i < faceSum.Length; i ++) {
				faceSum[i] = 0;
			};
			for (int i = 0; i < salary.Length; i ++) {
				countBitStack(salary[i], faceSum);
			};
			Console.Write("   Total");
			for (int i = 0; i < faceSum.Length; i ++) {
				Console.Write(PadStart(faceSum[i].ToString(), 5, " "));
			};
		}
	}
}
