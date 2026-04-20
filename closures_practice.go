package main

import (
	"fmt"
	"sort"
)

func main() {

	type person struct {
		FirstName string
		LastName  string
		age       int
	}

	people := []person{
		{"Pat", "patterson", 37},
		{"Tracy", "Bobdaugter", 23},
		{"Fred", "Fredson", 18},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)
}
