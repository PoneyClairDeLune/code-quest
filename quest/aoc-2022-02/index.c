#include <stdio.h>

char getReqScore(char req){
	return req - 64; // 0x65 is letter A
};

char getResScore(char res){
	return res - 87; // 0x88 is letter X
};

char getWinCriteria(char score) {
	// For part 1, deciding the correct ID to win
	return score % 3 + 1;
};

char getResponse(char req, char strategy) {
	// For part 2, deciding the correct response
	return ((req - 65) + (strategy - 89) + 3) % 3 + 1;
};

char getResultScoreOld(char req, char res) {
	// For part 1, calculating the score for the round
	char opponent = getReqScore(req);
	char result = getResScore(res);
	char crit = getWinCriteria(opponent);
	//printf("Choices: %i %i\n", opponent, result);
	if (result == crit) {
		//printf("You won!\n");
		result += 6;
	} else if (result == opponent) {
		//printf("You draw!\n");
		result += 3;
	} else {
		//printf("You lose!\n");
	};
	return result;
};

char getResultScore(char req, char res) {
	char opponent = getReqScore(req);
	char response = getResponse(req, res);
	response += 3 * (res - 88);
	return response;
};

int main() {
	char input[17] = {0};
	long sumScore = 0;
	//printf("Enter numbers to quit.\n");
	do {
		// Accepting input
		printf("Round: ");
		fgets(input, sizeof input, stdin);
		// Calculating score
		if (input[0] > 64) {
			sumScore += getResultScore(input[0], input[2]);
		};
	} while (input[0] > 64);
	printf("Your total score: %i\n", sumScore);
	return 0;
};
