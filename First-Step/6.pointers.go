package main

import "fmt"

func main() {
	var str string
	str = "Hello world"
	printString(&str)
}

func printString(str *string) {
	fmt.Println(*str)
}
