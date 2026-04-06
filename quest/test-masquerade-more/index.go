package main

import (
	"fmt"
	"math/rand"
	_ "math/rand"
	"strconv"
	"time"
)

func GetCurlVersion() string {
	// curl 8.0.0 was released on 20/03/2023.
	var timeCurrent int64 = time.Now().Unix() / 86400
	var timeStart int64 = time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC).Unix() / 86400
	var minorValue int = int((timeCurrent - timeStart) / 57)
	// Generate some variations.
	minorValue -= 1 + rand.Intn(4)
	return "8." + strconv.Itoa(minorValue) + ".0"
}

func main() {
	fmt.Println("Approximated versions:")
	fmt.Printf("curl: %s\n", GetCurlVersion())
}
