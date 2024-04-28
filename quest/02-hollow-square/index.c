#include <stdio.h>

char * usedChar(long length, char line, char col) {
	int crit = length - 1;
	if (
		line == 0 ||
		col == 0 ||
		line == crit ||
		col == crit
	) {
		return "*";
	} else {
		return " ";
	};
};

int main() {
	long length;
	do {
		printf("Length of the square's side (1~20): ");
		scanf("%ld", &length);
	} while (length < 1 || length > 20);
	for (char line = 0; line < length; line ++) {
		for (char col = 0; col < length; col ++) {
			printf(usedChar(length, line, col));
		};
		printf("\n");
	};
	return 0;
};
