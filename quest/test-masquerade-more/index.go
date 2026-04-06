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
	// Generate some variations upon the average release cadance of 56.67 days.
	var minorValue int = int((timeCurrent - timeStart) / 57)
	minorValue -= 1 + rand.Intn(4)
	return "8." + strconv.Itoa(minorValue) + ".0"
}
func GetFirefoxVersion() string {
	// Firefox 128 ESR was released on 09/07/2023.
	var timeCurrent int64 = time.Now().Unix() / 86400
	var timeStart int64 = time.Date(2024, 7, 29, 0, 0, 0, 0, time.UTC).Unix() / 86400
	// Generate some variations upon the average release cadance of 28 days, however Firefox had some delays averaging to 29.67 days in total.
	var timeDiff = timeCurrent - timeStart
	var majorValue int = int(timeDiff / 29) + 128
	majorValue -= 1 + rand.Intn(2)
	return strconv.Itoa(majorValue) + ".0"
}

func main() {
	fmt.Println("Approximated versions:")
	fmt.Printf("curl - %s\n", GetCurlVersion())
	fmt.Printf("Firefox - %s\n", GetFirefoxVersion())
}
