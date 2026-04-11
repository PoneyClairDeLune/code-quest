package main

import (
	"fmt"
	"math"
	"math/rand"
	_ "math/rand"
	"strconv"
	"time"
)

func GetChromeVersion() string {
	// Start from Chrome 144 released on 2026.1.13
	var startVersion int = 144
	var timeStart int64 = time.Date(2026, 1, 13, 0, 0, 0, 0, time.UTC).Unix() / 86400
	var timeCurrent int64 = time.Now().Unix() / 86400
	var timeDiff int = int((timeCurrent - timeStart - 30)) - int(math.Floor(math.Pow(rand.Float64(), 2) * 90))
	return strconv.Itoa(startVersion + (timeDiff / 30)) + ".0"
}
func GetCurlVersion() string {
	// curl 8.0.0 was released on 20/03/2023.
	var timeCurrent int64 = time.Now().Unix() / 86400
	var timeStart int64 = time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC).Unix() / 86400
	// Generate some variations upon the average release cadence of 56.67 days.
	var timeDiff int = int((timeCurrent - timeStart - 60)) - int(math.Floor(math.Pow(rand.Float64(), 2) * 165))
	var minorValue int = int(timeDiff / 57)
	return "8." + strconv.Itoa(minorValue) + ".0"
}
func GetFirefoxVersion() string {
	// Firefox 128 ESR was released on 09/07/2023.
	var timeCurrent int64 = time.Now().Unix() / 86400
	var timeStart int64 = time.Date(2024, 7, 29, 0, 0, 0, 0, time.UTC).Unix() / 86400
	// Generate some variations upon the average release cadence of 28 days, however Firefox had some delays averaging to 29.67 days in total.
	var timeDiff = timeCurrent - timeStart - 15 - int64(math.Floor(math.Pow(rand.Float64(), 2) * 45))
	var majorValue int = int(timeDiff / 29) + 128
	return strconv.Itoa(majorValue) + ".0"
}
func GetSafariVersion() string {
	var minorMap [25]int = [25]int{0, 0, 0, 1, 1,
		1, 2, 2, 2, 2, 3, 3, 3, 4, 4,
		4, 5, 5, 5, 5, 5, 6, 6, 6, 6}
	var anchoredTime time.Time = time.Now()
	var releaseYear int = anchoredTime.Year()
	var splitPoint time.Time = time.Date(releaseYear, 9, 23, 0, 0, 0, 0, time.UTC)
	if (anchoredTime.Compare(splitPoint) < 0) {
		releaseYear --
		splitPoint = time.Date(releaseYear, 9, 23, 0, 0, 0, 0, time.UTC)
	}
	splitPoint = splitPoint.AddDate(0, 0, int(math.Floor(math.Pow(rand.Float64(), 3) * 75)))
	var minorVersion = minorMap[(anchoredTime.Unix() - splitPoint.Unix()) / 1296000]
	return strconv.Itoa(releaseYear - 1999) + "." + strconv.Itoa(minorVersion)
}

func main() {
	fmt.Println("Approximated versions:")
	fmt.Printf("curl - %s\n", GetCurlVersion())
	fmt.Printf("Chrome - %s\n", GetChromeVersion())
	fmt.Printf("Firefox - %s\n", GetFirefoxVersion())
	fmt.Printf("Safari - %s\n", GetSafariVersion())
}
