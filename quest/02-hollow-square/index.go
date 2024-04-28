package main

import (
	"fmt"
	"strconv"
)

func getChar(length int16, line int16, col int16) string {
	var crit int16 = length - 1
	if line == 0 || line == crit || col == 0 || col == crit {
		return "*"
	} else {
		return " "
	}
}

func main() {
	var length int16 = 0
	var err error;
	for ok := true; ok; ok = (length < 1 || length > 20 || err != nil) {
		fmt.Printf("Length of the square's side (1~20): ")
		var line string
		// Why does Go automatically read line in a space-delimited fashion?
		fmt.Scanln(&line)
		length, err = strconv.ParseInt(line, 10, 64)
		//fmt.Printf("Input value: %d\n", length)
	}
	for line := int16(0); line < length; line ++ {
		for col := int16(0); col < length; col ++ {
			fmt.Printf(getChar(length, line, col))
		}
		fmt.Printf("\n")
	}
}
