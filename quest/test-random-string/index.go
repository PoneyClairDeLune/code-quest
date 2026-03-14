package main

import (
	"fmt"
	"math/rand"
)

//func getChManglingChar() string {
	//return string(" ,-_:;()"[rand.Int()&7]);
//}

var clientHintGreaseNA = []string{" ", "(", ":", "-", ".", "/", ")", ";", "=", "?", "_"}
var clientHintVersionNA = []string{"8", "99", "24"}

func getGreasedChInvalidBrand(seed int) string {
	return "\"Not" + clientHintGreaseNA[seed % len(clientHintGreaseNA)] + "A" + clientHintGreaseNA[(seed + 1) % len(clientHintGreaseNA)] + "Brand\";v=\"" + clientHintVersionNA[seed % len(clientHintVersionNA)] + "\"";
}

func main() {
	//fmt.Println("Some random characters: " + getChManglingChar());
	fmt.Println("Greased string:", getGreasedChInvalidBrand(145))
	fmt.Println("Greased string:", getGreasedChInvalidBrand(rand.Int()))
}