package main

import "fmt"
import "GoTutorial/go-test/pkg-math"

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := pkg_math.Average(xs)
	fmt.Println(avg)
}
