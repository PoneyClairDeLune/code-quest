package main

import (
	"fmt"
	"strconv"
)

func getChar(length int64, line int64, col int64) string {
	var crit int64 = line * col
	if crit == 0 || crit % (length - 1) == 0 {
		return "*"
	} else {
		return " "
	}
}

func main() {
	var length int64 = 0
	var err error;
	for ok := true; ok; ok = (length < 1 || length > 20 || err != nil) {
		fmt.Printf("Length of the square's side (1~20): ")
		var line string
		// Why does Go automatically read line in a space-delimited fashion?
		fmt.Scanln(&line)
		length, err = strconv.ParseInt(line, 10, 64)
		//fmt.Printf("Input value: %d\n", length)
	}
	for line := int64(0); line < length; line ++ {
		for col := int64(0); col < length; col ++ {
			fmt.Printf(getChar(length, line, col))
		}
		fmt.Printf("\n")
	}
}
