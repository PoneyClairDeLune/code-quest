package main;

import (
	"fmt"
	"bufio"
	"os"
);

func main() {
	var instruction string = " ";
	var inputMode byte = 0; // Tower
	var stackTower [][]byte;
	stdinScan := bufio.NewScanner(os.Stdin);
	for (len(instruction) > 0) {
		// Clear buffer and accepting user input
		instruction = "";
		fmt.Printf("In mode %d: ", inputMode);
		stdinScan.Scan();
		instruction = stdinScan.Text();
		if (inputMode == 1) {
			instruction = " ";
		};
		if (len(instruction) > 0 && instruction[0] > 31) {
			switch (inputMode) {
				case 0: {
					// Tower parsing mode
					if (instruction[1] > 47 && instruction[1] < 58) {
						inputMode = 1;
					} else {
						var towers uint16 = uint16(len(instruction)) / 4 + 1;
						for i := uint16(0); i < towers; i ++ {
							e := instruction[i * 4 + 1];
							if (len(stackTower) < int(i + 1)) {
								stackTower = append(stackTower, []byte{});
							};
							if (e > 32) {
								stackTower[i] = append(stackTower[i], e);
							};
						};
					};
					break;
				};
				case 1: {
					// Ending tower parsing
					inputMode = 2;
					for i0, e0 := range stackTower {
						fmt.Printf("Tower #%02d:", i0 + 1);
						last := len(e0) - 1;
						half := (last >> 1) + 1;
						for i1, e1 := range e0 {
							if (i1 < half) {
								e0[i1], e0[last - i1] = e0[last - i1], e1;
							};
						};
						for _, e1 := range e0 {
							fmt.Printf(" %c", e1);
						};
						fmt.Printf("\n");
					};
					break;
				};
				case 2: {
					// Instruction parsing
					var moveMode byte = 0;
					var moveCount rune = 0;
					var moveSource rune = 0;
					var moveTarget rune = 0;
					for _, e := range instruction {
						switch (moveMode) {
							case 0: {
								if (e == 32) {
									moveMode = 1;
								};
								break;
							};
							case 1: {
								if (e == 32) {
									moveMode = 2;
								} else {
									moveCount *= 10;
									moveCount += e - 48;
								};
								break;
							};
							case 2: {
								if (e == 32) {
									moveMode = 3;
								};
								break;
							};
							case 3: {
								if (e == 32) {
									moveMode = 4;
								} else {
									moveSource *= 10;
									moveSource += e - 48;
								};
								break;
							};
							case 4: {
								if (e == 32) {
									moveMode = 5;
								};
								break;
							};
							case 5: {
								if (e == 32) {
									moveMode = 6;
								} else {
									moveTarget *= 10;
									moveTarget += e - 48;
								};
								break;
							};
						};
					};
					moveSource --;
					moveTarget --;
					// Moving crates
					// Part 1 moving behaviour
					/* for mI := rune(0); mI < moveCount; mI ++ {
						fmt.Printf("Moving %c: S %d T %d\n", stackTower[moveSource][len(stackTower[moveSource]) - 1], moveSource, moveTarget);
						stackTower[moveTarget] = append(stackTower[moveTarget], stackTower[moveSource][len(stackTower[moveSource]) - 1]);
						stackTower[moveSource] = stackTower[moveSource][ : len(stackTower[moveSource]) - 1];
					}; */
					// Part 2 moving behaviour
					stackTower[moveTarget] = append(stackTower[moveTarget], stackTower[moveSource][len(stackTower[moveSource]) - int(moveCount) : ]...);
					stackTower[moveSource] = stackTower[moveSource][ : len(stackTower[moveSource]) - int(moveCount)];
					for i0, e0 := range stackTower {
						fmt.Printf("Tower #%02d:", i0 + 1);
						for _, e1 := range e0 {
							fmt.Printf(" %c", e1);
						};
						fmt.Printf("\n");
					};
					break;
				};
			};
		};
	};
};
