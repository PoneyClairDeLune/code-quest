package main;

import "fmt";

func groundPriority(e rune) int8 {
	if (e > 96) {
		return int8(e - 97); // Lower-case letter A
	} else {
		return int8(e - 39); // Upper-case letter A plus 26
	};
};

func main() {
	// Where the input is accepted
	var sackContent string = "A";
	// Initialize the sack matches
	var sackHas [5][52]bool; // 0 and i for part 1, all else for part 2
	// A place for calculating sum
	var prioritySum int32 = 0;
	var badgeSum int32 = 0;
	var badgeIndex uint16 = 0;
	// Accepting input for each round
	for (sackContent[0] > 64) {
		// Nullifying the string
		sackContent = " ";
		// Nullifying matches for each elf
		for i := range sackHas[0] {
			sackHas[0][i] = false;
		};
		for i := range sackHas[1] {
			sackHas[1][i] = false;
		};
		if (badgeIndex % 3 == 0) {
			// Nullifying matches for each group
			for i := range sackHas[2] {
				sackHas[2][i] = false;
				sackHas[3][i] = false;
				sackHas[4][i] = false;
			};
		};
		// Accepting input
		fmt.Printf("Sack content: ");
		fmt.Scanf("%s", &sackContent);
		if (sackContent[0] > 64) {
			var breakPoint = len(sackContent) / 2;
			for i, e := range sackContent {
				// Filling the sacks for individual elves
				var priority int8 = groundPriority(e);
				sackHas[i / breakPoint][priority] = true;
				// Filling the sacks for groups of elves
				sackHas[badgeIndex % 3 + 2][priority] = true;
			};
			for i, e := range sackHas[0] {
				if (e && sackHas[1][i]) {
					prioritySum += int32(i + 1);
				};
			};
			if (badgeIndex % 3 == 2) {
				for i, e := range sackHas[2] {
					if (e && sackHas[3][i] && sackHas[4][i]) {
						badgeSum += int32(i + 1);
					};
				};
			};
		};
		badgeIndex ++;
	};
	fmt.Printf("Sum for individual elves: %d\n", prioritySum);
	fmt.Printf("Sum for groups of elves: %d\n", badgeSum);
};
