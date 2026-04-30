package main

import (
	"fmt"
	"strings"
)

// Define a custom type
type Score int

// Define a function type
type Converter func(string) Score

func process(input string, conv Converter) Score {
	return conv(input)
}

func countLength(input string) Score {
	return Score(len(input))
}

func countVowels(input string) Score {
	count := 0
	for _, val := range input {
		if strings.ContainsRune("aeiou", val) {
			count++
		}
	}
	return Score(count)
}

func main() {

	result := process("aeiou", countVowels)
	fmt.Println("The Length of the input String is ", result)

}
