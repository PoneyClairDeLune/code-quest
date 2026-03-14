package main

import (
	"fmt"
	"math/rand"
)

func getChManglingChar() string {
	return string(" ,-_:;()"[rand.Int()&7]);
}

func main() {
	fmt.Println("Some random characters: " + getChManglingChar());
}