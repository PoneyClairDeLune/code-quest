#include <stdio.h>

// Since every single function is of single use anyway,
// there is no point to ever truly seperate the variables.
// So for the sake of reducing CPU time,
// why not pack them up, and hack our way through?

char smallest = 100;
char smallIndex, largest, largeIndex = 0;
int sum = 0;

void tableFill(char numbers[], size_t length) {
	printf("Please provide the numbers accordingly.\n");
	for (char i = 0; i < length; i ++) {
		do {
			printf("Figure %02i (0~100): ", i + 1);
			scanf("%d", &numbers[i]);
		} while (numbers[i] < 0 || numbers[i] > 100);
	};
	return;
};

void tablePrint(char numbers[], size_t length) {
	printf("\nFor the given array:");
	for (char i = 0; i < length; i ++) {
		char e = numbers[i];
		if (e < smallest) {
			smallest = e;
			smallIndex = i;
		};
		if (e > largest) {
			largest = e;
			largeIndex = i;
		};
		sum += e;
		printf(" %i", e);
	};
	return;
};

void statPrint(size_t length) {
	float mean = (float)sum / length;
	printf("\n\nHas the following properties:\nMaximum: %i (#%02i)\nMinimum: %i (#%02i)\nMean: %.1f\n", largest, largeIndex + 1, smallest, smallIndex + 1, mean);
	return;
};

int main() {
	char numbers[10];
	tableFill(numbers, 10);
	tablePrint(numbers, 10);
	statPrint(10);
	return 0;
};