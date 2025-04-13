package main

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
	Email   string
}

func (e *Employee) String() string {
	return fmt.Sprintf("Emp %s (Id: %d) and Address %s", e.Name, e.ID, e.Address)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m *Manager) String() string {
	return fmt.Sprintf("Manager %s (Id: %d) and Address %s", m.Name, m.ID, m.Address)
}

type Printer interface {
	String() string
}

type Person struct {
	Printer
}

func (p *Person) String() string {
	return p.Printer.String()
}

func main() {

	p := Person{&Employee{
		1, "John", "123 Main St", "555-555-5555", "john@example.com"}}
	fmt.Println(p)
}
