package main

import (
	"fmt"
)

func main() {

	message := "Hi👧👦"

	var runes = []rune(message)
	var bytes = []byte(message)

	fmt.Println(runes)
	fmt.Println(bytes)

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var s1 = greetings[:2]
	var s2 = greetings[1:4]
	var s3 = greetings[3:]
	fmt.Println(greetings)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
