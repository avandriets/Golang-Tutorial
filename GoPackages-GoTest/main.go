package main

import "fmt"
import "GoTutorial/GoPackages-GoTest/pkg-math"

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := pkg_math.Average(xs)
	fmt.Println(avg)
}
