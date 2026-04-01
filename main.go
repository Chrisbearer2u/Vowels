package main

import (
	"fmt"
	"strings"
)

func Vowels(in string) string {

	input := strings.Fields(in)

	val1 := "aeiouhAEIOUH"

	for i, word := range input {

		if string(input[i]) == "A" && strings.ContainsAny(string(word[0]), val1) {
			input[i] = "An"

		}

	}
	return strings.Join(input, " ")
}

func main() {
	fmt.Println(Vowels("There it was. A amazing rock!"))
}
