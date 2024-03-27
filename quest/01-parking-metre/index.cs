using System;
using System.Text;

namespace ParkingMetre {
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
		static double CalculateFee(double hours) {
			if (hours < 3) {
				return 2;
			} else if (hours < 19) {
				return 0.5 * hours + 0.5;
			} else {
				return 10;
			};
		}
		static void Main(string[] args) {
			double[] elapsed = new double[3];
			for (int i = 0; i < elapsed.Length; i ++) {
				Console.Write("How many hours did customer #{0} stay? ", i + 1);
				elapsed[i] = double.Parse(Console.ReadLine());
			};
			Console.WriteLine("Customer   Hours   Charge");
			for (int i = 0; i < elapsed.Length; i ++) {
				double hours = elapsed[i];
				Console.Write(PadStart((i + 1).ToString(), 8, " "));
				Console.Write(PadStart(hours.ToString(), 8, " "));
				Console.WriteLine(PadStart(CalculateFee(hours).ToString("n2"), 9, " "));
			};
		}
	}
}
