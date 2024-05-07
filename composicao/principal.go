package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s, (%s)", e.Name, e.ID)
}

type Inner struct {
	A int
}

type Outer struct {
	Inner
	S string
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Fabricio",
			ID:   "10021",
		},
		Reports: []Employee{},
	}

	fmt.Println(m)
	fmt.Println(m.Employee.ID)

	o := Outer{
		Inner: Inner{A: 10}, S: "hello",
	}

	fmt.Println(o.Double())
}
