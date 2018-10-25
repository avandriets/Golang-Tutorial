package main

import "fmt"

func main() {
	var x [5]int

	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// INIT array
	array := [5]float64{1, 2, 3, 4, 5}

	// RANGE
	var total float64 = 0
	for _, value := range array {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// SLICE
	fmt.Println("SLICE TEST")
	for _, value := range array[1:3] {
		fmt.Println(value)
	}

	// MAP
	fmt.Println("MAP TEST")
	y := make(map[string]int)
	y["key"] = 10
	fmt.Println(y["key"])
}
