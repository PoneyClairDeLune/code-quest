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
	var rLength int64 = 0
	var err error;
	for {
		fmt.Printf("Length of the square's side (1~20): ")
		var line string
		// Why does Go automatically read line in a space-delimited fashion?
		fmt.Scanln(&line)
		rLength, err = strconv.ParseInt(line, 10, 16)
		if rLength > 0 || rLength < 21 || err == nil {
			break
		}
	}
	length := int16(rLength)
	for line := int16(0); line < length; line ++ {
		for col := int16(0); col < length; col ++ {
			fmt.Printf(getChar(length, line, col))
		}
		fmt.Printf("\n")
	}
}
