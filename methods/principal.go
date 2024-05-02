package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total  %d, last update %v", c.total, c.lastUpdate)
}

func (p Person) String() string {
	return fmt.Sprintf("%s, %s, age %d", p.FirstName, p.LastName, p.Age)
}

func executeIncrement(c Counter) {
	c.Increment()
}

func executeIncrementModify(c *Counter) {
	c.Increment()
}

func testNuloOrZeo(c *Counter) {
	if c == nil {
		fmt.Println("veio nulo")
		return
	}

	c.Increment()
	fmt.Println("executou")
}

func main() {
	/*p := Person{"fabricio", "jacob", 39}

	fmt.Println(p.String())

	fmt.Println("inicio")
	c := Counter{}
	executeIncrement(c)
	fmt.Println("executeIncrement", c.String())
	executeIncrementModify(&c)
	fmt.Println("executeIncrementModify", c.String())*/

	var z *Counter
	testNuloOrZeo(z)

	var p Counter
	testNuloOrZeo(&p)
}
