package main

import (
	"fmt"
)

func main() {
	_, _, n := concatWord("George", "Prado")

	fmt.Println(n)
}

func concatWord(a, b string) (string, string, string) {
	name := a
	lastName := b
	fullName := a + " " + b

	return name, lastName, fullName
}
