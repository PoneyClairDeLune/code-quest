import "dart:io";
import "dart:convert";

String usedChar (int size, int line, int col) {
	int criteria = size - 1;
	if (
		line == 0 ||
		col == 0 ||
		line == criteria ||
		col == criteria
	) {
		return "*";
	} else {
		return " ";
	};
}

void main () {
	int size = 0;
	do {
		stdout.write("Length of the square's side (1~20): ");
		size = int.parse(stdin.readLineSync(encoding: utf8)!);
	} while (size < 1 || size > 20);
	for (int line = 0; line < size; line ++) {
		for (int col = 0; col < size; col ++) {
			stdout.write(usedChar(size, line, col));
		};
		stdout.write("\n");
	};
}
