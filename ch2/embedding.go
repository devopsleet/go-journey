package main

import "fmt"

// employee definition
type Employee struct {
	Name string
	ID   int
}

// value receiver on Employee type
func (e Employee) Description() string {
	return fmt.Sprintf("%s %s", e.Name, e.ID)
}

// Manager type
type Manager struct {
	Employee // Embedded field
	Reports  []Employee
}

func main() {

	m := Manager{
		Employee: Employee{
			Name: "XYZ",
			ID:   0002,
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)

}
