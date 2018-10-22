package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// (c *Circle) - receiver
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	var c *Circle

	// get (*Circle)
	c = new(Circle)
	fmt.Println(*c)

	c1 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c1)

	c2 := Circle{0, 0, 5}
	fmt.Println(c2)
	fmt.Println("area:", c2.area())

	a := new(Android)
	a.Talk()
}
