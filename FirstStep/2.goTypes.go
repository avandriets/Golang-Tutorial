package main

import (
	"fmt"
)

func main() {
	// STRING
	var x string = "Hello World"
	fmt.Println(x, " len=", len(x))

	x += ", another"
	fmt.Println(x)

	fmt.Println()

	// compare
	var a string = "hello"
	var b string = "world"
	fmt.Println("Compare values", a == b)

	fmt.Println()
	c := "Hello World"
	fmt.Println(c, " len=", len(c))

	// const
	const d string = "Hello World"
	fmt.Println(d)

	// INT
	var intVariable int = -1
	var unsignedInt uint = 1

	fmt.Println("INT variables")
	fmt.Println(intVariable, unsignedInt)

	// FLOAT
	var float32Variable float32 = 0.1
	var float64Variable float64 = 0.03

	fmt.Println("int variables")
	fmt.Println(float32Variable, float64Variable)

	MultipleVar()
}

func MultipleVar() {
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println("Multi var")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
