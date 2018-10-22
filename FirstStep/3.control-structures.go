package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	var a = 1

	if a % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
