#include <stdio.h>
#include <stdbool.h>
#include <ctype.h>

short stringLength(char *string) {
	// This function calculates string length by filtering invalid charaters.
	// Feel free to rewrite this function to conduct matching in a single pass.
	short length = 0;
	while (string[length] > 31) {
		length ++;
	};
	return length;
};

bool suffixMatch(char *suffix, char *string) {
	// This function matches strings in a reverse order.
	short suffixLen = stringLength(suffix);
	short stringLen = stringLength(string);
	if (suffixLen > stringLen) {
		return false;
	};
	short matchSuffix = 0;
	for (short index = suffixLen - 1; index >= 0; index --) {
		// Input is normalized to let the matching be case-insensitive.
		char e0 = tolower(string[stringLen - suffixLen + index]);
		char e1 = suffix[index];
		if (e0 == e1) {
			matchSuffix ++;
		};
	};
	return (matchSuffix == suffixLen);
};

int main() {
	// Define suffix in lower case.
	char *suffix = "er";
	// Allocate space for string input.
	char strings[5][128];
	for (char i = 0; i < 5; i ++) {
		// Accepting string input.
		printf("Write a string (127 chars max): ");
		fgets(strings[i], sizeof strings[i], stdin);
	};
	printf("\n");
	// Accessing strings.
	for (char i = 0; i < 5; i ++) {
		if (suffixMatch(suffix, strings[i])) {
			// Since inputs contain line feeds, you don't need to add your own.
			printf("Suffix \"%s\" is found in string: %s", suffix, strings[i]);
		};
	};
	return 0;
};