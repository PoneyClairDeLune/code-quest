#include <stdio.h>

char faceValueCount = 9;
short faceValues[] = {500, 200, 100, 50, 20, 10, 5, 2, 1}; // Easier to optimize when organized this way

void showBitStack(int bits, short sum[]) {
	printf("%8i", bits);
	// Show the bit stacks in one go
	for (char i = 0; i < faceValueCount; i ++) {
		short count = 0;
		while (bits >= faceValues[i]) {
			count ++;
			bits -= faceValues[i];
		};
		sum[i] += count;
		printf("%5i", count);
	};
	printf("\n");
	return;
};

int main() {
	char employeeCount = 5;
	int salary[employeeCount];
	// Record the salary of each employee
	for (char i = 0; i < 5; i ++) {
		printf("Salary of employee #%i: ", i + 1);
		scanf("%d", &salary[i]);
	};
	// Show how the salary should be given in terms of bits
	printf("\n  Salary  500  200  100   50   20   10    5    2    1\n");
	short faceSum[faceValueCount];
	// Initialize the array
	for (char i = 0; i < faceValueCount; i ++) {
		faceSum[i] = 0;
	};
	for (char i = 0; i < 5; i ++) {
		showBitStack(salary[i], faceSum);
	};
	// Show the total bit counts
	printf("   Total");
	for (char i = 0; i < faceValueCount; i ++) {
		printf("%5i", faceSum[i]);
	};
	printf("\n");
	return 0;
};