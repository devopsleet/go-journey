package main

import "fmt"

func main() {

	// var s string = "Hello World"
	// var b byte = s[6]
	// fmt.Println("Character is ", string(s[6]))
	// fmt.Println("character value is ", b)

	var s string = "Hello 😊"
	s2 := string(s[4:7])
	s3 := string(s[:5])
	s4 := string(s[6:])
	fmt.Println("Strings are ", s2, s3, s4)
	fmt.Println("length of a string", len(s))

	var a rune = 'x'
	var s1 string = string(a)
	fmt.Println(s1)

	var b byte = 'y'
	fmt.Println(b)

	c := 65
	fmt.Println(string(c))

	totalWins := map[string][]string{
		"Person1": []string{"Gagan", `"30 years old"`},
		"Person2": []string{"Dev", "40 years old"},
	}

	fmt.Println("Map view is ", totalWins)
	fmt.Println("Map with a make function")

	m := make(map[string]int, 10)
	fmt.Println("M is ", len(m))

	data := make(map[string]string)
	data["name"] = "Gagan"
	data["age"] = "32"
	data["Job"] = "Golang developer"

	v, ok := data["name"]
	v2, ok2 := data["married"]
	fmt.Printf(`%v data is present: %t`, v, ok)
	println()
	fmt.Printf("%v data is present: %t", v2, ok2)
	println()

	type person struct {
		name string
		age  int
		pet  string
	}

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet.name)

}
