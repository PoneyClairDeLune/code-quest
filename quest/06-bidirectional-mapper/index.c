#include <stdio.h>
#include <stdbool.h>

long long fileLength(char *path) {
	// Get the size of the file
	// However it later proved to be useless, so ignore this part please
	long long length = 0;
	FILE *target = fopen(path, "r");
	char unit = fgetc(target);
	while (unit != EOF) {
		length ++;
		unit = fgetc(target);
	};
	fclose(target);
	return length;
};

bool isSameStr(char *a, char *b) {
	bool same = true;
	bool terminated = false;
	long long i = 0;
	while (same && !terminated) {
		if (a[i] < 32 || b[i] < 32) {
			terminated = true;
		};
		same = a[i] == b[i];
		i ++;
	};
	return same;
};

long entryCount(char *path) {
	// Get the number of entries in the file
	bool isCtrl = true; // If the last character is a control character
	long count = 0; // The incremental counter of entries
	FILE *target = fopen(path, "r");
	char unit = fgetc(target); // The current iterated character
	while (unit != EOF) {
		if (unit < 32) {
			// Control characters
			isCtrl = true;
		} else {
			if (isCtrl) {
				count ++; // Increase the length
				isCtrl = false; // Reset control state
			};
		};
		unit = fgetc(target);
	};
	fclose(target);
	return count;
};

void entryLoad(char dict[][2][65], char *path) {
	// Load the map into a two dimentional string array
	FILE *target = fopen(path, "r");
	bool isCtrl = true;
	char col = 0; // 0 for key, 1 for value
	char row = 0;
	short index = -1;
	char unit = fgetc(target); // The iterated current character
	while (unit != EOF) {
		if (unit < 32) {
			// Control characters
			isCtrl = true;
		} else {
			if (isCtrl) {
				// Reset into a new entry
				dict[index][col][row] = 0; // Add a NULL to terminate the string
				index ++;
				col = 0;
				row = 0;
				isCtrl = false;
				//printf("CTRL: %i %i %i\n", index, col, row);
			};
			if (unit == ',') {
				// No need for strtok
				dict[index][col][row] = 0; // Add a NULL to terminate the string
				row = 0; // Reset the string pointer
				col ++; // Switch to the next column
				//printf("JUMP: %i %i %i\n", index, col, row);
			} else {
				dict[index][col][row] = unit;
				//printf("STOR: %i %i %i %c\n", index, col, row, unit);
				row ++;
			};
		};
		unit = fgetc(target);
	};
	dict[index][col][row] = 0; // Add a NULL to terminate the string
	fclose(target);
	return;
};

int main() {
	char *path = "./dict.txt"; // Define the target file path
	long length = entryCount(path); // Determine how many entries are in the map
	char dictMap[length][2][65]; // Allocate the map accordingly
	entryLoad(dictMap, path);
	char term[65] = {0};
	printf("Search for word: ");
	fscanf(stdin, "%s", term);
	bool notFound = true;
	// Search in Ponish first, then in Prench
	long i0 = 0; // Ponish or Prench
	long i1 = 0; // Index
	while (notFound && i0 < 2) {
		while (notFound && i1 < length) {
			notFound = !isSameStr(term, dictMap[i1][i0]);
			if (notFound) {
				i1 ++;
			} else {
				printf("Found at %i %i\n", i1, i0);
			};
		};
		if (notFound) {
			i1 = 0;
			i0 ++;
		};
	};
	if (notFound) {
		printf("Word \"%s\" could not be found.", term);
	} else {
		if (i0 == 1) {
			printf("[Prench] %s [Ponish] %s", term, dictMap[i1][1 - i0]);
		} else {
			printf("[Ponish] %s [Prench] %s", term, dictMap[i1][1 - i0]);
		};
	};
	printf("\n");
	return 0;
};