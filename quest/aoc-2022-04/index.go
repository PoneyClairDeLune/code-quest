package main;

import "fmt";

func main() {
	var assignment string = "!";
	var curGroup [2][2]uint16;
	var smallerStart uint8 = 0;
	var containCount uint16 = 0;
	var intersectCount uint16 = 0;
	for (assignment[0] > 32) {
		// Clear the previous input
		assignment = " ";
		// Reset smallerStart
		smallerStart = 0;
		// Accept user input
		fmt.Printf("Assignment: ");
		fmt.Scanf("%s", &assignment);
		if (assignment[0] > 32) {
			// Parse input
			var scanMode byte = 0; // First dash, comma, second dash
			var lastIndex = 0;
			for i, e := range assignment {
				switch (scanMode) {
					case 0: {
						if (e == '-') {
							fmt.Sscan(assignment[lastIndex : i], &curGroup[0][0]);
							lastIndex = i + 1;
							scanMode = 1;
						};
						break;
					};
					case 1: {
						if (e == ',') {
							fmt.Sscan(assignment[lastIndex : i], &curGroup[0][1]);
							lastIndex = i + 1;
							scanMode = 2;
						};
						break;
					};
					case 2: {
						if (e == '-') {
							fmt.Sscan(assignment[lastIndex : i], &curGroup[1][0]);
							lastIndex = i + 1;
							scanMode = 3;
						};
						break;
					};
				};
			};
			fmt.Sscan(assignment[lastIndex:], &curGroup[1][1]);
			// Deciding the range with a smaller start
			if (curGroup[0][0] > curGroup[1][0]) {
				smallerStart = 1;
			} else if (curGroup[0][0] == curGroup[1][0] && curGroup[0][1] > curGroup[1][1]) {
				smallerStart = 1;
			};
			for i0, e0 := range curGroup {
				for i1 := range e0 {
					fmt.Printf("%d ", curGroup[uint8(i0) ^ smallerStart][i1]);
				};
				fmt.Printf("\n");
			};
			// Containment counting
			if (curGroup[smallerStart][1] >= curGroup[smallerStart ^ 1][1]) {
				containCount ++;
			};
			if (curGroup[smallerStart][0] == curGroup[smallerStart ^ 1][0] && curGroup[smallerStart][1] < curGroup[smallerStart ^ 1][1]) {
				containCount ++;
			};
			// Intersection counting
			if (curGroup[smallerStart][1] >= curGroup[smallerStart ^ 1][0]) {
				intersectCount ++;
			};
		};
	};
	fmt.Printf("Assignments fully contained: %d\n", containCount);
	fmt.Printf("Assignments intersected: %d\n", intersectCount);
};
