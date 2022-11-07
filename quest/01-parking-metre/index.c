#include <stdio.h>
#include <math.h>

// Treat curly braces {} as scopes. A child scope can access variables within the same scope or in parent scopes, but not other scopes in the same level or its very own child scope.

double calculateFee(double time) {
	if (time >= 19) {
		return 10; // Remember to return the result.
	} else if (time > 3) {
		/*
		Do not write redundant conditions if they cannot be matched.
		Like in this case, do not write "time < 19" when it cannot be matched anyway.
		*/
		return 0.5 * time + 0.5;
	} else {
		return 2;
	};
};

int main() {
	double customers[3]; // Declare variables within the same scope or in the parent scope.
	/*
	Array starts at index 0, but its length isn't. The last index of an array equals to the length of said array subtracted by one.
	*/
	for (int c = 0; c < 3; c ++) {
		printf("How many hours has customer #%i stayed? ", c + 1);
		scanf("%lf", &customers[c]); // No need for line feeds when you'll hit enter anyway
	};
	printf("Customer    Hours    Charge\n"); // Mould your format beforehand
	double totalHour = 0;
	for (char c = 0; c < 3; c ++) {
		double fee = calculateFee(customers[c]);
		printf("%8i %8.2lf %9.2lf\n", c + 1, customers[c], fee); // Pay attention to string templates used
		totalHour += customers[c];
	};
	printf("   TOTAL%9.2lf%10.2lf\n", totalHour, calculateFee(totalHour));
	return 0; // 0 means success, 1 means failure
};
