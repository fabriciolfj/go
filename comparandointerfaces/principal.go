package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int
type TestInt int

func (d *TestInt) Double() {
	*d = *d * 2
}

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleSliceInt []int

func (d DoubleSliceInt) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoubleComparer(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func test(d DoubleInt) {
	d.Double()
	fmt.Println("receive", d)
}

func main() {
	var d1 DoubleInt = 10
	var d2 DoubleInt = 10

	fmt.Println(d1 == d2)

	DoubleComparer(&d1, &d2)
	DoubleComparer(&d1, &d1)

	var t1 TestInt = 10
	DoubleComparer(&d1, &t1)
	//test(d1)

	//var ds1 = DoubleSliceInt{1, 2, 3}
	//var ds2 = DoubleSliceInt{1, 2, 3}

	//DoubleComparer(ds1, ds2) //error
}
