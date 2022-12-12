#include <stdio.h>
#include <stdbool.h>

int main() {
	bool lastBlank = false;
	bool continueLoop = true;
	unsigned short calCount = 16384;
	long calories[16384] = {0};
	unsigned short calIdx = 0;
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
		if (calIdx >= calCount) {
			continueLoop = false;
		};
	} while (continueLoop);
	// Print the array out for debugging
	/* for (char i = 0; i < calCount; i ++) {
		printf("%7i ", calories[i]);
	}; */
	// Try to find the elf with the most calories
	lastBlank = false;
	continueLoop = true;
	calIdx = 0;
	long currentCal = 0;
	short currentElf = 0;
	long maxCal = 0;
	short maxElf = 0;
	while (continueLoop) {
		long thisCal = calories[calIdx];
		if (thisCal == 0) {
			if (lastBlank) {
				continueLoop = false;
				//printf("Loop exited at %i\n", calIdx);
			};
			lastBlank = true;
			currentCal = 0;
			currentElf ++;
			//printf("Next elf!\n");
		} else {
			lastBlank = false;
			currentCal += thisCal;
			//printf("Elf #%02i has %i calories\n", currentElf + 1, currentCal);
			if (currentCal > maxCal) {
				maxCal = currentCal;
				maxElf = currentElf;
				//printf("Record broke!\n");
			};
		};
		calIdx ++;
		if (calIdx >= calCount) {
			continueLoop = false;
		};
	};
	printf("Elf #%02i carries the most calories: %i\n", maxElf + 1, maxCal);
	return 0;
};
