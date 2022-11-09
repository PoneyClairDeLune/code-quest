#include <stdio.h>

// Since every single function is of single use anyway,
// there is no point to ever truly seperate the variables.
// So for the sake of reducing CPU time,
// why not pack them up, and hack our way through?

int main() {
	char numbers[10];
	printf("Please provide the numbers accordingly.\n");
	for (char i = 0; i < 10; i ++) {
		do {
			printf("Figure %02i (1~100): ", i + 1);
			scanf("%d", &numbers[i]);
		} while (numbers[i] < 1 || numbers[i] > 100);
	};
	char smallest = 100;
	char largest = 0;
	int sum = 0;
	printf("\nFor the given array of figures:\n");
	for (char i = 0; i < 10; i ++) {
		int e = numbers[i];
		if (e < smallest) {
			smallest = e;
		};
		if (e > largest) {
			largest = e;
		};
		sum += e;
		printf("%i ", e);
	};
	float mean = (float)sum / 10;
	printf("\n\nHas the following properties:\nMaximum: %i\nMinimum: %i\nMean: %.1f\n", largest, smallest, mean);
	return 0;
};