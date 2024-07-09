package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return c.r * math.Pi
}

func main() {
	circle := Circle{r: 2}
	result := circle.area()

	value := fmt.Sprintf("%.3f", result)
	fmt.Println(value)

}
