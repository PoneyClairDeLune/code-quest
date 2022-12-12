#include <stdio.h>
#include <stdbool.h>

int main() {
	bool lastBlank = false; // If the last input is blank
	bool continueLoop = true; // Loop toggle
	unsigned short calCount = 16384; // So I don't have to write 16384 everywhere...
	long calories[16384] = {0}; // Stores the list of calories
	unsigned short calIdx = 0; // The current index
	// Collect the list of calories
	do {
		printf("Calories carried: ");
		scanf("%ld", &calories[calIdx]);
		// Break out of the loop if there are two successive blank values
		if (calories[calIdx] < 1) {
			if (lastBlank) {
				continueLoop = false;
			};
			lastBlank = true;
		} else {
			lastBlank = false;
		};
		calIdx ++;
		// Break out of the loop if went out of the bound
		if (calIdx >= calCount) {
			continueLoop = false;
		};
	} while (continueLoop);
	lastBlank = false;
	continueLoop = true;
	calIdx = 0;
	long currentCal = 0;
	short currentElf = 0;
	long maxCals[3] = {0};
	short maxElves[3] = {0};
	while (continueLoop) {
		long thisCal = calories[calIdx];
		if (thisCal == 0) {
			if (lastBlank) {
				continueLoop = false;
			};
			lastBlank = true;
			currentCal = 0;
			currentElf ++;
		} else {
			lastBlank = false;
			currentCal += thisCal;
			if (currentCal > maxCals[0]) {
				// Just quickly get things done as the assignment demanded...
				maxCals[2] = maxCals[1];
				maxCals[1] = maxCals[0];
				maxCals[0] = currentCal;
				maxElves[2] = maxElves[1];
				maxElves[1] = maxElves[0];
				maxElves[0] = currentElf;
			};
		};
		calIdx ++;
		if (calIdx >= calCount) {
			continueLoop = false;
		};
	};
	printf("These elves have the most calories carried:\n");
	long long calorySum = 0;
	for (char i = 0; i < 3; i ++) {
		printf("#%04i%12i\n", maxElves[i] + 1, maxCals[i]);
		calorySum += maxCals[i];
	};
	printf("Total %11i\n", calorySum);
	return 0;
};
