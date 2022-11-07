#include <stdio.h>

char * usedChar(long length, char line, char col) {
	if (
		line * col == 0 ||
		line == length - 1 ||
		col == length - 1
	) {
	    return "*";
    } else {
    	return " ";
    };
};

int main() {
	long length;
	do {
		printf("Length of the square's side: ");
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