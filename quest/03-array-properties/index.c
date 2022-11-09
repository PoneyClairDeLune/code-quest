#include <stdio.h>

// The design of this quest is kind of... out of their mind.
// But I'm here to solve, not to complain, so whatever.

void tableFill(char targetArray[], char length) {
	for (char i = 0; i < length; i ++) {
		do {
			printf("Provide figure #%02i (0~100): ", i + 1);
			scanf("%d", &targetArray[i]);
		} while (targetArray[i] < 0 || targetArray[i] > 100);
	};
};

void tablePrint(char targetArray[], char length) {
	printf("Among the array:");
	for (char i = 0; i < length; i ++) {
		printf(" %i", targetArray[i]);
	};
	printf("\n");
};

void printMax(char targetArray[], char length) {
	char max = 0;
	char at = 0;
	for (char i = 0; i < length; i ++) {
		if (max < targetArray[i]) {
			max = targetArray[i];
			at = i;
		};
	};
	printf("There exists a maximum of %i at #%02i\n", max, at + 1);
};

void printMean(char targetArray[], char length) {
	int sum = 0;
	for (char i = 0; i < length; i ++) {
		sum += targetArray[i];
	};
	printf("And a mean of %.1f\n", (float)sum / length);
};

int main() {
	char length = 10;
	char dummy[length];
	tableFill(dummy, length);
	tablePrint(dummy, length);
	printMax(dummy, length);
	printMean(dummy, length);
	return 0;
};