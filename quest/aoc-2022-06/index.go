package main;

import (
	"fmt"
);

func main() {
	var packet string;
	fmt.Printf("Packet: ");
	fmt.Scanln(&packet);
	var startIdx int = 0;
	var msgIdx int = 0;
	var streamCount [26]uint16;
	var letterCount uint8 = 0;
	// Reset distinct letter counter
	for i := range streamCount {
		streamCount[i] = 0;
	};
	for i, e := range packet {
		// Stream header check
		if (i > 2) {
			hasSame := false;
			window := packet[i - 3 : i + 1];
			for i1 := 0; i1 < 3; i1 ++ {
				for i2 := i1 + 1; i2 < 4; i2 ++ {
					if (!hasSame) {
						hasSame = window[i1] == window[i2];
					};
				};
			};
			if (startIdx < 3 && !hasSame) {
				startIdx = i;
			};
		};
		// Message header check
		if (streamCount[e - 97] < 1) {
			letterCount ++;
		};
		streamCount[e - 97] ++;
		if (i > 12) {
			if (i > 13) {
				streamCount[packet[i - 14] - 97] --;
				if (streamCount[packet[i - 14] - 97] < 1) {
					letterCount --;
				};
			};
			if (msgIdx < 13) {
				if (letterCount >= 14) {
					msgIdx = i;
				};
			};
		};
		/* fmt.Printf("%02d(%02d):", i + 1, letterCount);
		for _, e1 := range streamCount {
			fmt.Printf(" %02d", e1);
		};
		fmt.Printf("\n"); */
	};
	fmt.Printf("Packet offset: %d\n", startIdx + 1);
	fmt.Printf("Message offset: %d\n", msgIdx + 1);
};
