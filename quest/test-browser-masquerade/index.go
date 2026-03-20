package main

import (
	"fmt"
	"strconv"
	"strings"
	"math/rand"
)

//func getChManglingChar() string {
	//return string(" ,-_:;()"[rand.Int()&7]);
//}

var clientHintGreaseNA = []string{" ", "(", ":", "-", ".", "/", ")", ";", "=", "?", "_"}
var clientHintVersionNA = []string{"8", "99", "24"}
var clientHintShuffle3 = [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
var clientHintShuffle4 = [][4]int{
	{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
	{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
	{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
	{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0}}
func getGreasedChInvalidBrand(seed int) string {
	return "\"Not" + clientHintGreaseNA[seed % len(clientHintGreaseNA)] + "A" + clientHintGreaseNA[(seed + 1) % len(clientHintGreaseNA)] + "Brand\";v=\"" + clientHintVersionNA[seed % len(clientHintVersionNA)] + "\"";
}
func getGreasedChOrder(brandLength int, seed int) []int {
	switch brandLength {
		case 1:
			return []int{0}
		case 2:
			return []int{seed % brandLength, (seed + 1) % brandLength}
		case 3:
			return clientHintShuffle3[seed % len(clientHintShuffle3)][:]
		default:
			return clientHintShuffle4[seed % len(clientHintShuffle4)][:]
	}
	return []int{}
}
func getClientHintUngreased(majorVersion int) []string {
	return []string {getGreasedChInvalidBrand(majorVersion),
	"\"Chromium\";v=\"" + strconv.Itoa(majorVersion) + "\"",
	"\"Google Chrome\";v=\"" + strconv.Itoa(majorVersion) + "\""}
}
func getGreasedChUa(majorVersion int) string {
	rawCh := getClientHintUngreased(majorVersion)
	shuffleMap := getGreasedChOrder(len(rawCh), majorVersion)
	shuffledCh := make([]string, len(rawCh))
	for i, e := range shuffleMap {
		shuffledCh[e] = rawCh[i]
	}
	return strings.Join(shuffledCh, ", ")
}


func main() {
	//fmt.Println("Some random characters: " + getChManglingChar());
	fmt.Println("Greased string:", getGreasedChInvalidBrand(145))
	fmt.Println("Greased string:", getGreasedChUa(145))
	fmt.Println("Greased string:", getGreasedChInvalidBrand(rand.Int()))
}