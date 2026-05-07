package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("The Name is %s and ID is %s", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {

	m := Manager{
		Employee: Employee{
			Name: "Gagan",
			ID:   "122345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.Description())
}
